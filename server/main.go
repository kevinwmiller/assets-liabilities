package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"assets-liabilities/config"
	"assets-liabilities/logging"
	"assets-liabilities/models/record"
	"assets-liabilities/models/user"
	recordRepository "assets-liabilities/repositories/record"
	userRepository "assets-liabilities/repositories/user"
	"assets-liabilities/server"
)

func main() {
	logger := logrus.New()
	logger.Out = os.Stdout

	envFile := flag.String("e", "", `Specifies the environment file`)
	flag.Parse()

	if *envFile != "" {
		err := godotenv.Load(*envFile)
		if err != nil {
			logger.Fatalf("Could not load environment file '%s'\n", *envFile)
		}
	}
	ctx := context.Background()

	cfg := config.LoadConfig(".", "config")

	logger.Info("Creating database\n")
	// TODO: Initialize the repositories with this.
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode))

	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	if cfg.DebugMode {
		db.LogMode(true)
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	db.SetLogger(logger)

	ctx = context.WithValue(ctx, config.Ctx{}, cfg)
	ctx = context.WithValue(ctx, logging.Ctx{}, logger)

	// Could probably organize the repositories a bit more efficiently though I suppose it isn't too bad to
	// specify one for each type. I'd rather have a specific repository for each entity instead of one generic one
	ctx = record.UseModel(ctx, record.New(recordRepository.NewPersistedRepository(db)))
	ctx = user.UseModel(ctx, user.New(userRepository.NewPersistedRepository(db)))

	logger.Info("Creating server\n")
	s := server.New(ctx)

	logger.Info("Starting server\n")
	logger.Fatal(s.Start())
}
