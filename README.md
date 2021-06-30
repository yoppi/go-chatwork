# go-chatwork

[ChatWork](https://www.chatwork.com/) client for Golang.

## ❗ Notice

This project is a fork of [yoppi/go-chatwork](https://github.com/yoppi/go-chatwork).

<details>

<summary>Original README</summary>

# ChatWork API Client for Golang

[ChatWork](http://www.chatwork.com/) client for Golang.

## Install

```
$ go get github.com/yoppi/go-chatwork
```

## Usage

```go
package main

import (
  chatwork "github.com/yoppi/go-chatwork"
)

func main() {
  chatwork := chatwork.NewClient(`api-key`)

  chatwork.Me()

  chatwork.MyStatus()

  chatwork.MyTasks(map[string]string {
    "assigned_by_account_id": "123",
    "status": "open"
  })

  ...
}
```

See more examples in `examples` directory.

</details>