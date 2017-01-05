package gochatwork

import (
	"encoding/json"
)

const BaseUrl = `https://api.chatwork.com/v1`

type Me struct {
	AccountId        int    `json:"account_id"`
	RoomId           int    `json:"room_id"`
	Name             string `json:"name"`
	ChatworkId       string `json:"chatwork_id"`
	OrganizationId   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	Title            string `json:"title"`
	Url              string `json:"url"`
	Introduction     string `json:"introduction"`
	Mail             string `json:"mail"`
	TelOrganization  string `json:"tel_organization"`
	TelExtension     string `json:"tel_extension"`
	TelMobile        string `json:"tel_mobile"`
	Skype            string `json:"skype"`
	Facebook         string `json:"facebook"`
	Twitter          string `json:"twitter"`
	AvatarImageUrl   string `json:"avatar_image_url"`
}

func (c *Client) Me() Me {
	ret := c.Get("/me", map[string]string{})
	var me Me
	json.Unmarshal(ret, &me)
	return me
}

type Status struct {
	UnreadRoomNum  int `json:"unread_room_num"`
	MentionRoomNum int `json:"mention_room_num"`
	MytaskRoomNum  int `json:"mytask_room_num"`
	UnreadNum      int `json:"unread_num"`
	MentionNum     int `json:"mention_num"`
	MyTaskNum      int `json:"mytask_num"`
}

func (c *Client) MyStatus() Status {
	ret := c.Get("/my/status", map[string]string{})
	var status Status
	json.Unmarshal(ret, &status)
	return status
}

type MyTask struct {
	Task
	Room struct {
		Roomid   int    `json:"room_id"`
		Name     string `json:"name"`
		IconPath string `json:"icon_path"`
	}
}

// params keys
//  - assigned_by_account_id
//  - status: [open, done]
func (c *Client) MyTasks(params map[string]string) []MyTask {
	ret := c.Get("/my/tasks", params)
	var tasks []MyTask
	json.Unmarshal(ret, &tasks)
	return tasks
}

type Contact struct {
	AccountId        int    `json:"account_id"`
	RoomId           int    `json:"room_id"`
	Name             string `json:"name"`
	ChatworkId       string `json:"chatwork_id"`
	OrganizationId   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageUrl   string `json:"avatar_image_url"`
}

func (c *Client) Contacts() []Contact {
	ret := c.Get("/contacts", map[string]string{})
	var contacts []Contact
	json.Unmarshal(ret, &contacts)
	return contacts
}

type Room struct {
	RoomId         int    `json:"room_id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Role           string `json:"role"`
	Sticky         bool   `json:"sticky"`
	UnreadNum      int    `json:"unread_num"`
	MentionNum     int    `json:"mention_num"`
	MytaskNum      int    `json:"mytask_num"`
	MessageNum     int    `json:"message_num"`
	FileNum        int    `json:"file_num"`
	TaskNum        int    `json:"task_num"`
	IconPath       string `json:"icon_path"`
	LastUpdateTime int64  `json:"last_update_time"`
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
//   * action_type: [leave, delete]
func (c *Client) DeleteRoom(roomId string, params map[string]string) []byte {
	return c.Delete("/rooms/"+roomId, params)
}

type Member struct {
	AccountId         int    `json:"account_id"`
	Role              string `json:"role"`
	Name              string `json:"name"`
	ChatworkId        string `json:"chatwork_id"`
	Organization_Id   int    `json:"organization_id"`
	Organization_Name string `json:"organization_name"`
	Department        string `json:"department"`
	AvatarImageUrl    string `json:"avatar_image_url"`
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
	AccountId      int    `json:"account_id"`
	Name           string `json:"name"`
	AvatarImageUrl string `json:"avatar_image_url"`
}

type Message struct {
	MessageId  int     `json:"message_id"`
	Account    Account `json:"account"`
	Body       string  `json:"body"`
	SendTime   int64   `json:"send_time"`
	UpdateTime int64   `json:"update_time"`
}

func (c *Client) RoomMessages(roomId string) []Message {
	ret := c.Get("/rooms/"+roomId+"/messages", map[string]string{})
	var messages []Message
	json.Unmarshal(ret, &messages)
	return messages
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
	TaskId            int     `json:"task_id"`
	Account           Account `json:"account"`
	AssignedByAccount Account `json:"assigned_by_account"`
	MessageId         int     `json:"message_id"`
	Body              string  `json:"body"`
	LimitTime         int64   `json:"limit_time"`
	Status            string  `json:"status"`
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
	FileId     int     `json:"file_id"`
	Account    Account `json:"account"`
	MessageId  int     `json:"message_id"`
	Filename   string  `json:"filename"`
	Filesize   int     `json:"filesize"`
	UploadTime int64   `json:"upload_time"`
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
