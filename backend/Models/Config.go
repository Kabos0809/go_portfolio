package Model

type BlogInterface interface {
	GetAllBlog() (*[]Blog, error)
	CreateBlog(blog *Blog) error
}

type 