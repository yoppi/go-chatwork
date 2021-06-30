package examples

import (
	"fmt"

	"github.com/lettenj61/go-chatwork"
)

func My(key string) {
	c := chatwork.NewClient(key)
	if s, err := c.MyStatus(); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("%+v\n", s)
	}

	if res, err := c.MyTasks(&chatwork.Params{}); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("%+v\n", res)
	}
}
