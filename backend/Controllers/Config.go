package Controllers

import (
	"github.com/kabos0809/go_portfolio/Models"
)

type ModelController struct {
	Model Models.ModelInterface
}

type BlogController struct {
	Model Models.BlogInterface
}

type WorkController struct {
	Model Models.WorkInterface
}

type ContactController struct {
	Model Models.ContactInterface
}

type UesrController struct {
	Model Models.UserInterface
}