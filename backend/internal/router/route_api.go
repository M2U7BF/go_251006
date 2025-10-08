package router_api

import (
	"go_251006/internal/api/search"
	"github.com/gin-gonic/gin"
)

func SetApiRouter (router *gin.Engine) {
	router.POST("/search", search.SearchHandler)
}