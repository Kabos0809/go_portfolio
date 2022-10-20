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
	//Tag []WorkTag `json:"tag" gorm:"foreignKey: WorkID"`
	Text string `json: "text" gorm:"type: Text; not null;"`
}

//type WorkTag struct {
//	Name string
//	WorkID uint64
//}

func (b *Work) TableName() string {
	return "work"
}

//func (b *WorkTag) TableName() string {
//	return "worktag"
//}