package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"mime"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/teamreviso/code/pkg/constants"
	"github.com/teamreviso/code/pkg/env"
	"github.com/teamreviso/code/pkg/graph/loaders"
	"github.com/teamreviso/code/pkg/graph/model"
	"github.com/teamreviso/code/pkg/models"
	"github.com/teamreviso/code/pkg/query"
	"github.com/teamreviso/code/pkg/service/messaging"
	"github.com/teamreviso/code/pkg/storage/dynamo"
	"github.com/teamreviso/code/pkg/utils"
)

// Author is the resolver for the author field.
func (r *commentNotificationPayloadValueResolver) Author(ctx context.Context, obj *model.CommentNotificationPayloadValue) (*models.User, error) {
	return loaders.GetUser(ctx, obj.AuthorID)
}

// Message is the resolver for the message field.
func (r *commentNotificationPayloadValueResolver) Message(ctx context.Context, obj *model.CommentNotificationPayloadValue) (*dynamo.Message, error) {
	currentUser, err := env.UserClaim(ctx)
	if err != nil {
		log.Errorf("error getting current user: %s", err)
		return nil, fmt.Errorf("please login")
	}

	dydb := env.Dynamo(ctx)

	channel, err := dydb.GetChannel(obj.DocumentID, obj.ChannelID)
	if err != nil {
		log.Errorf("error getting channel %s: %s", obj.ChannelID, err)
		return nil, fmt.Errorf("oops, something went wrong")
	}

	if !utils.Contains(channel.UserIDs, currentUser.Id) {
		return nil, fmt.Errorf("oops, something went wrong")
	}

	message, err := dydb.GetMessage(obj.ContainerID, obj.MessageID)
	if err != nil {
		log.Errorf("error getting messages for %s: %s", obj.ChannelID, err)
		return nil, fmt.Errorf("error getting messages: %s", err)
	}

	message, err = messaging.HydrateMentions(ctx, message)
	if err != nil {
		log.Errorf("error hydrating mentions: %s", err)
		return nil, fmt.Errorf("error hydrating mentions: %s", err)
	}

	return message, nil
}

// ID is the resolver for the id field.
func (r *messageResolver) ID(ctx context.Context, obj *dynamo.Message) (string, error) {
	return obj.MessageID, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *messageResolver) CreatedAt(ctx context.Context, obj *dynamo.Message) (*time.Time, error) {
	return utils.UnixNanoToTime(obj.CreatedAt), nil
}

// AiContent is the resolver for the aiContent field.
func (r *messageResolver) AiContent(ctx context.Context, obj *dynamo.Message) (*model.AiContent, error) {
	if obj.AIContent == nil {
		return nil, nil
	}

	concludingMessage := obj.AIContent.ConcludingMessage
	feedback := obj.AIContent.Feedback
	return &model.AiContent{
		ConcludingMessage: &concludingMessage,
		Feedback:          &feedback,
	}, nil
}

// LifecycleStage is the resolver for the lifecycleStage field.
func (r *messageResolver) LifecycleStage(ctx context.Context, obj *dynamo.Message) (model.LifecycleStage, error) {
	switch obj.LifecycleStage {
	case dynamo.MessageLifecycleStagePending:
		return model.LifecycleStagePending, nil
	case dynamo.MessageLifecycleStageRevised:
		return model.LifecycleStageRevised, nil
	case dynamo.MessageLifecycleStageCompleted:
		return model.LifecycleStageCompleted, nil
	case dynamo.MessageLifecycleStageRevising:
		return model.LifecycleStageRevising, nil
	default:
		env.SLog(ctx).Error("unknown lifecycle stage", slog.Int("stage", int(obj.LifecycleStage)))
		return model.LifecycleStagePending, fmt.Errorf("unknown lifecycle stage: %d", obj.LifecycleStage)
	}
}

