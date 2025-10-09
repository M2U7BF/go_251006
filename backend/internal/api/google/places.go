package google

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// APIレスポンス用の構造体
type PlacesResponse struct {
	Places []struct {
		DisplayName struct {
			Text string `json:"text"`
		} `json:"displayName"`
	} `json:"places"`
}

func FetchGooglePlacesTextSearch(searchText string) []string {
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
	req.Header.Set("X-Goog-FieldMask", "places.displayName.text") // 取得したいフィールドのみ

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

	var station_names []string
	for i := range placesResp.Places {
		if strings.Contains(placesResp.Places[i].DisplayName.Text, "駅") {
			station_names = append(station_names, placesResp.Places[i].DisplayName.Text)
		}
	}
	return station_names
}
