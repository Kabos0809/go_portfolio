package Models

import "gorm.io/gorm"

type ModelInterface interface {
	GetAllBlogByTag(blog *[]Blog, tag string) error
	GetAllTag(tags *[]BlogTag) error
}

type BlogInterface interface {
	GetAllBlog() (*[]Blog, error)
	GetBlogByID(id uint64) (*Blog, error)
	CreateBlog(blog *Blog) error
	UpdateBlog(blog *Blog) error
	DeleteBlog(id uint64)
}

type WorkInterface interface {
	GetAllWork() (*[]Work, error)
	GetWorkByID(id uint64) (*Work, error)
	CreateWork(work *Work) error
	UpdateWork(work *Work) error
	DeleteWork(id uint64) error
}

type ContactInterface interface {
	GetAllContact() (*[]Contact, error)
	GetContactByID(id uint64) (*Contact, error)
	CreateContact(contact *Contact) error
	ReadContact(contact *Contact) error
}

type Model struct {
	Db *gorm.DB
}