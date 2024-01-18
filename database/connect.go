package database

import (
	"explorer-daemon/config"
	"explorer-daemon/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"
)

// ConnectDB connect to db
func ConnectDB() {
	allModels := []interface{}{
		&model.Example{},
	}
	var err error
	p := config.EnvLoad("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic(err)
	}

	sqlLog := logger.New(log.New(os.Stdout, "[SQL] ", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.EnvLoad("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	fmt.Println(dsn)
	if DB, err = gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			PrepareStmt:                              true, // 开启自动更新UpdatedAt字段
			Logger:                                   sqlLog,
		}); err != nil {
		panic("failed to connect database")
	}

	//创表
	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			if err = DB.AutoMigrate(m); err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("Database Connected")
}
