package ftx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Side string

const (
	Buy  Side = "buy"
	Sell      = "sell"
)

type OrderType string

const (
	Limit  OrderType = "limit"
	Market           = "market"
)

type Order struct {
	ID            int
	ClientID      *string
	Market        string
	Type          string
	Side          string
	Price         float64
	Size          float64
	Status        string
	FilledSize    float64
	RemainingSize float64
	ReduceOnly    bool
	Liquidation   bool
	AvgFillPrice  *float64
	PostOnly      bool
	IOC           bool
	CreatedAt     time.Time
}

type Update struct {
	Channel string
	Type    string
	Data    json.RawMessage
}

type PlaceOrder struct {
	Market     string    `json:"market"`
	Side       Side      `json:"side"`
	Price      float64   `json:"price"`
	Type       OrderType `json:"type"`
	Size       float64   `json:"size"`
	ReduceOnly bool      `json:"reduceOnly,omitempty"`
	IOC        bool      `json:"ioc,omitempty"`
	PostOnly   bool      `json:"postOnly,omitempty"`
	ClientID   string    `json:"clientId,omitempty"`
}

type OrderHistoryOptions struct {
	Market    string `url:"market,omitempty"`
	StartTime int    `url:"start_time,omitempty"`
	EndTime   int    `url:"end_time,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

type ModifyOrderOptions struct {
	Price    float64 `json:"price,omitempty"`
	Size     float64 `json:"size,omitempty"`
	ClientID string  `json:"clientId,omitempty"`
}

func (c *Client) GetOrderHistory(opt *OrderHistoryOptions) ([]Order, *http.Response, error) {
	req, err := c.NewRequest("GET", "/orders/history", opt)
	if err != nil {
		return nil, nil, err
	}

	var v []Order
	resp, err := c.Do(req, true, &v)

	return v, resp, err
}

func (c *Client) PlaceOrder(opt *PlaceOrder) (*Order, *http.Response, error) {
	req, err := c.NewRequest("POST", "/orders", opt)
	if err != nil {
		return nil, nil, err
	}

	o := new(Order)
	resp, err := c.Do(req, true, &o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, err
}

func (c *Client) ModifyOrder(id int, opt *ModifyOrderOptions) (*Order, *http.Response, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("/orders/%d/modify", id), opt)
	if err != nil {
		return nil, nil, err
	}

	o := new(Order)
	resp, err := c.Do(req, true, &o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, err
}

type GetOpenOrdersOptions struct {
	Market string `url:"market"`
}

func (c *Client) GetOpenOrders(opt *GetOpenOrdersOptions) ([]Order, *http.Response, error) {
	req, err := c.NewRequest("GET", "/orders", opt)
	if err != nil {
		return nil, nil, err
	}

	var v []Order
	resp, err := c.Do(req, true, &v)

	return v, resp, err
}

type CancelAllOrdersOptions struct {
	Market string `json:"market,omitempty"`
}

func (c *Client) CancelAllOrders(opt *CancelAllOrdersOptions) (*http.Response, error) {
	req, err := c.NewRequest("DELETE", "/orders", opt)
	if err != nil {
		return nil, err
	}

	return c.Do(req, true, nil)
}
