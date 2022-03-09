package libs

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"log"
	"net/http"
	"strings"
)

func BiletKontrol(c *gin.Context) {

	// burayı kaldıracağız sonradan.
	log.Println("Access-Control-Allow-Origin eklendi.")
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		//c.AbortWithStatus(http.StatusOK)
	}

	var tokenStr string = c.Request.Header.Get("Authorization")
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 10)
	tokenStr = strings.Replace(tokenStr, "bearer ", "", 10)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.MySecretKey), nil
	})
	if err == nil && token.Valid {
		var user = models.User{}
		config.DB.First(&user, int64(ParseToken(tokenStr, "id").(float64)))
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Lütfen tekrar sisteme giriş yapınız."})
			c.Abort()
		}
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Lütfen tekrar sisteme giriş yapınız."})
		c.Abort()
	}
}
