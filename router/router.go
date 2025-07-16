package router

import (
	"github.com/Ze-Victor/taste-match-api/taste-match-api/api/user"
	"github.com/Ze-Victor/taste-match-api/taste-match-api/application_context"
	"github.com/gin-gonic/gin"
)

func SetupRestAPI(router *gin.Engine, applicationContext *application_context.ApplicationContext) *gin.Engine {
	prepareRoutes(router, applicationContext)
	//prepareAuthRoutes(router, applicationContext)
	return router
}

func prepareRoutes(router *gin.Engine, applicationContext *application_context.ApplicationContext) {
	routerGroup := router.Group("/v1")
	user.SetupRoutes(routerGroup, applicationContext.UserController)
}

// func prepareAuthRoutes(router *gin.Engine, applicationContext *application_context.ApplicationContext) {
// 	routerGroup := router.Group("table")
// 	entity_example.SetupRoutes(routerGroup, applicationContext.ExampleController)
// }
