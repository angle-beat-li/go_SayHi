package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/go_sayhi_db?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		fmt.Print(err.Error())
	}

}