// Attachments is the resolver for the attachments field.
func (r *messageResolver) Attachments(ctx context.Context, obj *dynamo.Message) ([]model.AttachmentValue, error) {
	log := env.SLog(ctx)
	out := make([]model.AttachmentValue, len(obj.Attachments.Attachments))

	for i, attachment := range obj.Attachments.Attachments {
		switch v := attachment.Value.(type) {
		case *models.Attachment_Document:
			out[i] = model.Selection{
				Content: v.Document.Content,
				Start:   v.Document.Start,
				End:     v.Document.End,
			}
		case *models.Attachment_Revision:
			rev := model.Revision{
				Start:   v.Revision.Start,
				End:     v.Revision.End,
				Updated: v.Revision.Updated,
			}

			out[i] = rev
		case *models.Attachment_Content:
			ac := model.AttachmentContent{
				Text: v.Content.Text,
				Role: v.Content.Role,
			}

			out[i] = ac

		case *models.Attachment_Error:
			ae := model.AttachmentError{
				Title: v.Error.Title,
				Text:  v.Error.Text,
				Error: v.Error.Error,
			}

			out[i] = ae
		case *models.Attachment_File:
			if v.File.ContentType == "" {
				extension := filepath.Ext(v.File.Filename)
				contentType := mime.TypeByExtension(extension)
				contentType = strings.Split(contentType, ";")[0]
				if contentType == "" && extension == ".md" {
					contentType = "text/markdown"
				}
				v.File.ContentType = contentType
			}
			af := model.AttachmentFile{
				ID:          v.File.Id,
				Filename:    v.File.Filename,
				ContentType: v.File.ContentType,
			}

			out[i] = af
		case *models.Attachment_RevisoDocument:
			ard := model.AttachedRevisoDocument{
				ID:    v.RevisoDocument.Id,
				Title: v.RevisoDocument.Title,
			}

			out[i] = ard
		default:
			log.Error("unknown attachment type", slog.String("type", fmt.Sprintf("%T", v)))
			return nil, fmt.Errorf("unknown attachment type: %T", v)
		}
	}

	return out, nil
}

// ForkedMessageIds is the resolver for the forkedMessageIds field.
func (r *messageResolver) ForkedMessageIds(ctx context.Context, obj *dynamo.Message) ([]string, error) {
	return obj.ForkedMessages, nil
}

// Metadata is the resolver for the metadata field.
func (r *messageResolver) Metadata(ctx context.Context, obj *dynamo.Message) (*model.MsgMetadata, error) {
	metadata := &model.MsgMetadata{
		ContentAddress:       obj.MessageMetadata.ContentAddress,
		ContentAddressAfter:  obj.MessageMetadata.ContentAddressAfter,
		ContentAddressBefore: obj.MessageMetadata.ContentAddressBefore,
	}

	switch status := obj.MessageMetadata.RevisionStatus; status {
	case models.RevisionStatus_REVISION_STATUS_UNSPECIFIED:
		metadata.RevisionStatus = model.MessageRevisionStatusUnspecified
	case models.RevisionStatus_REVISION_STATUS_ACCEPTED:
		metadata.RevisionStatus = model.MessageRevisionStatusAccepted
	case models.RevisionStatus_REVISION_STATUS_DECLINED:
		metadata.RevisionStatus = model.MessageRevisionStatusDeclined
	}

	if obj.MessageMetadata.ContentAddressAfterTimestamp != nil {
		time := obj.MessageMetadata.ContentAddressAfterTimestamp.AsTime()
		metadata.ContentAddressAfterTimestamp = &time
	}

	llmChoice := obj.MessageMetadata.Llm
	if llmChoice != models.LLM_CHOICE_LLM_CHOICE_UNSPECIFIED {
		switch llmChoice {
		case models.LLM_CHOICE_LLM_CLAUDE:
			metadata.Llm = model.MsgLlmClaude
		case models.LLM_CHOICE_LLM_GPT4O:
			metadata.Llm = model.MsgLlmGpt4o
		}
	}

	return metadata, nil
}

