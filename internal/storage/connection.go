package storage

import (
	"fmt"
	"os"
	"sync"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"github.com/AlparslanKaraguney/trux-task/pkg/logging/gormlogger"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var once sync.Once
var connection *gorm.DB

func Connection(logrusLogger *logrus.Logger) *gorm.DB {
	once.Do(func() {
		connection = initialize(logrusLogger)
	})

	return connection
}

func initialize(logrusLogger *logrus.Logger) *gorm.DB {
	databaseLogLevel := os.Getenv("DATABASE_LOG_LEVEL")
	DBLogLevel := logger.Info
	switch databaseLogLevel {
	case "silent":
		DBLogLevel = logger.Silent
	case "error":
		DBLogLevel = logger.Error
	case "warn":
		DBLogLevel = logger.Warn
	default:
		DBLogLevel = logger.Info
	}
	gormLogger := gormlogger.NewLogrusGORMLogger(logrusLogger, DBLogLevel)

	dbUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timezone=Europe/Istanbul",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrusLogger.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&models.SmartModel{}, &models.SmartFeature{})

	return db
}
