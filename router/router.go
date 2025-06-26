package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.luizalabs.com/luizalabs/luizalabs-commons-golang/reply"
	"gitlab.luizalabs.com/taste-match-api/api/entity_example"
	"gitlab.luizalabs.com/taste-match-api/application_context"
)

func SetupRestAPI(router *gin.Engine, applicationContext *application_context.ApplicationContext) *gin.Engine {
	prepareRoutes(router, applicationContext)
	//prepareAuthRoutes(router, applicationContext)
	createResponseRouteNotFound(router)
	return router
}

func prepareRoutes(router *gin.Engine, applicationContext *application_context.ApplicationContext) {
	routerGroup := router.Group("/api")
	entity_example.SetupRoutes(routerGroup, applicationContext.ExampleController)
}

// func prepareAuthRoutes(router *gin.Engine, applicationContext *application_context.ApplicationContext) {
// 	routerGroup := router.Group("table")
// 	entity_example.SetupRoutes(routerGroup, applicationContext.ExampleController)
// }

func createResponseRouteNotFound(router *gin.Engine) {
	router.NoRoute(func(ginContext *gin.Context) {
		reply.RouteNotFound(ginContext.Writer)
	})
}
