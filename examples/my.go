package main

import (
	"flag"
	"fmt"

	chatwork "github.com/griffin-stewie/go-chatwork"
)

var apiKey string

func init() {
	flag.StringVar(&apiKey, "key", "", "Chatwork API key")
	flag.Parse()
}

func main() {
	c := chatwork.NewClient(apiKey)

	{
		b, e := c.MyStatus()
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	{
		b, e := c.MyTasks(map[string]string{})
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	fmt.Printf("%+v\n", c.RateLimit())
}
