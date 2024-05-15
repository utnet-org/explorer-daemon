package database

import (
	"explorer-daemon/config"
	"explorer-daemon/model"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"
)

// ConnectDB connect to db
func ConnectDB() {
	allModels := []interface{}{
		&model.Example{},
		&model.Chip{},
	}
	var err error
	p := config.EnvLoad("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic(err)
	}

	var newLog = &log.Logger{
		Out: os.Stdout,
		Formatter: &log.TextFormatter{
			// 确保日志时间格式与log包的LstdFlags相似
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}
	//sqlLog := logger.New(log.New(os.Stdout, "[SQL] ", log.LstdFlags), logger.Config{
	sqlLog := logger.New(newLog, logger.Config{
		SlowThreshold:             500 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.EnvLoad("DB_HOST"), port, config.EnvLoad("DB_USER"), config.EnvLoad("DB_PASSWORD"), config.EnvLoad("DB_NAME"))
	log.Debug(dsn)
	if DB, err = gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			PrepareStmt:                              true, // 开启自动更新UpdatedAt字段
			Logger:                                   sqlLog,
		}); err != nil {
		log.Panic("failed to connect database")
	}

	//创表
	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			if err = DB.AutoMigrate(m); err != nil {
				log.Panic(err)
			}
		}
	}
	log.Info("Database connected")
}
