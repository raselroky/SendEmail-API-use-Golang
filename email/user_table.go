package email

type User struct {
	//gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile" gorm:"unique"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

func (b *Email) TableNames() string {
	return "users"
}
