package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type SearchForm struct {
	Query string `form:"query" json:"query"`
	Limit int    `form:"limit" json:"limit"`
}

func main() {
	router := gin.Default()

	router.POST("/search", func(c *gin.Context) {
		var form SearchForm

		// フォームデータをバインド
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("受け取ったフォーム:", form)
		result := "検索結果: " + form.Query

		// レスポンスを返す
		c.JSON(http.StatusOK, gin.H{
			"query":  form.Query,
			"limit":  form.Limit,
			"result": result,
		})
	})

	router.Run() // デフォルトで0.0.0.0:8080
}
