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
		b, e := c.Rooms()
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	{
		b, e := c.Room(`room-id`)
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	{
		b, e := c.RoomMembers(`room-id`)
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	{
		b, e := c.RoomMessages(`room-id`, map[string]string{})
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	{
		b, e := c.RoomMessage(`room-id`, `message-id`)
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	{
		b, e := c.RoomTasks(`room-id`, map[string]string{})
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	{
		b, e := c.RoomTask(`room-id`, `task-id`)
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	{
		b, e := c.RoomFiles(`room-id`, map[string]string{})
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	{
		b, e := c.RoomFile(`room-id`, `file-id`, map[string]string{})
		fmt.Printf("%+v\n", b)
		fmt.Printf("%+v\n", e)
	}

	// POST
	c.CreateRoom(map[string]string{
		"name":              "Test Room",
		"members_admin_ids": `user-id`,
		"description":       "テスト",
	})

	// PUT
	c.UpdateRoom(`room-id`, map[string]string{
		"name":        "テストルーム",
		"description": "Update description",
	})
	c.UpdateRoomMembers(`room-id`, map[string]string{
		"members_admin_ids":    `user-id`,
		"members_member_ids":   `user-id`,
		"members_readonly_ids": `user-id`,
	})

	// DELETE
	c.DeleteRoom(`room-id`, map[string]string{
		"action_type": "delete",
	})

	fmt.Printf("%+v\n", c.RateLimit())
}
