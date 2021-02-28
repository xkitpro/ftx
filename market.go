package ftx

import (
	"fmt"
	"net/http"
	"time"
)

type Candle struct {
	Close     float64
	High      float64
	Low       float64
	Open      float64
	StartTime string
}

type GetHistoricalPricesOptions struct {
	Resolution int   `url:"resolution"`
	Limit      int   `url:"limit,omitempty"`
	StartTime  int64 `url:"start_time,omitempty"`
	EndTime    int64 `url:"end_time,omitempty"`
}

func (c *Client) GetHistoricalPrices(market string, opt *GetHistoricalPricesOptions) ([]Candle, *http.Response, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("/markets/%s/candles", market), opt)
	if err != nil {
		return nil, nil, err
	}

	var v []Candle
	resp, err := c.Do(req, false, &v)

	return v, resp, err
}

type Trade struct {
	ID          int       `json:"id"`
	Liquidation bool      `json:"liquidation"`
	Price       float64   `json:"price"`
	Side        Side      `json:"side"`
	Size        float64   `json:"size"`
	Time        time.Time `json:"time"`
}

type GetTradesOptions struct {
	StartTime int64 `url:"start_time,omitempty"`
	EndTime   int64 `url:"end_time,omitempty"`
	Limit     int   `url:"limit,omitempty"`
}

func (c *Client) GetTrades(market string, opt *GetTradesOptions) ([]Trade, *http.Response, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("/markets/%s/trades", market), opt)
	if err != nil {
		return nil, nil, err
	}

	var v []Trade
	resp, err := c.Do(req, false, &v)

	return v, resp, err
}
