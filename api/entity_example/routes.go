package entity_example

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(routerGroup *gin.RouterGroup, exampleController *EntityExampleController) {

	routerGroup.GET("/health", exampleController.Find)
	routerGroup.GET("/user", exampleController.FindAllUsers)

}
