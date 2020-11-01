package hdnsclient

import (
	"fmt"
	"net/http"
)

//RecordResp is a struct to unmarshall a single record response into
type RecordResp struct {
	Record interface{} `json:"record"`
}

//ListResp is a struct to unmarshall a list of records response into
type ListResp struct {
	Records []Record `json:"records"`
	Meta    Meta     `json:"meta"`
}

//Meta holds information by the server about pagination of the response
type Meta struct {
	Pagination Pagination `json:"pagination"`
}

//Pagination holds information about the specific pagination of the response
type Pagination struct {
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	LastPage     int `json:"last_page"`
	TotalEntries int `json:"total_entries"`
}

//Record is a single dns entry
type Record struct {
	Type     string     `json:"type"`
	ID       string     `json:"id"`
	Created  CustomTime `json:"created"`
	Modified CustomTime `json:"modified"`
	ZoneID   string     `json:"zone_id"`
	Name     string     `json:"name"`
	Value    string     `json:"value"`
	TTL      int        `json:"ttl"`
}

//RecordsListOptions configures the request
type RecordsListOptions struct {
	ItemsPerPage int `json:"per_page"`
	Page         int `json:"page"`
}

//GetRecord fetches a single record by ID
func (c *Client) GetRecord(recordID string) (*RecordResp, error) {
	// Get Record (GET https://dns.hetzner.com/api/v1/records/1)

	// Create request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/records/%s", c.baseURL, recordID), nil)

	if err != nil {
		return nil, err
	}

	res := RecordResp{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil

}

//GetRecords fetches a list of records by zoneID
func (c *Client) GetRecords(zoneID string, options *RecordsListOptions) (*ListResp, error) {
	// Get Records (GET https://dns.hetzner.com/api/v1/records?zone_id=1)
	limit := 100
	page := 1
	if options != nil {
		limit = options.ItemsPerPage
		page = options.Page
	}

	// Create request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/records?zone_id=%s&per_page=%d&page=%d", c.baseURL, zoneID, limit, page), nil)

	if err != nil {
		return nil, err
	}

	res := ListResp{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
