package email

type Email struct {
	//gorm.Model
	Keywords string `json:"keywords"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}

func (b *Email) TableName() string {
	return "email_key"
}
