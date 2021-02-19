package ftx

import (
	"fmt"
	"net/http"
)

type Candle struct {
	Close float64
	High  float64
	Low   float64
	Open  float64
}

type GetHistoricalPricesOptions struct {
	Resolution int `url:"resolution"`
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
