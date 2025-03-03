package models

type Flight struct {
	Airline  string `json:"airline"`
	Duration int64  `json:"total_duration"`
	Price    int64  `json:"price"`
}
