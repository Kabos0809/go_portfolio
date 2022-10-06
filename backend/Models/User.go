package Models

import (
	"gorm.io/gorm"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
	"github.com/joho/godotenv"
)

func (m Model) CreateUser(username string, password string) error {
	tx := m.Db.Begin()
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err := tx.Create(&User{UName: username, Pass: string(hash)}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) CheckPassword(username string, password string) error {
	var user User
	m.Db.Where(&User, "username = ?", username).Scan(&user)
	if err != bcrypt.CompareHashAndPassword([]byte(user.Pass), Hash([]byte(password))); err != nil {
		return err
	}
	return nil
}