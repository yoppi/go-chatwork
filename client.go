package gochatwork

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Http interface {
	Get()
	Post()
	Put()
	Delete()
}

type Client struct {
	ApiKey  string
	BaseUrl string
	Http
}

func NewClient(apiKey string) *Client {
	return &Client{ApiKey: apiKey, BaseUrl: BaseUrl}
}

func (c *Client) Get(endpoint string, params map[string]string) string {
	return c.execute("GET", endpoint, params)
}

func (c *Client) Post(endpoint string, params map[string]string) string {
	return c.execute("POST", endpoint, params)
}

func (c *Client) Put(endpoint string, params map[string]string) string {
	return c.execute("PUT", endpoint, params)
}

func (c *Client) Delete(endpoint string, params map[string]string) string {
	return c.execute("DELETE", endpoint, params)
}

func (c *Client) buildUrl(baseUrl, endpoint string, params map[string]string) string {
	query := make([]string, len(params))
	for k := range params {
		query = append(query, k+"="+params[k])
	}
	return baseUrl + endpoint + "?" + strings.Join(query, "&")
}

func (c *Client) buildBody(params map[string]string) url.Values {
	body := url.Values{}
	for k := range params {
		body.Add(k, params[k])
	}
	return body
}

func (c *Client) parseBody(resp *http.Response) string {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(body)
}

func (c *Client) execute(method, endpoint string, params map[string]string) string {
	httpClient := &http.Client{}

	var (
		req        *http.Request
		requestErr error
	)

	if method != "GET" {
		req, requestErr = http.NewRequest(method, c.BaseUrl+endpoint, bytes.NewBufferString(c.buildBody(params).Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, requestErr = http.NewRequest(method, c.buildUrl(c.BaseUrl, endpoint, params), nil)
	}
	if requestErr != nil {
		panic(requestErr)
	}

	req.Header.Add("X-ChatWorkToken", c.ApiKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}

	return c.parseBody(resp)
}
