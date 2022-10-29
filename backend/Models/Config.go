package Models

import "gorm.io/gorm"

type ModelInterface interface {
	//GetBlogByTag(blog *[]Blog, tag string) error
	//GetAllBlogTag(tags *[]BlogTag) error
	//GetWorkByTag(works *[]Work, tag string) error
	//GetAllWorkTag(tags *[]WorkTag) error
}

type UserInterface interface {
	CreateUser(username string, password string) error
	CheckPassword(username string, password string) error
	CreateRefresh(id uint64, token string, exp int64) (map[string]string, error)
}

type BlogInterface interface {
	GetAllBlog() (*[]Blog, error)
	GetBlogByID(id uint64) (*Blog, error)
	CreateBlog(blog *Blog) error
	UpdateBlog(blog *Blog) error
	DeleteBlog(id uint64) error
	IncrementSeeBlog(blog *Blog) error
	ChangeBlogIsActive(id uint64) error
}

type WorkInterface interface {
	GetAllWork() (*[]Work, error)
	GetWorkByID(id uint64) (*Work, error)
	CreateWork(work *Work) error
	UpdateWork(work *Work) error
	DeleteWork(id uint64) error
	IncrementSeeWork(work *Work) error
	ChangeWorkIsActive(id uint64) error
}

type ContactInterface interface {
	GetAllContact() (*[]Contact, error)
	GetContactByID(id uint64) (*Contact, error)
	CreateContact(contact *Contact) error
	ReadContact(c *Contact) error
	DeleteContact(id uint64) error
}

type Model struct {
	Db *gorm.DB
}
