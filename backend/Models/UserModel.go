package Models

type User struct {
	ID uint64 `json:"id" gorm:"AUTO_INCREMENT; not null;"`
	UName string `json:"uname" gorm:"not null; unique;" bind:"required"`
	Password string `json:"password" gorm:"not null;" bind:"required"`
}

func (b *User) TableName() string {
	return "user"
}