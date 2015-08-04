package gochatwork

import (
	"encoding/json"
	"time"
)

const BaseURL = `https://api.chatwork.com/v1`

type Me struct {
	AccountID        int    `json:"account_id"`
	RoomID           int    `json:"room_id"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	Title            string `json:"title"`
	URL              string `json:"url"`
	Introduction     string `json:"introduction"`
	Mail             string `json:"mail"`
	TelOrganization  string `json:"tel_organization"`
	TelExtension     string `json:"tel_extension"`
	TelMobile        string `json:"tel_mobile"`
	Skype            string `json:"skype"`
	Facebook         string `json:"facebook"`
	Twitter          string `json:"twitter"`
	AvatarImageURL   string `json:"avatar_image_url"`
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
	AccountID        int    `json:"account_id"`
	RoomID           int    `json:"room_id"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageURL   string `json:"avatar_image_url"`
}

func (c *Client) Contacts() []Contact {
	ret := c.Get("/contacts", map[string]string{})
	var contacts []Contact
	json.Unmarshal(ret, &contacts)
	return contacts
}

type Room struct {
	RoomID         int    `json:"room_id"`
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

func (c *Client) Room(roomID string) Room {
	ret := c.Get("/rooms/"+roomID, map[string]string{})
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
func (c *Client) UpdateRoom(roomID string, params map[string]string) []byte {
	return c.Put("/rooms/"+roomID, params)
}

// params key
//   * action_type: [leave, delete]
func (c *Client) DeleteRoom(roomID string, params map[string]string) []byte {
	return c.Delete("/rooms/"+roomID, params)
}

type Member struct {
	AccountID        int    `json:"account_id"`
	Role             string `json:"role"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageURL   string `json:"avatar_image_url"`
}

func (c *Client) RoomMembers(roomID string) []Member {
	ret := c.Get("/rooms/"+roomID+"/members", map[string]string{})
	var members []Member
	json.Unmarshal(ret, &members)
	return members
}

// params keys
//   * members_admin_ids
//   - members_member_ids
//   - members_readonly_ids
func (c *Client) UpdateRoomMembers(roomID string, params map[string]string) []byte {
	return c.Put("/rooms/"+roomID+"/members", params)
}

type Account struct {
	AccountID      int    `json:"account_id"`
	Name           string `json:"name"`
	AvatarImageURL string `json:"avatar_image_url"`
}

type Message struct {
	MessageID  int     `json:"message_id"`
	Account    Account `json:"account"`
	Body       string  `json:"body"`
	SendTime   int64   `json:"send_time"`
	UpdateTime int64   `json:"update_time"`
}

func (m Message) SendDate() time.Time {
	return time.Unix(m.SendTime, 0)
}

func (m Message) UpdateDate() time.Time {
	return time.Unix(m.UpdateTime, 0)
}

type Messages []Message

func (c *Client) RoomMessages(roomID string, params map[string]string) Messages {
	ret := c.Get("/rooms/"+roomID+"/messages", params)
	var msgs Messages
	json.Unmarshal(ret, &msgs)
	return msgs
}

func (c *Client) PostRoomMessage(roomID string, body string) []byte {
	return c.Post("/rooms/"+roomID+"/messages", map[string]string{"body": body})
}

func (c *Client) RoomMessage(roomID, messageID string) Message {
	ret := c.Get("/rooms/"+roomID+"/messages/"+messageID, map[string]string{})
	var message Message
	json.Unmarshal(ret, &message)
	return message
}

type Task struct {
	TaskID            int     `json:"task_id"`
	Account           Account `json:"account"`
	AssignedByAccount Account `json:"assigned_by_account"`
	MessageID         int     `json:"message_id"`
	Body              string  `json:"body"`
	LimitTime         int64   `json:"limit_time"`
	Status            string  `json:"status"`
}

func (c *Client) RoomTasks(roomID string) []Task {
	ret := c.Get("/rooms/"+roomID+"/tasks", map[string]string{})
	var tasks []Task
	json.Unmarshal(ret, &tasks)
	return tasks
}

// params keys
//   * body
//   * to_ids
//   - limit
func (c *Client) PostRoomTask(roomID string, params map[string]string) []byte {
	return c.Post("/rooms/"+roomID+"/tasks", params)
}

func (c *Client) RoomTask(roomID, taskID string) Task {
	ret := c.Get("/rooms/"+roomID+"/tasks/"+taskID, map[string]string{})
	var task Task
	json.Unmarshal(ret, &task)
	return task
}

type File struct {
	FileID      int     `json:"file_id"`
	Account     Account `json:"account"`
	MessageID   int     `json:"message_id"`
	Filename    string  `json:"filename"`
	Filesize    int     `json:"filesize"`
	UploadTime  int64   `json:"upload_time"`
	DownloadURL string  `json:"download_url"`
}

// params key
//   - account_id
func (c *Client) RoomFiles(roomID string, params map[string]string) []File {
	ret := c.Get("/rooms/"+roomID+"/files", params)
	var files []File
	json.Unmarshal(ret, &files)
	return files
}

func (c *Client) RoomFile(roomID, fileID string, params map[string]string) File {
	ret := c.Get("/rooms/"+roomID+"/files/"+fileID, params)
	var file File
	json.Unmarshal(ret, &file)
	return file
}
