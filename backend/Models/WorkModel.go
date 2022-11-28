package Models

import (
	"time"
	//"image"
)

type Work struct {
	ID uint `json: "id" gorm:"prinmaryKey; AUTO_INCREMENT; not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime;"`
	Title string `json:"title" gorm:"size: 50; type: Text; not null;"`
	//Thumbnail image.Image `json:"thumbnail"`
	//WorkTags []WorkTag `gorm:"many2many:work_tags;"`
	Text string `json: "text" gorm:"type: Text; not null;"`
	SeeCount uint64 `json:"see_count" gorm:"default: 0; not null;"`
	IsActive bool `json:"is_active" gorm:"default: true"`
}

type WorkTag struct {
	id uint `json:"id" gorm:"primaryKey; AUTO_INCREMENT; not null;`
	Name string `json:"name" gorm:"not null; unique;`
	Works []Work `gorm:"many2many:work_tags;"`
}

func (b *Work) TableName() string {
	return "work"
}