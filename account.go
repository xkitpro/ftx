package ftx

import "net/http"

type Position struct {
	Cost                         float64
	EntryPrice                   *float64
	EstimatedLiquidationPrice    *float64
	Future                       string
	InitialMarginRequirement     float64
	LongOrderSize                float64
	MaintenanceMarginRequirement float64
	NetSize                      float64
	OpenSize                     float64
	RealizedPnl                  float64
	ShortOrderSize               float64
	Side                         Side
	Size                         float64
	UnrealizedPnl                float64
	CollateralUsed               float64

	RecentAverageOpenPrice float64
	RecentBreakEvenPrice   float64
}

type GetPositionsOptions struct {
	ShowAvgPrice bool `url:"showAvgPrice,omitempty"`
}

func (c *Client) GetPositions(opt *GetPositionsOptions) ([]Position, *http.Response, error) {
	req, err := c.NewRequest("GET", "/positions", opt)
	if err != nil {
		return nil, nil, err
	}

	var v []Position
	resp, err := c.Do(req, true, &v)

	return v, resp, err
}
