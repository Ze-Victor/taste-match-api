package entity_example

import (
	"net/http"

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
	// Retorna status 200 OK e um objeto JSON
	httpContext.JSON(http.StatusOK, gin.H{
		"message": "Busca realizada com sucesso!",
		"data": gin.H{ // Exemplo com dados aninhados
			"id":   1,
			"name": "Exemplo",
		},
	})
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
