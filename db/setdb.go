package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func ConnectDB() (Db *gorm.DB) {

	//driver := os.Getenv("Driver")
	//host := os.Getenv("Host")
	//port := os.Getenv("Port")
	//user := os.Getenv("User")
	//password:= os.Getenv("Password1")
	//dbname := os.Getenv("Dbname")

	dsn := "host=localhost user=postgres password=12345 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("Successfully created connection to database")
	}

	return Db
}
