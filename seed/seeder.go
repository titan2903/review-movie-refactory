package seed

import (
	"log"
	"review_movie/model"

	"github.com/jinzhu/gorm"
)

var users = []model.User{
	model.User{
		FullName: "Steven William",
		Email:    "steven@gmail.com",
		Password: "password",
		Role: "admin",
	},
	model.User{
		FullName: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
		Role: "admin",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&model.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&model.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&model.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}