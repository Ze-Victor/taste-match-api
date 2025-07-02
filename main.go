package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Ze-Victor/taste-match-api/taste-match-api/application_context"
	"github.com/Ze-Victor/taste-match-api/taste-match-api/internal/config"
	connection "github.com/Ze-Victor/taste-match-api/taste-match-api/internal/database/postgresql"
	"github.com/Ze-Victor/taste-match-api/taste-match-api/middlewares/server"
	"github.com/Ze-Victor/taste-match-api/taste-match-api/router"
)

var appContext *application_context.ApplicationContext

func init() {
	config.Load()
	db, err := connection.GetConnection()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}

	fmt.Println("[Method: main.init()] Database connection estabilished successfully...")
	appContext = application_context.NewApplicationContext(db)

}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())

	if config.Get().Env == "development" {
		ginEngine.Use(gin.Logger())
	}

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
