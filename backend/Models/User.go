package Models

import (
	"time"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
	"github.com/kabos0809/go_portfolio/backend/Config"
)

func (m Model) CreateUser(username string, password string) error {
	tx := m.Db.Begin()
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err := tx.Create(&User{UserName: username, Password: string(hash)}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) CheckPassword(username string, password string) error {
	var user User
	m.Db.Find(&User{}, "user_name = ?", username).Scan(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func CreateJWT(username string, id uint64) (map[string]string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "AccessToken"
	claims["id"] = id
	claims["username"] = username
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(5 * time.Hour).Unix()

	ACCESS_TOKEN_SECRETKEY := os.Getenv("ACCESS_TOKEN_SECRETKEY")
	t, err := token.SignedString([]byte(ACCESS_TOKEN_SECRETKEY))
	if err != nil {
		return nil, err
	}

	rtoken := jwt.New(jwt.SigningMethodHS256)

	rtclaims := rtoken.Claims.(jwt.MapClaims)
	rtclaims["sub"] = "RefreshToken"
	rtclaims["id"] = id
	rtclaims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	REFRESH_TOKEN_SECRETKEY := os.Getenv("REFRESH_TOKEN_SECRETKEY")
	rt, err := rtoken.SignedString([]byte(REFRESH_TOKEN_SECRETKEY))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"AccessToken": t,
		"RefreshToken": rt,
	}, nil
}

func (m Model) CreateRefresh(id uint64, token string, exp int64) (map[string]string, error) {
	var user *User
	m.Db.First(&user, id).Scan(&user)
	t, err := CreateJWT(user.UserName, id)
	if err != nil {
		return nil, err
	}

	if err := SetBlackList(token, exp); err != nil {
		return nil, err
	}

	return t, nil
}

func SetBlackList(token string, exp int64) error {
	c := Config.ConnRedis()
	defer c.Close()

	nowTime := time.Now()
	expTime := time.Unix(int64(exp), 0)

	timeSub := expTime.Sub(nowTime).Seconds()

	_, err := c.Do("SET", token, byte(exp))
	_, err = c.Do("EXPIRE", token, int64(timeSub))
	if err != nil {
		return err
	}
	
	return nil
}

func CheckBlackList(token string) error {
	c := Config.ConnRedis()
	defer c.Close()

	_, err :=redis.String(c.Do("GET", token))
	if err != nil {
		return err
	}
	return nil
}