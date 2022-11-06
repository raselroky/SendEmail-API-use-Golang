package email

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"main.go/db"
)

func CreateMailData(c *gin.Context) {

	email_key := Email{}
	c.ShouldBindJSON(&email_key)
	err := CreateEmailTable(&email_key)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(200, email_key)

}

//
func GetAll_email(email_key *[]Email) (err error) {
Db:=db.ConnectDB()
	if err = Db.Find(&email_key).Error; err != nil {
		return err
	}
	return nil
}

func GetAll_Email_key(c *gin.Context) {
	var email_key []Email
	err := GetAll_email(&email_key)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(200, email_key)
}
