package Models

import (
	"gorm.io/gorm"
)

func (m Model) GetAllBlog() (*[]Blog, error) {
	var blogs []Blog
	tx := m.Db.Preload("Tag").Begin()
	if err := tx.Find(&article).Error; err != nil {
		tx.Rollback()
		return &article, err
	}
	tx.Commit()
	return &article, err
}

func (m Model) GetBlogByID() (*Blog, error) {
	var blog *Blog
	tx := m.Db.Preload("Tag").Begin()
	if err := tx.Where("id = ?", id).Find(blog).Error; err != nil {
		tx.Rollback()
		return blog, err
	}
	tx.Commit()
	return blog, err
}

func (m Model) CreateBlog(blog *Blog) error {
	tx := m.Db.Begin()
	if err := tx.Create(blog).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (m Model) UpdateBlog(blog *Blog) error {
	tx := m.Db.Begin()
	if err := tx.Save(blog).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (m Model) DeleteBlog(id uint64) error {
	tx := m.Db.Begin()
	if err := tx.Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (m Model) GetAllBlogTag(tags *[]BlogTag) error {
	if err := m.Db.Select("name").Group("name").Find(tags).Error; err != nil {
		return err
	}
	return nil
}

func (m Model) GetBlogByTag(blogs *[]Blog, tag string) error {
	tx := m.Db.Preload("Tag").Begin()
	if err := tx.Joins("inner join blogtag on blog.id = blogtag.blog_id").Where("blogtag.name = ?", tag).Preload("Tag").Find(blogs).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}