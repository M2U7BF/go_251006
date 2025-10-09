package search

import (
	"fmt"
	"go_251006/internal/api/google"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchRequest struct {
    Address              string `json:"address"`
    LLimitTravelExpenses string `json:"l_limit_travel_expenses"`
    ULimitTravelExpenses string `json:"u_limit_travel_expenses"`
}

func SearchHandler(c *gin.Context) {
	var req SearchRequest

	// フォームデータをバインド
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// geocodingAPI呼び出し
	// geo, err := google.FetchGeocode(req.Address)
	// if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// }
	var station_names []string
	station_names = google.FetchGooglePlacesTextSearch(req.Address + " 最寄り駅")

	fmt.Printf("受け取ったフォーム: %s,%s,%s\n", req.Address, req.LLimitTravelExpenses, req.ULimitTravelExpenses)
	fmt.Println(station_names)
	// fmt.Printf("lat:%g,lan:%g\n", geo.Results[0].Geometry.Location.Lat, geo.Results[0].Geometry.Location.Lng)
	// fmt.Printf("lat:%g,lan:%g\n", gpl.Place.displayName)
	// result := "検索結果: " + req.Address

	// レスポンスを返す
	c.JSON(http.StatusOK, gin.H{
		"address":                 req.Address,
		"l_limit_travel_expenses": req.LLimitTravelExpenses,
		"u_limit_travel_expenses": req.ULimitTravelExpenses,
	})
}