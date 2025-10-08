package router_api

import (
	"go_251006/internal/api/search"
	"go_251006/internal/api/google"
	"github.com/gin-gonic/gin"
)

func SetApiRouter (router *gin.Engine) {
	router.POST("/search", search.SearchHandler)
	router.GET("/geocode", google.GetGeocodeHandler)
}