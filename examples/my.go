package main

import (
        chatwork "github.com/yoppi/go-chatwork"
	"flag"
        "fmt"
)

var apiKey = `api-key`

func init() {
	flag.StringVar(&apiKey, "key", "", "Chatwork API key")
	flag.Parse()
}

func main() {
        c := chatwork.NewClient(apiKey)
        fmt.Printf("%+v\n", c.MyStatus())
        fmt.Printf("%+v\n", c.MyTasks(map[string]string{}))
}
