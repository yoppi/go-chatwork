package main

import (
        chatwork "github.com/yoppi/go-chatwork"
        "fmt"
)

var ApiKey = `api-key`

func main() {
        c := chatwork.NewClient(ApiKey)
        fmt.Print(c.MyStatus())
        fmt.Print(c.MyTasks(map[string]string{}))
}
