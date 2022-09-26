package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (m Model) GetAllBlog() (*[]Blog, error) {
	var blogs []Blog
	
}