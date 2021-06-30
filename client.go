package chatwork

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var ErrHttpFail = errors.New("chatwork: bad request")

const (
	baseURL = `https://api.chatwork.com/v2`
)

type Http interface {
	Get()
	Post()
	Put()
	Delete()
}

type Client struct {
	apiKey  string
	baseURL string
	Http
}

type Params = map[string]string

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey, baseURL: baseURL}
}

func (c *Client) Get(endpoint string, params *Params) ([]byte, error) {
	return c.execute("GET", endpoint, params)
}

func (c *Client) Post(endpoint string, params *Params) ([]byte, error) {
	return c.execute("POST", endpoint, params)
}

func (c *Client) Put(endpoint string, params *Params) ([]byte, error) {
	return c.execute("PUT", endpoint, params)
}

func (c *Client) Delete(endpoint string, params *Params) ([]byte, error) {
	return c.execute("DELETE", endpoint, params)
}

func (c *Client) buildUrl(baseUrl, endpoint string, params *Params) string {
	query := &url.Values{}
	for k, v := range *params {
		query.Add(k, v)
	}
	return baseUrl + endpoint + "?" + query.Encode()
}

func (c *Client) buildBody(params *Params) url.Values {
	body := url.Values{}
	for k, v := range *params {
		body.Add(k, v)
	}
	return body
}

func (c *Client) parseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	if isErrorStatus(resp) {
		return nil, httpFailure(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func isErrorStatus(resp *http.Response) bool {
	code4XX := resp.StatusCode - 400
	code5XX := resp.StatusCode - 500

	is4XX := code4XX > 0 && code4XX < 100
	is5XX := code5XX > 0 && code5XX < 100

	return is4XX || is5XX
}

func httpFailure(resp *http.Response) error {
	return fmt.Errorf("bad request: %s: %w", resp.Status, ErrHttpFail)
}

func (c *Client) execute(method, endpoint string, params *Params) ([]byte, error) {
	httpClient := &http.Client{}

	var (
		req    *http.Request
		reqerr error
	)

	if method != "GET" {
		req, reqerr = http.NewRequest(method, c.baseURL+endpoint, bytes.NewBufferString(c.buildBody(params).Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, reqerr = http.NewRequest(method, c.buildUrl(c.baseURL, endpoint, params), nil)
	}
	if reqerr != nil {
		return nil, reqerr
	}

	req.Header.Add("X-ChatWorkToken", c.apiKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return c.parseBody(resp)
}