// ParentMessageID is the resolver for the parentMessageId field.
func (r *messageResolver) ParentMessageID(ctx context.Context, obj *dynamo.Message) (*string, error) {
	if strings.HasPrefix(obj.ContainerID, dynamo.MsgPrefix) {
		parentMsgID := strings.TrimPrefix(obj.ContainerID, dynamo.MsgPrefix)
		return &parentMsgID, nil
	}

	return nil, nil
}

// User is the resolver for the user field.
func (r *messageResolver) User(ctx context.Context, obj *dynamo.Message) (*models.User, error) {
	_, err := env.UserClaim(ctx)
	if err != nil {
		log.Errorf("error getting user: %s", err)
		return nil, fmt.Errorf("please login")
	}

	return loaders.GetUser(ctx, obj.UserID)
}

// Replies is the resolver for the replies field.
func (r *messageResolver) Replies(ctx context.Context, obj *dynamo.Message) ([]*dynamo.Message, error) {
	log := env.SLog(ctx)
	dydb := env.Dynamo(ctx)

	messages, err := dydb.GetRepliesToMessage(obj.MessageID)
	if err != nil {
		log.Error("error getting messages for container", "container", obj.ContainerID, "error", err)
		return nil, fmt.Errorf("error getting messages: %s", err)
	}

	return messages, nil
}

// ReplyingUsers is the resolver for the replyingUsers field.
func (r *messageResolver) ReplyingUsers(ctx context.Context, obj *dynamo.Message) ([]*models.User, error) {
	_, err := env.UserClaim(ctx)
	if err != nil {
		log.Error("error getting replyingUsers", "error", err)
		return nil, fmt.Errorf("please login")
	}

	return loaders.GetUsers(ctx, obj.ReplyingUserIds)
}

// CreateAskAiThread is the resolver for the createAskAiThread field.
func (r *mutationResolver) CreateAskAiThread(ctx context.Context, documentID string) (*dynamo.Thread, error) {
	log := env.SLog(ctx)
	currentUser, err := env.UserClaim(ctx)
	if err != nil {
		log.Error("error getting current user", "error", err)
		return nil, fmt.Errorf("please login")
	}

	// TODO: Check if they have access to the document

	thread := &dynamo.Thread{
		DocID:  documentID,
		UserID: currentUser.Id,
		Title:  constants.DefaultThreadTitle,
	}
	err = env.Dynamo(ctx).CreateThread(thread)
	if err != nil {
		log.Error("error creating thread", "error", err)
		return nil, fmt.Errorf("error creating thread: %s", err)
	}

	messaging.CreateAiThreadProactiveMessage(ctx, documentID, currentUser.Id, thread.ThreadID)

	return thread, nil
}

