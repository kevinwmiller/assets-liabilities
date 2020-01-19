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

	logger.Info("Creating server\n")
	s := server.New(ctx)

	logger.Info("Starting server\n")
	logger.Fatal(s.Start())
}
