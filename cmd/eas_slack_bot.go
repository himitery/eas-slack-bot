package main

import (
	"github.com/gin-gonic/gin"
	"github.com/himitery/eas-slack-bot/internal/http/expo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if godotenv.Load("env/.env") != nil {
		log.Fatalln("environment variables not set")
	}

	port := func(port string) string {
		if port != "" {
			return port
		}
		return "3000"
	}(os.Getenv("APP_PORT"))

	router := gin.Default()

	eas := router.Group("/eas")
	eas.POST("/build", expo.BuildHook)
	eas.POST("/submit", expo.SubmitHook)

	if router.Run(":"+port) == nil {
		log.Fatalln("Failed to start server...")
	}
	log.Println("server started on :" + port)
}
