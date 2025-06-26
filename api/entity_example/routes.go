package entity_example

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(routerGroup *gin.RouterGroup, exampleController *EntityExampleController) {

	routerGroup.GET("/address/search", exampleController.Find)
	routerGroup.POST("/address", exampleController.Create)
	routerGroup.PATCH("/address/:uuid", exampleController.Update)
	routerGroup.DELETE("/address/:uuid", exampleController.Delete)

}
