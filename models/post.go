package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID            string    `gorm:"type:uuid;primaryKey"`
	Profile_ID    string    `gorm:"type:uuid;not null"`
	Collection_ID string    `gorm:"type:uuid"`
	Caption       string    `gorm:"type:varchar(255);not null"`
	Location      string    `gorm:"type:varchar(255);not null"`
	Created       time.Time `gorm:"type:timestamp without time zone;not null"`
	Image         string    `gorm:"type:varchar(255);not null"`
	Image2        string    `gorm:"type:varchar(255)"`
	Drink_Number  int       `gorm:"type:int2"`
	Like_Count    int       `gorm:"type:int4;not null"`
}

func (post *Post) BeforeCreate(db *gorm.DB) error {
	if post.Collection_ID == "" {
		post.Collection_ID = uuid.Nil.String()
	}

	post.ID = uuid.New().String()
	post.Created = time.Now().Local()
	post.Created.Format(time.RFC3339)

	post.Like_Count = 0
	return nil
}
