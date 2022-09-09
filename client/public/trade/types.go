package trade

import (
	"time"
)

type Trade struct {
	ID          int64     `json:"id"`
	Liquidation bool      `json:"liquidation"`
	Price       float64   `json:"price"`
	Side        string    `json:"side"`
	Size        float64   `json:"size"`
	Time        time.Time `json:"time"`
}

func (t Trade) Pri() float32 {
	return float32(t.Price)
}

func (t Trade) Siz() float32 {
	return float32(t.Size)
}

type ListRequest struct {
	ProductCode string `json:"-"`
	Limit       int    `json:"limit"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
}

type ListResponse struct {
	Result []Trade `json:"result"`
}
