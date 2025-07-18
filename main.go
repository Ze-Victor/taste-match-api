package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

	configCors := cors.Config{
		AllowOrigins: []string{"https://taste-match-app.onrender.com", "http://localhost:3000"},

		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},

		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},

		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}

	ginEngine.Use(cors.New(configCors))

	if config.Get().Env == "development" {
		ginEngine.Use(gin.Logger())
	}

	engine := router.SetupRestAPI(ginEngine, appContext)
	return engine

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Get().Server.Port
	}

	addr := fmt.Sprintf(":%s", port)

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
