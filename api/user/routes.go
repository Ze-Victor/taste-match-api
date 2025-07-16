package user

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(routerGroup *gin.RouterGroup, userController *UserController) {

	routerGroup.GET("/health", userController.Find)
	routerGroup.GET("/user", userController.FindAllUsers)
	routerGroup.GET("/user/:id", userController.Find)
	routerGroup.POST("/user", userController.Create)
	routerGroup.DELETE("/user/:id", userController.Delete)
	routerGroup.PATCH("/user/:id", userController.Update)

}
