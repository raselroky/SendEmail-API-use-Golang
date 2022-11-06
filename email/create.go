package email

import "main.go/db"

func CreateEmailTable(email_key *Email) (err error) {

	if err = db.ConnectDB().Create(&email_key).Error; err != nil {
		return err
	}
	return nil
}

//Show email table data
