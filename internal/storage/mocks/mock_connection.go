package mocks

import (
	"sync"

	"github.com/AlparslanKaraguney/trux-task/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "modernc.org/sqlite" // Import the pure Go SQLite driver to enable SQLite support in non-Cgo environments
)

var once sync.Once
var connection *gorm.DB

// MockConnection initializes and returns a test SQLite database connection
func MockConnection() (*gorm.DB, func()) {
	once.Do(func() {
		connection = initializeMock()
	})

	cleanupMock := func() {
		connection.Exec("DELETE FROM smart_models")
		connection.Exec("DELETE FROM smart_features")
	}

	return connection, cleanupMock
}

func initializeMock() *gorm.DB {
	// Connect to SQLite in-memory database using modernc.org/sqlite
	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        "file::memory:?cache=shared", // Use in-memory database
	}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Enable foreign key constraints
	db.Exec("PRAGMA foreign_keys = ON;")

	// Auto migrate the schema
	db.AutoMigrate(&models.SmartModel{}, &models.SmartFeature{})
	return db
}
