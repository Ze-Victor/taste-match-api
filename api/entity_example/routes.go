package entity_example

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(routerGroup *gin.RouterGroup, exampleController *EntityExampleController) {

	// routerGroup.GET("/health", exampleController.Find)
	// routerGroup.GET("/user", exampleController.FindAllUsers)
	// routerGroup.POST("/user", exampleController.Create)
	// routerGroup.DELETE("/user/:id", exampleController.Delete)
	// routerGroup.PATCH("/user/:id", exampleController.Update)

}
