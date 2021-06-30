package examples

import (
	"fmt"

	"github.com/lettenj61/go-chatwork"
)

func Rooms(key string) {
	// GET
	c := chatwork.NewClient(key)

	if res, err := c.Rooms(); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("%+v\n", res)
	}

	// fmt.Printf("%+v\n", c.Rooms())
	// fmt.Printf("%+v\n", c.Room(`room-id`))
	// fmt.Printf("%+v\n", c.RoomMembers(`room-id`))
	// fmt.Print(c.RoomMessages(`room-id`))
	// fmt.Printf("%+v\n", c.RoomMessage(`room-id`, `message-id`))
	// fmt.Printf("%+v\n", c.RoomTasks(`room-id`))
	// fmt.Printf("%+v\n", c.RoomTask(`room-id`, `task-id`))
	// fmt.Printf("%+v\n", c.RoomFiles(`room-id`, map[string]string{}))
	// fmt.Printf("%+v\n", c.RoomFile(`room-id`, `file-id`))

	// POST
	// c.CreateRoom(map[string]string{
	// 	"name":              "Test Room",
	// 	"members_admin_ids": `user-id`,
	// 	"description":       "テスト",
	// })

	// PUT
	// c.UpdateRoom(`room-id`, map[string]string{
	// 	"name":        "テストルーム",
	// 	"description": "Update description",
	// })
	// c.UpdateRoomMembers(`room-id`, map[string]string{
	// 	"members_admin_ids":    `user-id`,
	// 	"members_member_ids":   `user-id`,
	// 	"members_readonly_ids": `user-id`,
	// })

	// DELETE
	// c.DeleteRoom(`room-id`, map[string]string{
	// 	"action_type": "delete",
	// })
}
