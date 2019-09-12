package cmd

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangxusc/auto_record/pkg/config"
	"github.com/tangxusc/auto_record/pkg/db"
	"github.com/tangxusc/auto_record/pkg/record"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func NewCommand(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:   "start",
		Short: "start application",
		Run: func(cmd *cobra.Command, args []string) {
			rand.Seed(time.Now().Unix())
			config.InitLog()

			db.Conn(ctx)
			defer db.Disconnection(ctx)

			record.Start(ctx)
			defer record.Stop(ctx)

			<-ctx.Done()
		},
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
	config.BindParameter(command)

	return command
}

func HandlerNotify(cancel context.CancelFunc) {
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, os.Kill)
		<-signals
		cancel()
	}()
}
