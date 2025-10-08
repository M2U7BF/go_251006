package google

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// APIレスポンス用の構造体
type PlacesResponse struct {
	Places []struct {
		DisplayName struct {
			Text string `json:"text"`
		} `json:"displayName"`
	} `json:"places"`
}

func FetchGooglePlacesTextSearch(searchText string) {
	url := "https://places.googleapis.com/v1/places:searchText"

	// リクエストボディ
	fmt.Printf("searchText:%s\n", searchText)
	requestBody := map[string]string{
		"textQuery": searchText,
	}

	bodyBytes, _ := json.Marshal(requestBody)

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
	fmt.Printf("respBody:%s\n", respBody)

	var placesResp PlacesResponse
	if err := json.Unmarshal(respBody, &placesResp); err != nil {
		panic(err)
	}

	for i, place := range placesResp.Places {
		fmt.Printf("Place %d: %s\n", i, place.DisplayName.Text)
	}
}
