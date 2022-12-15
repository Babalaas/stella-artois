package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID            string `gorm:"primaryKey"`
	User_ID       string `gorm:"type:uuid;not null"`
	Collection_ID string `gorm:"type:uuid"`
	Caption       string `gorm:"type:varchar(255);not null"`
	Location      string `gorm:"type:varchar(255);not null"`
	Created_At    string `gorm:"type:timestamp;not null"`
	Image         string `gorm:"type:varchar(255);not null"`
	Image2        string `gorm:"type:varchar(255)"`
	Drink_Number  int    `gorm:"type:int2"`
	Like_Count    int    `gorm:"type:int4;not null"`
}

func (entity *Post) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.Created_At = time.Now().Local().String()
	entity.Like_Count = 0
	return nil
}
