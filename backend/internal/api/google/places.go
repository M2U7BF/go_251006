package google

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Places []struct {
	DisplayName struct {
		Text string `json:"text"`
	} `json:"displayName"`
	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
}

// APIレスポンス用の構造体
type PlacesResponse struct {
	Places Places `json:"places"`
}

func FetchGooglePlacesTextSearch(searchText string) Places {
	url := "https://places.googleapis.com/v1/places:searchText"

	// リクエストボディ
	fmt.Printf("searchText:%s\n", searchText)
	requestBody := map[string]string{
		"textQuery":    searchText,
		"languageCode": "ja",
	}

	bodyBytes, _ := json.Marshal(requestBody)
	fmt.Printf("bodyBytes:%s\n", string(bodyBytes))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", GetGoogleMapsAPIKey())
	req.Header.Set("X-Goog-FieldMask", "places.displayName.text,places.location") // 取得したいフィールドのみ

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var placesResp PlacesResponse
	if err := json.Unmarshal(respBody, &placesResp); err != nil {
		panic(err)
	}

	var places Places
	for i := range placesResp.Places {
		if strings.Contains(placesResp.Places[i].DisplayName.Text, "駅") {
			places = append(places, placesResp.Places[i])
		}
	}
	return places
}
