package models

type User struct {
	ID      int64   `json:"id"`
	FName   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   int64   `json:"phone"`
	Height  float64 `json:"height"`
	Married bool    `json:"married"`
}
