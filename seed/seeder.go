package seed

import (
	"log"
	"review_movie/user"

	"github.com/jinzhu/gorm"
)

var users = []user.User{
	user.User{
		FullName: "Steven William",
		Email:    "steven@gmail.com",
		Password: "password",
		Role: "admin",
	},
	user.User{
		FullName: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
		Role: "admin",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&user.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&user.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&user.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}