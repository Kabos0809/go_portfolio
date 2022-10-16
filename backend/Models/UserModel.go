package Models

type User struct {
	ID uint64 `json:"id" gorm:"primaryKey; AUTO_INCREMENT; not null;"`
	UserName string `json:"username" gorm:"not null; unique;" binding:"required"`
	Password string `json:"password" gorm:"not null;" binding:"required"`
}

func (b *User) TableName() string {
	return "user"
}