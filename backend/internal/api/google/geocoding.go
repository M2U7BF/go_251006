package google

import (
	"log"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

type GeocodingRequest struct {
	Latitude   string `json:"latitude"`
	Longtitude string `json:"longtitude"`
}

// Google Geocoding APIのレスポンス全体
type GeocodeResponse struct {
	Results []GeocodeResult `json:"results"`
	Status  string          `json:"status"`
}

// 各住所結果
type GeocodeResult struct {
	FormattedAddress  string             `json:"formatted_address"`
	Geometry          Geometry           `json:"geometry"`
	PlaceID           string             `json:"place_id"`
	Types             []string           `json:"types"`
	AddressComponents []AddressComponent `json:"address_components"`
}

// 住所コンポーネント（例：都道府県、市区町村、郵便番号など）
type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// 座標やビューポート情報
type Geometry struct {
	Location     LatLng   `json:"location"`
	LocationType string   `json:"location_type"`
	Viewport     Viewport `json:"viewport"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Viewport struct {
	Northeast LatLng `json:"northeast"`
	Southwest LatLng `json:"southwest"`
}

func FetchGeocode(address, apiKey string) (*GeocodeResponse, error) {
	endpoint := "https://maps.googleapis.com/maps/api/geocode/json"

	// URLエンコード
	encodedAddr := url.QueryEscape(address)

	fullURL := fmt.Sprintf("%s?address=%s&key=%s", endpoint, encodedAddr, apiKey)

	// HTTP GETリクエスト
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request geocode API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	// レスポンスを読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// JSONデコード
	var result GeocodeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &result, nil
}

func GetGoogleMapsAPIKey() string {
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_MAPS_API_KEY not set")
		return ""
	}
	return apiKey
}

func GetGeocodeHandler(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address parameter is required"})
		return
	}

	// Google API呼び出し
	geo, err := FetchGeocode(address, GetGoogleMapsAPIKey())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 結果をそのまま返す
	c.JSON(http.StatusOK, geo)
}
