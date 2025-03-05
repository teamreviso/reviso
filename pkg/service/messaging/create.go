package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/teamreviso/code/pkg/background/wire"
	"github.com/teamreviso/code/pkg/constants"
	"github.com/teamreviso/code/pkg/env"
	"github.com/teamreviso/code/pkg/models"
	"github.com/teamreviso/code/pkg/storage/dynamo"
)

func CreateMessage(
	ctx context.Context,
	docID string,
	msg *dynamo.Message,
) error {
	log := env.SLog(ctx)
	dydb := env.Dynamo(ctx)

	// always ensure docID is set
	msg.DocID = docID

	log.Info("creating message", "msg", msg)
	err := dydb.CreateMessage(msg)
	if err != nil {
		return fmt.Errorf("[messaging] error creating dynamodb message: %w", err)
	}

	err = PublishMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("[messaging] error publishing message: %w", err)
	}

	return nil
}

func CreateAiThreadMessage(
	ctx context.Context,
	docID string,
	msg *dynamo.Message,
) (bool, error) {
	log := env.SLog(ctx)
	dydb := env.Dynamo(ctx)
	bg := env.Background(ctx)

	// If the message is from a user, check if they have exceeded their message limit
	if msg.UserID != constants.RevisoUserID {
		limit, err := MessageLimit(ctx, msg.UserID)
		if err != nil {
			log.Error("error getting message limit", "error", err)
			return false, fmt.Errorf("error getting message limit: %s", err)
		}
		if !limit.Open() {
			log.Error("message limit exceeded", "userID", msg.UserID, "limit", limit.Open())
			return false, fmt.Errorf("message limit exceeded")
		}
	}

	// always ensure docID is set
	msg.DocID = docID

	log.Info("creating message", "msg", msg)
	err := dydb.CreateMessage(msg)
	if err != nil {
		return false, fmt.Errorf("[messaging] error creating dynamodb message: %w", err)
	}

	err = PublishMessage(ctx, msg)
	if err != nil {
		return false, fmt.Errorf("[messaging] error publishing message: %w", err)
	}

	if msg.UserID == constants.RevisoUserID {
		// Exit early, if Reviso sent the message
		// Also exit early if the message is still pending (ie in realtime chat)
		return false, nil
	}

	log.Info("[messaging] Building AI response message", "channelID", msg.ChannelID)
	aiMsg := &dynamo.Message{
		ContainerID:    msg.ContainerID,
		ChannelID:      msg.ChannelID,
		AuthorID:       fmt.Sprintf("!%s", msg.AuthorID),
		UserID:         constants.RevisoUserID,
		Content:        "",
		LifecycleStage: dynamo.MessageLifecycleStagePending,
		MessageMetadata: &models.MessageMetadata{
			AllowDraftEdits: msg.MessageMetadata.AllowDraftEdits,
		},
	}

	// Setting the content address before as the message's content address if it exists
	if msg.MessageMetadata.ContentAddress != "" {
		aiMsg.MessageMetadata.ContentAddressBefore = msg.MessageMetadata.ContentAddress
	}

	_, err = CreateAiThreadMessage(ctx, docID, aiMsg)
	if err != nil {
		return false, fmt.Errorf("[messaging] error creating ai message: %w", err)
	}

	// Increment the message usage, if the user is not Revis and the message was create successfully
	err = IncrMessageUsage(ctx, msg.UserID)
	if err != nil {
		return false, fmt.Errorf("[messaging] error increasing message usage: %w", err)
	}

	_, err = bg.Enqueue(
		ctx,
		&wire.RespondToThread{
			DocId:    docID,
			ThreadId: msg.ChannelID,
			AuthorId: aiMsg.AuthorID,
			UserId:   msg.UserID,

			InputMessageId:  msg.MessageID,
			OutputMessageId: aiMsg.MessageID,
		},
	)
	if err != nil {
		return false, fmt.Errorf("[messaging] error enqueueing job: %w", err)
	}

	message_type := "ask"
	if msg.MessageMetadata.AllowDraftEdits {
		message_type = "revise"
	}

	log.Info("created message", "event", "message_created", "eventType", message_type)

	return true, nil
}

func CreateAiThreadProactiveMessage(ctx context.Context, docID string, authorID string, threadID string) (bool, error) {
	log := env.Log(ctx)
	bg := env.Background(ctx)

	log.Info("[messaging] Building proactive AI message")

	aiMsg := &dynamo.Message{
		DocID:          docID,
		ContainerID:    fmt.Sprintf("%s%s", dynamo.AiThreadPrefix, threadID),
		ChannelID:      threadID,
		MessageID:      uuid.New().String(),
		CreatedAt:      time.Now().Unix(),
		AuthorID:       fmt.Sprintf("!%s", authorID),
		UserID:         constants.RevisoUserID,
		Content:        "",
		LifecycleStage: dynamo.MessageLifecycleStagePending,
	}

	_, err := CreateAiThreadMessage(ctx, docID, aiMsg)
	if err != nil {
		return false, fmt.Errorf("[messaging] error creating ai message: %w", err)
	}

	_, err = bg.Enqueue(
		ctx,
		&wire.ProactiveAiMessage{
			Type:        wire.ProactiveAiMessageType_PROACTIVE_AI_MESSAGE_TYPE_NEW_DOCUMENT,
			DocId:       docID,
			ContainerId: aiMsg.ContainerID,
			MessageId:   aiMsg.MessageID,
			ThreadId:    threadID,
		},
	)
	if err != nil {
		return false, fmt.Errorf("[messaging] error enqueueing job: %w", err)
	}

	return true, nil
}
