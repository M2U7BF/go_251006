package main

import (
	"log"
	"go_251006/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// envファイル読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	router_api.SetApiRouter(r)
	r.Run() // デフォルトで0.0.0.0:8080
}
