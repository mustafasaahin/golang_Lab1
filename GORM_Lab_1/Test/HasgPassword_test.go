package Test

import (
	"fmt"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"testing"
)

func TestHash(t *testing.T) {
	pp := config.HashPassword("ms1998ms")
	fmt.Println(pp)
}
