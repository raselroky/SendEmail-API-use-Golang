package main

import (
	"github.com/gin-gonic/gin"
	"main.go/db"
	"main.go/email"
)
func main() {
	Db := db.ConnectDB()
	Db.AutoMigrate(&email.Email{})
	Db.AutoMigrate(&email.User{})

	r := gin.Default()

	r.POST("/email", email.CreateMailData)
	r.POST("/user", email.SendMail)
	//r.POST("/takemail",email.Take_mail)
	//email.SendMail("welcome")
	r.Run(":5000")

}
