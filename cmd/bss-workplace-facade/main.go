package main

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-workplace-api/internal/database"
	"github.com/ozonmp/bss-workplace-api/internal/processor"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/ozonmp/bss-workplace-api/internal/config"
)

var (
	batchSize uint = 2
)

func main() {

	ctx := context.Background()

	var configFileName = "config.yml"
	//var configFileName = "config_local.yml"

	if err := config.ReadConfigYML(configFileName); err != nil {
		log.Fatal().Msgf("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(ctx, dsn, cfg.Database.Driver)
	if err != nil {
		log.Fatal().Msgf("Failed init postgres", err.Error())
		return
	}
	defer db.Close()

	eventPprocessor, err := processor.NewEventsProcessor(cfg, db)
	if err != nil {
		panic(err)
	}
	eventPprocessor.StartProcessor(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Info().Msgf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Info().Msgf("ctx.Done: %v", done)
	}

}
