package config

import (
	"fmt"
	"gorm.io/gorm"
	"os"
)

type DBconfig struct {
	User string
	DB string
	Pass string
	Host string
}