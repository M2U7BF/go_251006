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
	fmt.Printf("受け取ったフォーム: %s,%s,%s\n", req.Address, req.LLimitTravelExpenses, req.ULimitTravelExpenses)

	geo, err := google.FetchGeocode(req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	places := google.FetchGooglePlacesTextSearch(req.Address + " 最寄り駅")
	// fmt.Println(places)

	fmt.Println(places)
	for i := range places {
		res2, err := google.FetchGoogleDirections(
			geo.Results[0].Geometry.Location.Lat,
			geo.Results[0].Geometry.Location.Lng,
			places[i].Location.Latitude,
			places[i].Location.Longitude,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("name:%s, DistanceMeters: %d\n", places[i].DisplayName.Text, res2.Routes[0].DistanceMeters)
	}

	// レスポンスを返す
	c.JSON(http.StatusOK, gin.H{
		"l_limit_travel_expenses": req.LLimitTravelExpenses,
		"u_limit_travel_expenses": req.ULimitTravelExpenses,
	})
}
