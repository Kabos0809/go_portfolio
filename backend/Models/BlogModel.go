package models

import "time"

type Blog struct {
	ID uint `gorm:"prinmaryKey; AUTO_INCREMENT; not null;"`
	Title string `json:"title" gorm:"size: 50; type: Text; not null;"`
	Category []BlogTag `json:"category" gorm:"foreignKey:BlogID"`
	Text string `json:"text" gorm:"type: Text; not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;"`
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

