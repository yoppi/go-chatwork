package examples

import (
	"fmt"

	"github.com/lettenj61/go-chatwork"
)

func Me(key string) {
	c := chatwork.NewClient(key)
	if me, err := c.Me(); err != nil {
		fmt.Printf("error: %w\n", err)
	} else {
		fmt.Printf("%+v\n", me)
	}
}
