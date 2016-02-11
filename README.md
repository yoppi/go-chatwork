# ChatWork API Client for Golang

[ChatWork](http://www.chatwork.com/) client for Golang.

## Install

```
$ go get github.com/griffin-stewie/go-chatwork
```

## Usage

```go
package main

import (
  chatwork "github.com/griffin-stewie/go-chatwork"
)

func main() {
  chatwork := chatwork.NewClient(`api-key`)

  chatwork.Me()

  chatwork.MyStatus()

  chatwork.MyTasks(map[string]string {
    "assigned_by_account_id": "123",
    "status": "open"
  })

  chatwork.RateLimit()

  ...
}
```

### for GoogleAppEngine/Go

```go
c := appengine.NewContext(r)
client := urlfetch.Client(c)
chatwork := chatwork.NewClient(`api-key`)
chatwork.HTTPClient = client
```

See more examples in `examples` directory.
