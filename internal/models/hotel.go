package models

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Hotel struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
	Price    int64    `json:"price"`
}