// CreateAskAiThreadMessage is the resolver for the createAskAiThreadMessage field.
func (r *mutationResolver) CreateAskAiThreadMessage(ctx context.Context, documentID string, threadID string, input model.MessageInput) (*dynamo.Message, error) {
	log := env.SLog(ctx)
	currentUser, err := env.UserClaim(ctx)
	if err != nil {
		log.Error("error getting current user", "error", err)
		return nil, fmt.Errorf("please login")
	}

	docAccessTbl := env.Query(ctx).DocumentAccess
	accessRow, err := docAccessTbl.Where(docAccessTbl.DocumentID.Eq(documentID), docAccessTbl.UserID.Eq(currentUser.Id)).Where(docAccessTbl.AccessLevel.In(constants.AccessLevelsWithEdit...)).First()
	if err != nil {
		log.Error("error checking access", "error", err)
		return nil, fmt.Errorf("error checking access: %s", err)
	}
	if accessRow == nil {
		log.Error("user does not have access")
		return nil, fmt.Errorf("you don't have access to this document")
	}

	messageAttachments := make([]*models.Attachment, 0, len(input.Attachments))
	for _, attachmentInput := range input.Attachments {
		if attachmentInput.Type == model.AttachmentInputTypeDraft {
			messageAttachments = append(messageAttachments, &models.Attachment{
				Value: &models.Attachment_RevisoDocument{
					RevisoDocument: &models.RevisoDocument{
						Id:    attachmentInput.ID,
						Title: attachmentInput.Name,
					},
				},
			})
		} else if attachmentInput.Type == model.AttachmentInputTypeFile {
			f := &models.FileAttachment{
				Id:       attachmentInput.ID,
				Filename: attachmentInput.Name,
			}

			if attachmentInput.ContentType != nil {
				f.ContentType = *attachmentInput.ContentType
			}

			messageAttachments = append(messageAttachments, &models.Attachment{
				Value: &models.Attachment_File{
					File: f,
				},
			})
		}
	}

	msg := &dynamo.Message{
		DocID:          documentID,
		ContainerID:    fmt.Sprintf("%s%s", dynamo.AiThreadPrefix, threadID),
		ChannelID:      threadID,
		MessageID:      uuid.New().String(),
		CreatedAt:      time.Now().Unix(),
		AuthorID:       input.AuthorID,
		UserID:         currentUser.Id,
		Content:        input.Content,
		LifecycleStage: dynamo.MessageLifecycleStageCompleted,
		Attachments: &models.AttachmentList{
			Attachments: messageAttachments,
		},
		MessageMetadata: &models.MessageMetadata{
			AllowDraftEdits: input.AllowDraftEdits,
			Intent:          models.MessageIntent_MESSAGE_INTENT_UNSPECIFIED,
			ContentAddress:  input.ContentAddress,
		},
	}

	if input.Llm != nil {
		switch *input.Llm {
		case model.MsgLlmClaude:
			msg.MessageMetadata.Llm = models.LLM_CHOICE_LLM_CLAUDE
		case model.MsgLlmGpt4o:
			msg.MessageMetadata.Llm = models.LLM_CHOICE_LLM_GPT4O
		}
	} else {
		msg.MessageMetadata.Llm = models.LLM_CHOICE_LLM_CHOICE_UNSPECIFIED
	}

	if input.Selection != nil {
		selection := models.Attachment{
			Value: &models.Attachment_Document{
				Document: &models.DocumentSelection{
					Id:      uuid.New().String(),
					Start:   input.Selection.Start,
					End:     input.Selection.End,
					Content: input.Selection.Content,
				},
			},
		}
		msg.Attachments.Attachments = append(msg.Attachments.Attachments, &selection)
	}

	// TODO: Update when we have signal about if the user is just asking
	if input.AllowDraftEdits {
		msg.MessageMetadata.Intent = models.MessageIntent_MESSAGE_INTENT_GENERATE
	}

	_, err = messaging.CreateAiThreadMessage(ctx, documentID, msg)
	if err != nil {
		log.Error("error creating message", "error", err)
		return nil, err
	}

	msg, err = messaging.HydrateMentions(ctx, msg)
	if err != nil {
		log.Error("error hydrating mentions", "error", err)
		return nil, fmt.Errorf("error hydrating mentions: %s", err)
	}

	return msg, nil
}

