package main

import (
        chatwork "github.com/yoppi/go-chatwork"
        "fmt"
)

var apiKey = `api-key`

func main() {
        c := chatwork.NewClient(apiKey)
        fmt.Print(c.Contacts())
}
