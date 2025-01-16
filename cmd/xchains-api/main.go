package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/xchains-api/cmd/xchains-api/cli"
	"github.com/scalarorg/xchains-api/internal/api"
	"github.com/scalarorg/xchains-api/internal/config"
	"github.com/scalarorg/xchains-api/internal/observability/healthcheck"
	"github.com/scalarorg/xchains-api/internal/observability/metrics"
	"github.com/scalarorg/xchains-api/internal/services"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Debug().Msg("failed to load .env file")
	}
}

func main() {
	ctx := context.Background()

	// setup cli commands and flags
	if err := cli.Setup(); err != nil {
		log.Fatal().Err(err).Msg("error while setting up cli")
	}

	// load config
	cfgPath := cli.GetConfigPath()
	cfg, err := config.New(cfgPath)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("error while loading config file: %s", cfgPath))
	}

	// initialize metrics with the metrics port from config
	metricsPort := cfg.Metrics.GetMetricsPort()
	metrics.Init(metricsPort)

	services, err := services.New(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("error while setting up staking services layer")
	}

	// Check if the replay flag is set
	if cli.GetReplayFlag() {
		// log.Info().Msg("Replay flag is set. Starting replay of unprocessable messages.")
		// err := scripts.ReplayUnprocessableMessages(ctx, cfg, queues, services.DbClient)
		// if err != nil {
		// 	log.Fatal().Err(err).Msg("error while replaying unprocessable messages")
		// }
		// return

		fmt.Println("Not implemented")
	}

	healthcheck.StartHealthCheckCron(ctx, cfg.Server.HealthCheckInterval)

	apiServer, err := api.New(ctx, cfg, services)
	if err != nil {
		log.Fatal().Err(err).Msg("error while setting up staking api service")
	}
	if err = apiServer.Start(); err != nil {
		log.Fatal().Err(err).Msg("error while starting staking api service")
	}
}
