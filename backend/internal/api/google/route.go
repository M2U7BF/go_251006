package google

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Google Directions API レスポンス構造体
type RouteResponse struct {
	Routes []struct {
		DistanceMeters int    `json:"distanceMeters"`
		Duration       string `json:"duration"`
		Polyline       struct {
			EncodedPolyline string `json:"encodedPolyline"`
		} `json:"polyline"`
	} `json:"routes"`
}

// 送信用リクエスト構造体
type RouteRequest struct {
	Origin struct {
		Location struct {
			LatLng struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"latLng"`
		} `json:"location"`
	} `json:"origin"`

	Destination struct {
		Location struct {
			LatLng struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"latLng"`
		} `json:"location"`
	} `json:"destination"`

	TravelMode             string `json:"travelMode"`
	ComputeAlternativeRoutes bool   `json:"computeAlternativeRoutes"`
	RouteModifiers struct {
		AvoidTolls    bool `json:"avoidTolls"`
		AvoidHighways bool `json:"avoidHighways"`
		AvoidFerries  bool `json:"avoidFerries"`
	} `json:"routeModifiers"`
	LanguageCode string `json:"languageCode"`
	Units        string `json:"units"`
}

// Google Routes API呼び出し関数
func FetchGoogleDirections(origin_lat float64, origin_long float64, dest_lat float64, dest_long float64) (*RouteResponse, error) {

	// --- 構造体初期化 ---
	reqBody := RouteRequest{
		TravelMode:             "WALK",
		ComputeAlternativeRoutes: false,
		LanguageCode:           "ja",
		Units:                  "METRIC",
	}

	// 出発地
	reqBody.Origin.Location.LatLng.Latitude = origin_lat
	reqBody.Origin.Location.LatLng.Longitude = origin_long

	// 到着地
	reqBody.Destination.Location.LatLng.Latitude = dest_lat
	reqBody.Destination.Location.LatLng.Longitude = dest_long
	fmt.Printf("routes args: %g,%g,%g,%g\n", origin_lat, origin_long, dest_lat, dest_long)
	// ルート条件
	reqBody.RouteModifiers.AvoidTolls = false
	reqBody.RouteModifiers.AvoidHighways = false
	reqBody.RouteModifiers.AvoidFerries = false

	url := os.Getenv("GOOGLE_ROUTES_API_ENDPOINT")

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", GetGoogleMapsAPIKey())
	req.Header.Set("X-Goog-FieldMask", "routes.duration,routes.distanceMeters,routes.polyline.encodedPolyline")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var routeResp RouteResponse
	if err := json.Unmarshal(respBody, &routeResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &routeResp, nil
}
