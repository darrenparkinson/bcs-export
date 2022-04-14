// Copyright 2022 Darren Parkinson. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run gen-accessors.go
package ciscobcs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
	"golang.org/x/time/rate"
)

type Client struct {
	// BaseURL for BCS API.  Set to https://demo.api.csco-bcs.com/v2 using `ciscobcs.New()`, or set directly.
	BaseURL string

	//HTTP Client to use for making requests, allowing the user to supply their own if required.
	HTTPClient *http.Client

	//CustomerID for specific customer/API Key combination.
	CustomerID string

	//API Key for Cisco BCS.
	APIKey string

	lim *rate.Limiter
}

func NewClient(customerID, apikey string, client *http.Client) (*Client, error) {
	if apikey == "" {
		return nil, ErrMissingAPIKey
	}
	if customerID == "" {
		return nil, ErrMissingCustomerID
	}
	if client == nil {
		client = &http.Client{
			Timeout: 30 * time.Second,
		}
	}
	rl := rate.NewLimiter(150, 1) // this is not documented, so we'll limit to 150/s
	c := &Client{
		// BaseURL:    "https://demo.api.csco-bcs.com/v2",
		BaseURL:    "https://api.csco-bcs.com/v2",
		HTTPClient: client,
		CustomerID: customerID,
		APIKey:     apikey,
		lim:        rl,
	}
	return c, nil
}

// ListBCSObject provides a generic function for listing the different BCS Object types which only differ in the list of Items it returns.
func ListBCSObject[T any](ctx context.Context, c *Client, customerID string, urlStr string, options ...*ListOptions) (*BCSListResponse[T], error) {
	var res BCSListResponse[T]
	err := c.makeListRequest(ctx, urlStr, &res, options...)
	return &res, err
}

// ListAllBCSObjects provides a generic function for listing all items for a particular BCS Object type without needing the user to perform pagination.
func ListAllBCSObjects[T any](ctx context.Context, c *Client, customerID string, urlStr string) ([]*T, error) {
	perPage := 500
	page := 1
	var t []*T
	for {
		res, err := ListBCSObject[T](ctx, c, customerID, urlStr, &ListOptions{Page: Int(page), PerPage: Int(perPage)})
		if err != nil {
			return nil, err
		}
		t = append(t, res.Items...)
		if page >= res.Pages {
			break
		}
		page++
	}
	return t, nil
}

func (c *Client) getURLForSlug(slug string) string {
	u, _ := url.Parse(fmt.Sprintf("%s/customer/%s", c.BaseURL, c.CustomerID))
	u.Path = path.Join(u.Path, slug)
	return u.String()
}

// makeListRequest is a helper function to specifically make requests for list items with list options and a url
func (c *Client) makeListRequest(ctx context.Context, urlStr string, v interface{}, options ...*ListOptions) error {
	urlStr, err := addRequestOptionsSlice(urlStr, options)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}
	if err := c.makeRequest(ctx, req, &v); err != nil {
		return err
	}
	return nil
}

// makeRequest provides a single function to add common items to the request.
// It will unmarshall the json body to interface provided in v.
func (c *Client) makeRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req.Header.Add("x-api-key", c.APIKey)
	if !c.lim.Allow() {
		c.lim.Wait(ctx)
	}
	rc := req.WithContext(ctx)
	res, err := c.HTTPClient.Do(rc)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var ciscobcsErr error
		switch res.StatusCode {
		case 400:
			ciscobcsErr = ErrBadRequest
		case 401:
			ciscobcsErr = ErrUnauthorized
		case 403:
			ciscobcsErr = ErrForbidden
		case 500:
			ciscobcsErr = ErrInternalError
		default:
			ciscobcsErr = ErrUnknown
		}
		return ciscobcsErr
	}
	if res.StatusCode == http.StatusCreated {
		return nil
	}
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func addRequestOptionsSlice(url string, options []*ListOptions) (string, error) {
	for _, option := range options {
		var err error
		url, err = addOptions(url, option)
		if err != nil {
			return "", err
		}
	}
	return url, nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Float64 is a helper routine that allocates a new Float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
