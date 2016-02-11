package gochatwork

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// HTTP interface of HTTP METHODS's methods
type HTTP interface {
	Get()
	Post()
	Put()
	Delete()
}

// ChatWorkError is error model
type ChatWorkError struct {
	Errors []string `json:"errors"`
}

// error returns Errors
func (e *ChatWorkError) error() error {
	if len(e.Errors) == 0 {
		return nil
	}

	return errors.New(strings.Join(e.Errors, ", "))
}

// Client ChatWork HTTP client
type Client struct {
	APIKey  string
	BaseURL string
	HTTP
	HTTPClient      *http.Client
	latestRateLimit *RateLimit
}

// NewClient returns ChatWork HTTP Client
func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey, BaseURL: BaseURL}
}

// Get GET method
func (c *Client) Get(endpoint string, params map[string]string) ([]byte, error) {
	return c.execute("GET", endpoint, params)
}

// Post POST method
func (c *Client) Post(endpoint string, params map[string]string) ([]byte, error) {
	return c.execute("POST", endpoint, params)
}

// Put PUT method
func (c *Client) Put(endpoint string, params map[string]string) ([]byte, error) {
	return c.execute("PUT", endpoint, params)
}

// Delete DELETE method
func (c *Client) Delete(endpoint string, params map[string]string) ([]byte, error) {
	return c.execute("DELETE", endpoint, params)
}

func (c *Client) buildURL(baseURL, endpoint string, params map[string]string) string {
	query := make([]string, len(params))
	for k := range params {
		query = append(query, k+"="+params[k])
	}
	return baseURL + endpoint + "?" + strings.Join(query, "&")
}

func (c *Client) buildBody(params map[string]string) url.Values {
	body := url.Values{}
	for k := range params {
		body.Add(k, params[k])
	}
	return body
}

func (c *Client) parseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []byte(``), err
	}

	if resp.StatusCode != 200 {
		var er ChatWorkError
		json.Unmarshal(body, &er)
		return []byte(``), er.error()
	}

	return body, nil
}

func (c *Client) execute(method, endpoint string, params map[string]string) ([]byte, error) {
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{}
	}

	var (
		req        *http.Request
		requestErr error
	)

	if method != "GET" {
		req, requestErr = http.NewRequest(method, c.BaseURL+endpoint, bytes.NewBufferString(c.buildBody(params).Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, requestErr = http.NewRequest(method, c.buildURL(c.BaseURL, endpoint, params), nil)
	}
	if requestErr != nil {
		panic(requestErr)
	}

	req.Header.Add("X-ChatWorkToken", c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	c.latestRateLimit = c.rateLimit(resp)

	if err != nil {
		return []byte(``), err
	}

	return c.parseBody(resp)
}

func (c *Client) rateLimit(resp *http.Response) *RateLimit {
	limit, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Limit"))
	remaining, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))
	resetTime, _ := strconv.ParseInt(resp.Header.Get("X-RateLimit-Reset"), 10, 64)

	return &RateLimit{Limit: limit, Remaining: remaining, ResetTime: resetTime}
}
