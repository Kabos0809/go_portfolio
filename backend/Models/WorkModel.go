package models

import "time"

type Work struct {
	ID uint `gorm:"prinmaryKey; AUTO_INCREMENT; not null;"`
	Title string `json:"title" gorm:"size: 50; type: Text; not null;"`
	Category []WorkTag `json:"category" gorm:"foreignKey: WorkID"`
	Text string `json: "text" gorm:"type: Text; not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;"`
}

type WorkTag struct {
	Name string
	WorkID uint64
}

func (b *Work) TableName() string {
	return "work"
}

func (b *WorkTag) TableName() string {
	return "worktag"
}