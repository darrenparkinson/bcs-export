package main

import (
	"time"

	"github.com/schollz/progressbar/v3"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (app *application) mustLoadDatabaseConfig() {
	if app.dsn == "" {
		app.logger.Fatal().Msg("missing database dsn environment")
	}
	dbLogger := logger.New(
		app.logger,
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(sqlserver.Open(app.dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		app.logger.Fatal().Err(err).Msg("problem connecting to the database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		app.logger.Fatal().Err(err).Msg("problem connecting to the database")
	}
	// gorm.DefaultCallback.Create().Remove("mssql:set_identity_insert")
	app.db = db
	app.sqlDB = sqlDB
}

func uploadObjectsToDatabase[T any](db *gorm.DB, objects []*T) {
	// startTime := time.Now()
	const maxConnections = 40
	semaphore := make(chan struct{}, maxConnections)
	results := make(chan string)
	length := len(objects)

	if length < 1 {
		return
	}

	db.AutoMigrate(objects[0])

	for _, o := range objects {
		go func(o *T) {
			semaphore <- struct{}{} // use up a connection
			db.Save(&o)
			results <- "."
			<-semaphore // free up a connection
		}(o)
	}
	bar := progressbar.Default(int64(length))
	for i := 0; i < length; i++ {
		<-results
		bar.Add(1)
	}
	close(semaphore)
	close(results)
	// seconds := time.Since(startTime).Seconds()
	// app.logger.Info().Float64("seconds", seconds).Msg("time since start")
}
