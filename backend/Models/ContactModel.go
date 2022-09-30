package Models

import "time"

type Contact struct {
	ID uint64 `json:"id" gorm:"primaryKey; not null;"`
	SendTime time.Time `json:"send_time" gorm:"autoCreatetime"`
	Subject string `json:"subject" gorm:"size: 50; type: Text; not null;"`
	Email string `json:"email" gorm:"not null;"`
	Text string `json:"text" gorm:"size: 20000; type: Text; not null;"`
	Read bool `json:"read" gorm:"default: true;"`
}

func (b *Contact) TableName() {
	return "contact"
}