package main

import (
        chatwork "github.com/yoppi/go-chatwork"
        "fmt"
)

var apiKey = `api-key`

func main() {
	// GET
        c := chatwork.NewClient(apiKey)
        fmt.Print(c.Rooms())
        fmt.Print(c.Room("room-id"))
        fmt.Print(c.RoomMembers("room-id"))
        fmt.Print(c.RoomMessages("room-id"))
        fmt.Print(c.RoomMessage("room-id", "message-id"))
        fmt.Print(c.RoomTasks("room-id"))
        fmt.Print(c.RoomTask("room-id", "task-id"))
        fmt.Print(c.RoomFiles("room-id", map[string]string{}))
        fmt.Print(c.RoomFile("room-id", "file-id"))

	// POST
	fmt.Println(c.CreateRoom(map[string]string {
		"name": "Test Room",
		"members_admin_ids": `user-id`,
		"description": "テスト",
	}))
}
