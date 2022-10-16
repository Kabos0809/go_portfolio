package Models

import (
	"time"
	"image"
)

type Blog struct {
	ID uint64 `json:"id" gorm:"prinmaryKey; AUTO_INCREMENT; not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime;"`
	Thumbnail image.Image `json:"thumbnail"`
	Title string `json:"title" gorm:"size: 50; type: Text; not null;"`
	Tag []BlogTag `json:"tag" gorm:"foreignKey:BlogID"`
	Text string `json:"text" gorm:"type: Text; not null;"`
}

type BlogTag struct {
	Name string
	BlogID uint64
}

func (b *Blog) TableName() string {
	return "blog"
}

func (b *BlogTag) TableName() string {
	return "blogtag"
}

