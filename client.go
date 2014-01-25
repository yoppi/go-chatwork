package gochatwork

import (
        "net/http"
        "strings"
        "io/ioutil"
        "log"
)

type Http interface {
        Get()
        Post()
        Put()
        Delete()
}

const BaseUrl = `https://api.chatwork.com/v1`

type Client struct {
        ApiKey string
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

func (c *Client) requestUrl(baseUrl, endpoint string, params map[string]string) string {
        query := make([]string, len(params))
        for k := range params {
                query = append(query, k + "=" + params[k])
        }
        return baseUrl + endpoint + "?" + strings.Join(query, "&")
}

func (c *Client) parseBody(resp *http.Response) string {
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                log.Println(err)
                return ""
        }
        return string(body)
}

func (c *Client) execute(method, endpoint string, params map[string]string) string {
        httpClient := &http.Client{}

        req, requestErr := http.NewRequest(method, c.requestUrl(c.BaseUrl, endpoint, params), nil)
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
