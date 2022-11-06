package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"main.go/db"
)

var c *gin.Context

type UserId struct {
	Id int64 `json:"id"`
}

func Take_mail(id int64) string {
	Db := db.ConnectDB()
	var mail []User
	var s string
	Db.Table("users").Select("email").Where("id = ?", id).Take(&mail)
	for _, item := range mail {
		s += string(item.Email)
	}
	return s
}
func Take_body(bd string) string {
	Db := db.ConnectDB()
	var body []Email
	var s string
	Db.Table("email_key").Select("body").Where("keywords = ?", bd).Take(&body)
	for _, item := range body {
		s += string(item.Body)
	}
	return s
}
func Take_subject(bd string) string {
	Db := db.ConnectDB()
	var subject []Email
	var s string
	Db.Table("email_key").Select("subject").Where("keywords = ?", bd).Take(&subject)

	for _, item := range subject {
		s += string(item.Subject)
	}
	return s
}

func SendMail(c *gin.Context) {
	Db := db.ConnectDB()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	user := User{}
	id := UserId{}
	c.ShouldBindJSON(&id)
	Db.Table("users").Where("id = ?", id.Id).Take(&user)
	user2 := User{
		Name:     user.Name,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Password: user.Password,
		Active:   user.Active,
	}

	//email
	em := Take_mail(id.Id)
	var to []string
	to = append(to, em)
	from := os.Getenv("From")
	host := os.Getenv("smtpHost")
	port := os.Getenv("smtpPort")
	password := os.Getenv("Password")

	auth := smtp.PlainAuth("", from, password, host)
	t, _ := template.ParseFiles("/home/raselroky/Documents/Demo./email/test.html")

	var body bytes.Buffer
	subject := Take_subject("thanks")
	bodys := Take_body("thanks")
	html := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + bodys
	body.Write([]byte("Subject: " + subject + "\n" + html))

	t.Execute(&body, user2)

	er := smtp.SendMail(host+":"+port, auth, from, to, body.Bytes())
	if er != nil {
		fmt.Print(er)
	} else {
		fmt.Println("Successfully send Email!")
	}

}
