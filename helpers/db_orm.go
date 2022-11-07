package helpers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	dsn := "host=db-eu-1.cyhypnvh2iog.eu-central-1.rds.amazonaws.com port=5432 dbname=auditt-taks2 user=admindb password=VLir6cqvbc9zMwC sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
