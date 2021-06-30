package examples

import (
	"fmt"

	"github.com/lettenj61/go-chatwork"
)

func Contacts(key string) {
	c := chatwork.NewClient(key)

	if res, err := c.Contacts(); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("%+v\n", res)
	}
}
