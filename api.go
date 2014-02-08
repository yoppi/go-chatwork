package gochatwork

import (
	"encoding/json"
)

const BaseUrl = `https://api.chatwork.com/v1`

func (c *Client) Me() []byte {
	return c.Get("/me", map[string]string{})
}

func (c *Client) MyStatus() []byte {
	return c.Get("/my/status", map[string]string{})
}

// params keys
//  - assigned_by_account_id
//  - status: [open, done]
func (c *Client) MyTasks(params map[string]string) []byte {
	return c.Get("/my/tasks", params)
}

func (c *Client) Contacts() []byte {
	return c.Get("/contacts", map[string]string{})
}

type Room struct {
	Room_Id int
	Name string
	Type string
	Role string
	Sticky bool
	Unread_Num int
	Mention_Num int
	Mytask_Num int
	Message_Num int
	File_Num int
	Task_Num int
	Icon_Path string
	Last_Update_Time int64
}

func (c *Client) Rooms() []Room {
	ret := c.Get("/rooms", map[string]string{})
	var rooms []Room
	json.Unmarshal(ret, &rooms)
	return rooms
}

func (c *Client) Room(roomId string) Room {
	ret := c.Get("/rooms/"+roomId, map[string]string{})
	var room Room
	json.Unmarshal(ret, &room)
	return room
}

// params keys
//   * name
//   * members_admin_ids
//   - description
//   - icon_preset
//   - members_member_ids
//   - members_readonly_ids
func (c *Client) CreateRoom(params map[string]string) []byte {
	return c.Post("/rooms", params)
}

// params keys
//   - description
//   - icon_preset
//   - name
func (c *Client) UpdateRoom(roomId string, params map[string]string) []byte {
	return c.Put("/rooms/"+roomId, params)
}

// params key
//   * action_type
func (c *Client) DeleteRoom(roomId string, params map[string]string) []byte {
	return c.Delete("/rooms/"+roomId, params)
}

type Member struct {
	Account_Id int
	Role string
	Name string
	Chatwork_Id string
	Organization_Id int
	Organization_Name string
	Department string
	Avatar_Image_Url string
}

func (c *Client) RoomMembers(roomId string) []Member {
	ret := c.Get("/rooms/"+roomId+"/members", map[string]string{})
	var members []Member
	json.Unmarshal(ret, &members)
	return members
}

// params keys
//   * members_admin_ids
//   - members_member_ids
//   - members_readonly_ids
func (c *Client) UpdateRoomMembers(roomId string, params map[string]string) []byte {
	return c.Put("/rooms/"+roomId+"/members", params)
}

type Account struct {
	Account_Id int
	Name string
	Avatar_Image_Url string
}

type Message struct {
	Message_Id int
	Account Account
	Body string
	Send_Time int64
	Update_Time int64
}

// XXX: Not yet implement
func (c *Client) RoomMessages(roomId string) []byte {
	return c.Get("/rooms/"+roomId+"/messages", map[string]string{})
}

func (c *Client) PostRoomMessage(roomId string, body string) []byte {
	return c.Post("/rooms/"+roomId+"/messages", map[string]string{"body": body})
}

func (c *Client) RoomMessage(roomId, messageId string) Message {
	ret := c.Get("/rooms/"+roomId+"/messages/"+messageId, map[string]string{})
	var message Message
	json.Unmarshal(ret, &message)
	return message
}

type Task struct {
	Task_Id int
	Account Account
	Assigned_By_Account Account
	Message_Id int
	Body string
	Limit_Time int64
	Status string
}

func (c *Client) RoomTasks(roomId string) []Task {
	ret := c.Get("/rooms/"+roomId+"/tasks", map[string]string{})
	var tasks []Task
	json.Unmarshal(ret, &tasks)
	return tasks
}

// params keys
//   * body
//   * to_ids
//   - limit
func (c *Client) PostRoomTask(roomId string, params map[string]string) []byte {
	return c.Post("/rooms/"+roomId+"/tasks", params)
}

func (c *Client) RoomTask(roomId, taskId string) Task {
	ret := c.Get("/rooms/"+roomId+"/tasks/"+taskId, map[string]string{})
	var task Task
	json.Unmarshal(ret, &task)
	return task
}

type File struct {
	File_Id int
	Account Account
	Message_Id int
	Filename string
	Filesize int
	Upload_Time int64
}

// params key
//   - account_id
func (c *Client) RoomFiles(roomId string, params map[string]string) []File {
	ret := c.Get("/rooms/"+roomId+"/files", params)
	var files []File
	json.Unmarshal(ret, &files)
	return files
}

func (c *Client) RoomFile(roomId, fileId string) File {
	ret := c.Get("/rooms/"+roomId+"/files/"+fileId, map[string]string{})
	var file File
	json.Unmarshal(ret, &file)
	return file
}
