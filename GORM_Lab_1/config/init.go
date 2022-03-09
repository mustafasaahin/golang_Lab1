package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"strings"
	"time"
)

var (
	appDBName           = "mustafa"
	appDBHost           = "localhost"
	appDBUserName       = "postgres"
	appDBUserPassword   = "31298"
	DB                  *gorm.DB
	TokenExpireDuration = time.Hour * 1
	MySecretKey         = "mus28ta06fa20"
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
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println("GetOutboundIP", err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}
