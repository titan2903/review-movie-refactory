package seed

import (
	"log"
	"review_movie/entities"

	"github.com/jinzhu/gorm"
)

var users = []entities.User{
	entities.User{
		FullName: "Steven William",
		Email:    "steven@gmail.com",
		Password: "password",
		Role: "admin",
	},
	entities.User{
		FullName: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
		Role: "admin",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&entities.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&entities.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&entities.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}