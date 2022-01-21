package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
)

func InitDb() {
	var err error

	dsn := "host=db user=pg password=admin dbname=recall port=5432 sslmode=disable TimeZone=Asia/Singapore"
	DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
