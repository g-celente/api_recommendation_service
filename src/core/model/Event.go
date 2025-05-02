package model

type EventRecommendation struct {
	ID        uint    `json:"id"`
	Title     string  `json:"title"`
	Theme     string  `json:"theme"`
	Location  string  `json:"location"`
	Price     float64 `json:"price"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
}
