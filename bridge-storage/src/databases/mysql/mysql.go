package mysql

import (
	"bridge-storage/src/models/account_model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func Connect() {
	var err error
	maxAttempts := 10

	for i := 0; i < maxAttempts; i++ {
		DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/bridge?parseTime=true"), &gorm.Config{})

		if err == nil {
			log.Println("Successfully connected to the database")
			return
		}

		log.Printf("Attempt %d failed to connect to the database: %v. Retrying in 1 second...\n", i+1, err)
		time.Sleep(1 * time.Second)
	}

	panic("Could not connect with the database after multiple attempts!")
}

func AutoMigrate() {
	err := DB.AutoMigrate(
		account_model.Account{},
	)
	if err != nil {
		panic("Could not migrate the database!")
	} else {
		log.Printf("Database migrated\n")
	}
}

func GetDB() *gorm.DB {
	return DB
}
