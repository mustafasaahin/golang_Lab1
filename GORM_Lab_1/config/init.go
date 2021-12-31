package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	appDBName         = "mustafa"
	appDBHost         = "localhost"
	appDBUserName     = "postgres"
	appDBUserPassword = "31298"
	DB                *gorm.DB
)

func InitDB() {
	cnnString := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s sslmode=disable",
		appDBHost,
		appDBUserName,
		appDBUserPassword,
		appDBName)

	var err error
	DB, err = gorm.Open(postgres.Open(cnnString), &gorm.Config{})
	if err != nil {
		log.Println("DB Error", err)
	} else {
		log.Println("DB Connected")
	}
}
