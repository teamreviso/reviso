package jobs

import (
	"context"

	"github.com/teamreviso/code/pkg/background/wire"
	"github.com/teamreviso/code/pkg/env"
)

func SendEmailJob(ctx context.Context, args *wire.SendEmail) error {
	log := env.Log(ctx)

	log.Info("responding to message", "args", args)

	return env.SES(ctx).SendRawEmail(
		args.From,
		args.To,
		args.Subject,
		args.Txtbody,
		args.Htmlbody,
	)
}
