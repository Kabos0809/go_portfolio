package Models

import (
	"gorm.io/gorm"
)

func (m Model) GetAllBlog() (*[]Blog, error) {
	var blogs []Blog
	tx := m.Db.Preload("Tag").Begin()
	err := tx.Find(&article).Error
	if err != nil {
		tx.Rollback()
		return &article, err
	}
	tx.Commit()
	return &article, err
}

func (m Model) GetBlogByID() (*Blog, error) {
	var blog *Blog
	tx := m.Db.Preload("Tag").Begin()
	err := tx.Where("id = ?", id).Find(blog).Error
	if err != nil {
		tx.Rollback()
		return blog, err
	}
	tx.Commit()
	return blog, err
}

func (m Model) CreateBlog(blog *Blog) error {
	tx := m.Db.Begin()
	err := tx.Create(blog).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (m Model) UpdateBlog(blog *Blog) error {
	tx := m.Db.Begin()
	err := tx.Save(blog).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (m Model) DeleteBlog(id uint64) error {
	tx := m.Db.Begin()
	err := tx.Where("id = ?", id).Delete(&Article{}).Error()
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}