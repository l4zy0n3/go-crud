package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/l4zy0n3/go-crud/api/models"
)

var books = []models.Book{
	{
		Name:   "Hound of the baskervilles",
		Author: "Sir Arthur Conan Doyle",
	},
	{
		Name:   "Immortals of Meluha",
		Author: "Amish Tripathi",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Book{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Book{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range books {
		err = db.Debug().Model(&models.Book{}).Create(&books[i]).Error
		if err != nil {
			log.Fatalf("cannot seed books table: %v", err)
		}
	}
}
