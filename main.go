package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.luizalabs.com/taste-match-api/application_context"
	"gitlab.luizalabs.com/taste-match-api/internal/config"
	"gitlab.luizalabs.com/taste-match-api/middlewares/server"
	"gitlab.luizalabs.com/taste-match-api/router"
)

var appContext *application_context.ApplicationContext

func init() {
	config.Load()
	// 	db, err := connection.GetConnection()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		os.Exit(3)
	// 	}

	// 	fmt.Println("[Method: main.init()] Database connection estabilished successfully...")
	appContext = application_context.NewApplicationContext()

}

func setupRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())

	// if config.Get().Env == "development" {
	// 	ginEngine.Use(gin.Logger())
	// }

	engine := router.SetupRestAPI(ginEngine, appContext)
	return engine

}

func main() {
	addr := fmt.Sprintf(`%s:%s`, config.Get().Server.Host, config.Get().Server.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: setupRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	server.WaitingForShutdown(srv)
}
