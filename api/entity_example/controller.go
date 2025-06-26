package entity_example

import (
	"github.com/gin-gonic/gin"
)

type EntityExampleController struct {
	EntityExampleBusiness EntityExampleBusiness
}

func NewEntityExampleController(entityExampleBusiness EntityExampleBusiness) *EntityExampleController {
	return &EntityExampleController{
		EntityExampleBusiness: entityExampleBusiness,
	}
}

func (controller *EntityExampleController) Find(httpContext *gin.Context) {
	// TODO impl this
	return
}

func (controller *EntityExampleController) Update(httpContext *gin.Context) {
	// TODO impl this
	return
}

func (controller *EntityExampleController) Create(httpContext *gin.Context) {
	// TODO impl this
	return
}

func (controller *EntityExampleController) Delete(httpContext *gin.Context) {
	// TODO impl this
	return
}
