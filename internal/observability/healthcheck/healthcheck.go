package healthcheck

import (
	"context"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger = log.Logger

func SetLogger(customLogger zerolog.Logger) {
	logger = customLogger
}

func StartHealthCheckCron(ctx context.Context, cronTime int) error {
	c := cron.New()
	logger.Info().Msg("Initiated Health Check Cron")

	if cronTime == 0 {
		cronTime = 60
	}

	// cronSpec := fmt.Sprintf("@every %ds", cronTime)

	// _, err := c.AddFunc(cronSpec, func() {
	// 	queueHealthCheck(queues)
	// })

	// if err != nil {
	// 	return err
	// }

	c.Start()

	go func() {
		<-ctx.Done()
		logger.Info().Msg("Stopping Health Check Cron")
		c.Stop()
	}()

	return nil
}
