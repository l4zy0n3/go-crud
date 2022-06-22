package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	ID     uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name   string `gorm:"size:255;not null;unique" json:"name"`
	Author string `gorm:"size:255;not null;" json:"author"`
}

func (b *Book) GetAllBooks(db *gorm.DB) (*[]Book, error) {
	var err error
	books := []Book{}
	err = db.Debug().Model(&Book{}).Limit(100).Find(&books).Error
	if err != nil {
		return &[]Book{}, err
	}
	return &books, err
}
