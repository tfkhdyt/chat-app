package gorm

import (
	"context"
	"fmt"
	"os"

	"github.com/tfkhdyt/chat-app/server/user"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(lc fx.Lifecycle) *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&user.User{}); err != nil {
		panic(err)
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			if err := db.AutoMigrate(&user.User{}); err != nil {
				panic(err)
			}

			return nil
		},
	})

	return db
}
