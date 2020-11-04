package hdnsclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURLV1 = "https://dns.hetzner.com/api/v1"
)

//Client struct for a http client with api key and request url
type Client struct {
	baseURL    string
	apiKey     string
	HTTPClient *http.Client
}

//NewClient returns a http client with a set connection timeout
func NewClient(apiKey string) *Client {
	return &Client{
		baseURL: baseURLV1,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

type successResponse struct {
	//Data interface{} `json:{}`
	Data interface{}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	// Headers
	req.Header.Add("Accept", "application/json; charset=utf-8")
	req.Header.Add("Auth-API-Token", c.apiKey)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		return parseFormErr
	}

	// Fetch Request
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Api Errors
	if resp.StatusCode != http.StatusOK {
		var errResp errorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
		}
		return errors.New(errResp.Message)
	}

	// Read Response Body
	//response := successResponse{
	//	Data: v,
	//}
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}
