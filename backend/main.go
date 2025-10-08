package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type SearchRequest struct {
    Address              string `json:"address"`
    LLimitTravelExpenses string `json:"l_limit_travel_expenses"`
    ULimitTravelExpenses string `json:"u_limit_travel_expenses"`
}

func main() {
	router := gin.Default()

	router.POST("/search", func(c *gin.Context) {
		var req SearchRequest

		// フォームデータをバインド
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("受け取ったフォーム: %s,%s,%s\n", req.Address, req.LLimitTravelExpenses, req.ULimitTravelExpenses)
		// result := "検索結果: " + req.Address

		// レスポンスを返す
		c.JSON(http.StatusOK, gin.H{
			"address":  req.Address,
			"l_limit_travel_expenses":  req.LLimitTravelExpenses,
			"u_limit_travel_expenses":  req.ULimitTravelExpenses,
		})
	})

	router.Run() // デフォルトで0.0.0.0:8080
}
