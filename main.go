package main

import (
	"apistartup/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_backend_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("selamat connection ke database telah berhasil")

	var users []user.User

	db.Find(&users)

	for _, user := range users {
		fmt.Println("============")
		fmt.Println(user.Name)
		fmt.Println(user.Email)
		fmt.Println("============")
	}
}
