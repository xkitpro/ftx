package ftx

import "net/http"

type Future struct {
	Ask          float64
	Bid          float64
	Change1H     float64
	Change24H    float64
	ChangeBod    float64
	Description  string
	Enabled      bool
	Expired      bool
	Group        string
	Last         float64
	Name         string
	Perpetual    bool
	VolumeUSD24H float64
}

func (c *Client) ListFutures() ([]Future, *http.Response, error) {
	req, err := c.NewRequest("GET", "/futures", nil)
	if err != nil {
		return nil, nil, err
	}

	var v []Future
	resp, err := c.Do(req, false, &v)

	return v, resp, err
}
