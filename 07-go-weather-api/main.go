package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		city = "London"
	}

	response := WeatherResponse{
		City:        city,
		Temperature: 25.5,
		Condition:   "Sunny",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/weather", weatherHandler)

	fmt.Println("Weather API running on :5050")

	http.ListenAndServe(":5050", nil)
}
