package db

import (
	"database/sql"
	"encoding/json"
	"fastgo/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var db *gorm.DB

type DbConfig struct {
	Dsn          string        `json:"dsn"`
	MaxIdleCount int           `json:"maxIdleCount"`
	MaxOpenCount int           `json:"maxOpenCount"`
	MaxLifetime  time.Duration `json:"maxLifetime"`
}

func New() *gorm.DB {
	mstr := configs.Get("mysql").String()
	var dbconfig DbConfig
	_ = json.Unmarshal([]byte(mstr), &dbconfig)

	sqlDB, err := sql.Open("mysql", dbconfig.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
		CreateBatchSize: 100,
	})

	if err != nil {
		panic(err)
	}
	db = gormDB
	perr := sqlDB.Ping()
	if perr != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(dbconfig.MaxIdleCount)
	sqlDB.SetMaxOpenConns(dbconfig.MaxOpenCount)
	sqlDB.SetConnMaxLifetime(dbconfig.MaxLifetime * time.Second)

	log.Println("Connected to Mysql server...")

	return db
}
