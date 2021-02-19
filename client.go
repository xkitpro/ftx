package ftx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
)

type Client struct {
	client  *http.Client
	baseURL *url.URL
	key     string
	secret  []byte
	sub     string
}

type Response struct {
	Result  interface{}
	Success bool
	Error   string
}

func NewClient(key, secret, sub string) *Client {
	baseURL, _ := url.Parse("https://ftx.com/api")
	return &Client{
		client:  http.DefaultClient,
		baseURL: baseURL,
		key:     key,
		secret:  []byte(secret),
		sub:     sub,
	}
}

func (c *Client) NewRequest(method string, path string, opt interface{}) (*http.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	var body io.Reader
	if method == "POST" || method == "DELETE" {
		b, err := json.Marshal(opt)
		if err != nil {
			return nil, err
		}

		fmt.Println(string(b))

		body = bytes.NewBuffer(b)
	}

	if method == "GET" {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}

		u.RawQuery = v.Encode()
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request, sign bool, v interface{}) (*http.Response, error) {
	if sign {
		req = c.sign(req)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := new(Response)
	response.Result = v

	r := io.TeeReader(resp.Body, os.Stderr)
	// r := resp.Body
	err = json.NewDecoder(r).Decode(&response)

	return resp, err
}

func (client *Client) sign(req *http.Request) *http.Request {
	b := new(bytes.Buffer)
	if req.Body != nil {
		rc, _ := req.GetBody()
		b.ReadFrom(rc)
	}

	ts := strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)
	signaturePayload := ts + req.Method + req.URL.RequestURI() + b.String()
	signature := sign(client.secret, signaturePayload)

	req.Header.Set("FTX-KEY", client.key)
	req.Header.Set("FTX-SIGN", signature)
	req.Header.Set("FTX-TS", ts)
	if client.sub != "" {
		req.Header.Set("FTX-SUBACCOUNT", client.sub)
	}
	return req
}