// UpdateMessageRevisionStatus is the resolver for the updateMessageRevisionStatus field.
func (r *mutationResolver) UpdateMessageRevisionStatus(ctx context.Context, containerID string, messageID string, status model.MessageRevisionStatus, contentAddress string) (*dynamo.Message, error) {
	//log := env.SLog(ctx)

	_, err := env.UserClaim(ctx)
	if err != nil {
		return nil, err
	}

	msg, err := messaging.UpdateRevisionStatus(ctx, containerID, messageID, status, contentAddress)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

// DocumentID is the resolver for the documentId field.
func (r *notificationResolver) DocumentID(ctx context.Context, obj *dynamo.Notification) (string, error) {
	return obj.DocID, nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *notificationResolver) CreatedAt(ctx context.Context, obj *dynamo.Notification) (*time.Time, error) {
	return utils.UnixNanoToTime(obj.CreatedAt), nil
}

// Payload is the resolver for the payload field.
func (r *notificationResolver) Payload(ctx context.Context, obj *dynamo.Notification) (model.NotificationPayloadValue, error) {
	switch val := obj.Payload.Value.(type) {
	case *models.NotificationPayload_Comment:
		comment := model.CommentNotificationPayloadValue{}
		switch val.Comment.Type {
		case models.CommentType_Comment:
			comment.CommentType = model.CommentNotificationTypeComment
		case models.CommentType_Reply:
			comment.CommentType = model.CommentNotificationTypeReply
		case models.CommentType_Mention:
			comment.CommentType = model.CommentNotificationTypeMention
		}

		comment.DocumentID = val.Comment.DocumentId
		comment.ChannelID = val.Comment.ChannelId
		comment.ContainerID = val.Comment.ContainerId
		comment.MessageID = val.Comment.MessageId
		comment.AuthorID = val.Comment.AuthorId

		return &comment, nil
	}

	return nil, errors.New("unknown notification payload type")
}

// GetAskAiThreads is the resolver for the getAskAiThreads field.
func (r *queryResolver) GetAskAiThreads(ctx context.Context, documentID string) ([]*dynamo.Thread, error) {
	log := env.SLog(ctx)
	currentUser, err := env.UserClaim(ctx)
	if err != nil {
		log.Error("error getting current user", "error", err)
		return nil, fmt.Errorf("please login")
	}

	if currentUser == nil {
		return nil, fmt.Errorf("please login")
	}

	level, err := query.AccessLevelForDocument(env.Query(ctx), documentID, currentUser.Id)
	if err != nil {
		log.Error("error getting access level", "error", err)
		return nil, fmt.Errorf("error getting access level: %s", err)
	}

	if level == constants.AccessLevelAdmin {
		threads, err := env.Dynamo(ctx).GetThreadsForDoc(documentID)
		if err != nil {
			log.Error("error getting threads", "error", err)
			return nil, fmt.Errorf("error getting threads: %s", err)
		}

		return threads, nil
	}

	if !utils.Contains(constants.AccessLevelsWithEdit, level) {
		return nil, fmt.Errorf("no access")
	}

	threads, err := env.Dynamo(ctx).GetThreadsForDocForUser(documentID, currentUser.Id)
	if err != nil {
		log.Error("error getting threads", "error", err)
		return nil, fmt.Errorf("error getting threads: %s", err)
	}

	if len(threads) == 0 {
		thread := &dynamo.Thread{
			DocID:  documentID,
			UserID: currentUser.Id,
			Title:  constants.DefaultThreadTitle,
		}
		err = env.Dynamo(ctx).CreateThread(thread)
		if err != nil {
			log.Error("error creating thread", "error", err)
			return nil, fmt.Errorf("error creating thread: %s", err)
		}

		messaging.CreateAiThreadProactiveMessage(ctx, documentID, currentUser.Id, thread.ThreadID)

		return []*dynamo.Thread{thread}, nil
	}

	return threads, nil
}

// GetAskAiThreadMessages is the resolver for the getAskAiThreadMessages field.
func (r *queryResolver) GetAskAiThreadMessages(ctx context.Context, documentID string, threadID string) ([]*dynamo.Message, error) {
	log := env.SLog(ctx)
	currentUser, err := env.UserClaim(ctx)
	if err != nil {
		return nil, fmt.Errorf("please login")
	}

	// TODO: Check if they have access to the document
	_, err = env.Dynamo(ctx).GetThreadsForDocForUser(documentID, currentUser.Id)
	if err != nil {
		log.Error("error getting threads for doc", "error", err)
		// use doesn't have access to this thread
		return nil, nil
	}

	return env.Dynamo(ctx).GetMessagesForThread(threadID)
}

// MessageUpserted is the resolver for the messageUdserted field.
func (r *subscriptionResolver) MessageUpserted(ctx context.Context, documentID string, channelID string) (<-chan *dynamo.Message, error) {
	log := env.SLog(ctx)
	dydb := env.Dynamo(ctx)

	currentUser, err := env.UserClaim(ctx)
	if err != nil {
		log.Error("error getting current user", "error", err)
		return nil, fmt.Errorf("please login")
	}

	thread, err := dydb.GetThreadForUser(documentID, channelID, currentUser.Id)
	if err != nil {
		log.Error("error getting thread", "channelID", channelID, "error", err)
	}
	if thread == nil {
		// Check for admin access
		level, err := query.AccessLevelForDocument(env.Query(ctx), documentID, currentUser.Id)
		if err != nil {
			log.Error("error getting access level", "error", err)
			return nil, fmt.Errorf("you don't have access to this channel")
		}

		if level == constants.AccessLevelAdmin {
			thread, err = dydb.GetThread(documentID, channelID)
			if err != nil {
				log.Error("error creating thread", "error", err)
				return nil, fmt.Errorf("error creating thread: %s", err)
			}
		} else {
			return nil, fmt.Errorf("you don't have access to this channel")
		}
	}

	ch := make(chan *dynamo.Message)
	go messaging.ListenForMessages(ctx, ch, documentID, channelID)
	return ch, nil
}

// ThreadUpserted is the resolver for the threadUpserted field.
func (r *subscriptionResolver) ThreadUpserted(ctx context.Context, documentID string) (<-chan *dynamo.Thread, error) {
	log := env.SLog(ctx)

	currentUser, err := env.UserClaim(ctx)
	if err != nil {
		log.Error("error getting current user", "error", err)
		return nil, fmt.Errorf("please login")
	}

	ch := make(chan *dynamo.Thread)
	go messaging.ListenForThreads(ctx, ch, documentID, currentUser.Id)
	return ch, nil
}

// ID is the resolver for the id field.
func (r *threadResolver) ID(ctx context.Context, obj *dynamo.Thread) (string, error) {
	return obj.ThreadID, nil
}

// DocumentID is the resolver for the documentId field.
func (r *threadResolver) DocumentID(ctx context.Context, obj *dynamo.Thread) (string, error) {
	return obj.DocID, nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *threadResolver) UpdatedAt(ctx context.Context, obj *dynamo.Thread) (*time.Time, error) {
	return utils.UnixNanoToTime(obj.UpdatedAt), nil
}

// Messages is the resolver for the messages field.
func (r *threadResolver) Messages(ctx context.Context, obj *dynamo.Thread) ([]*dynamo.Message, error) {
	panic(fmt.Errorf("not implemented: Messages - messages"))
}

// User is the resolver for the user field.
func (r *threadResolver) User(ctx context.Context, obj *dynamo.Thread) (*models.User, error) {
	_, err := env.UserClaim(ctx)
	if err != nil {
		log.Errorf("error getting user: %s", err)
		return nil, fmt.Errorf("please login")
	}

	return loaders.GetUser(ctx, obj.UserID)
}

// CommentNotificationPayloadValue returns CommentNotificationPayloadValueResolver implementation.
func (r *Resolver) CommentNotificationPayloadValue() CommentNotificationPayloadValueResolver {
	return &commentNotificationPayloadValueResolver{r}
}

// Message returns MessageResolver implementation.
func (r *Resolver) Message() MessageResolver { return &messageResolver{r} }

// Notification returns NotificationResolver implementation.
func (r *Resolver) Notification() NotificationResolver { return &notificationResolver{r} }

// Thread returns ThreadResolver implementation.
func (r *Resolver) Thread() ThreadResolver { return &threadResolver{r} }

type commentNotificationPayloadValueResolver struct{ *Resolver }
type messageResolver struct{ *Resolver }
type notificationResolver struct{ *Resolver }
type threadResolver struct{ *Resolver }
