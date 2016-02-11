package gochatwork

import (
	"encoding/json"
	"time"
)

// BaseURL ChatWork API endpooint URL
const BaseURL = `https://api.chatwork.com/v1`

// Me model
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

// Me GET "/me"
func (c *Client) Me() (Me, error) {
	ret, err := c.Get("/me", map[string]string{})
	var me Me
	json.Unmarshal(ret, &me)
	return me, err
}

// Status model
type Status struct {
	UnreadRoomNum  int `json:"unread_room_num"`
	MentionRoomNum int `json:"mention_room_num"`
	MytaskRoomNum  int `json:"mytask_room_num"`
	UnreadNum      int `json:"unread_num"`
	MentionNum     int `json:"mention_num"`
	MyTaskNum      int `json:"mytask_num"`
}

// MyStatus GET "/my/status"
func (c *Client) MyStatus() (Status, error) {
	ret, err := c.Get("/my/status", map[string]string{})
	var status Status
	json.Unmarshal(ret, &status)
	return status, err
}

// MyTask model
type MyTask struct {
	Task
	Room struct {
		Roomid   int    `json:"room_id"`
		Name     string `json:"name"`
		IconPath string `json:"icon_path"`
	}
}

// MyTasks GET "/my/tasks"
// params keys
//  - assigned_by_account_id
//  - status: [open, done]
func (c *Client) MyTasks(params map[string]string) ([]MyTask, error) {
	ret, err := c.Get("/my/tasks", params)
	var tasks []MyTask
	json.Unmarshal(ret, &tasks)
	return tasks, err
}

// Contact model
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

// Contacts GET "/contacts"
func (c *Client) Contacts() ([]Contact, error) {
	ret, err := c.Get("/contacts", map[string]string{})
	var contacts []Contact
	json.Unmarshal(ret, &contacts)
	return contacts, err
}

// Room model
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
	Description    string `json:"description"`
}

// Rooms GET "/rooms"
func (c *Client) Rooms() ([]Room, error) {
	ret, err := c.Get("/rooms", map[string]string{})
	var rooms []Room
	json.Unmarshal(ret, &rooms)
	return rooms, err
}

// Room GET "/rooms/{room_id}"
func (c *Client) Room(roomID string) (Room, error) {
	ret, err := c.Get("/rooms/"+roomID, map[string]string{})
	var room Room
	json.Unmarshal(ret, &room)
	return room, err
}

// CreateRoom POST "/rooms"
// params keys
//   * name
//   * members_admin_ids
//   - description
//   - icon_preset
//   - members_member_ids
//   - members_readonly_ids
func (c *Client) CreateRoom(params map[string]string) ([]byte, error) {
	return c.Post("/rooms", params)
}

// UpdateRoom PUT "/rooms/{room_id}"
// params keys
//   - description
//   - icon_preset
//   - name
func (c *Client) UpdateRoom(roomID string, params map[string]string) ([]byte, error) {
	return c.Put("/rooms/"+roomID, params)
}

// DeleteRoom DELETE "/rooms/{room_id}"
// params key
//   * action_type: [leave, delete]
func (c *Client) DeleteRoom(roomID string, params map[string]string) ([]byte, error) {
	return c.Delete("/rooms/"+roomID, params)
}

// Member model
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

// RoomMembers GET "/rooms/{room_id}/members"
func (c *Client) RoomMembers(roomID string) ([]Member, error) {
	ret, err := c.Get("/rooms/"+roomID+"/members", map[string]string{})
	var members []Member
	json.Unmarshal(ret, &members)
	return members, err
}

// UpdateRoomMembers PUT "/rooms/{room_id}/members"
// params keys
//   * members_admin_ids
//   - members_member_ids
//   - members_readonly_ids
func (c *Client) UpdateRoomMembers(roomID string, params map[string]string) ([]byte, error) {
	return c.Put("/rooms/"+roomID+"/members", params)
}

// Account model
type Account struct {
	AccountID      int    `json:"account_id"`
	Name           string `json:"name"`
	AvatarImageURL string `json:"avatar_image_url"`
}

// Message model
type Message struct {
	MessageID  int     `json:"message_id"`
	Account    Account `json:"account"`
	Body       string  `json:"body"`
	SendTime   int64   `json:"send_time"`
	UpdateTime int64   `json:"update_time"`
}

// SendDate time.Time representation of SendTime
func (m Message) SendDate() time.Time {
	return time.Unix(m.SendTime, 0)
}

// UpdateDate time.Time representation of UpdateTime
func (m Message) UpdateDate() time.Time {
	return time.Unix(m.UpdateTime, 0)
}

// Messages slice of Message
type Messages []Message

// RoomMessages GET "/rooms/{room_id}/messages"
func (c *Client) RoomMessages(roomID string, params map[string]string) (Messages, error) {
	ret, err := c.Get("/rooms/"+roomID+"/messages", params)
	var msgs Messages
	json.Unmarshal(ret, &msgs)
	return msgs, err
}

// PostRoomMessage POST "/rooms/{room_id}/messages"
func (c *Client) PostRoomMessage(roomID string, body string) ([]byte, error) {
	return c.Post("/rooms/"+roomID+"/messages", map[string]string{"body": body})
}

// RoomMessage GET "/rooms/{room_id}/messages/{message_id}"
func (c *Client) RoomMessage(roomID, messageID string) (Message, error) {
	ret, err := c.Get("/rooms/"+roomID+"/messages/"+messageID, map[string]string{})
	var message Message
	json.Unmarshal(ret, &message)
	return message, err
}

// Task model
type Task struct {
	TaskID            int     `json:"task_id"`
	Account           Account `json:"account"`
	AssignedByAccount Account `json:"assigned_by_account"`
	MessageID         int     `json:"message_id"`
	Body              string  `json:"body"`
	LimitTime         int64   `json:"limit_time"`
	Status            string  `json:"status"`
}

// LimitDate time.Time representation of LimitTime
func (t Task) LimitDate() time.Time {
	return time.Unix(t.LimitTime, 0)
}

// RoomTasks GET "/rooms/{room_id}/tasks"
// params keys
//  - account_id
//  - assigned_by_account_id
//  - status: [open, done]
func (c *Client) RoomTasks(roomID string, params map[string]string) ([]Task, error) {
	ret, err := c.Get("/rooms/"+roomID+"/tasks", params)
	var tasks []Task
	json.Unmarshal(ret, &tasks)
	return tasks, err
}

// PostRoomTask POST "/rooms/{room_id}/tasks"
// params keys
//   * body
//   * to_ids
//   - limit
func (c *Client) PostRoomTask(roomID string, params map[string]string) ([]byte, error) {
	return c.Post("/rooms/"+roomID+"/tasks", params)
}

// RoomTask GET "/rooms/{room_id}/tasks/tasks_id"
func (c *Client) RoomTask(roomID, taskID string) (Task, error) {
	ret, err := c.Get("/rooms/"+roomID+"/tasks/"+taskID, map[string]string{})
	var task Task
	json.Unmarshal(ret, &task)
	return task, err
}

// File model
type File struct {
	FileID      int     `json:"file_id"`
	Account     Account `json:"account"`
	MessageID   int     `json:"message_id"`
	Filename    string  `json:"filename"`
	Filesize    int     `json:"filesize"`
	UploadTime  int64   `json:"upload_time"`
	DownloadURL string  `json:"download_url"`
}

// UploadDate time.Time representation of UploadTime
func (f File) UploadDate() time.Time {
	return time.Unix(f.UploadTime, 0)
}

// RoomFiles GET "/rooms/{room_id}/files/"
// params key
//   - account_id
func (c *Client) RoomFiles(roomID string, params map[string]string) ([]File, error) {
	ret, err := c.Get("/rooms/"+roomID+"/files", params)
	var files []File
	json.Unmarshal(ret, &files)
	return files, err
}

// RoomFile GET "/rooms/{room_id}/files/{file_id}"
// params key
//   - create_download_url: ["0", "1"]
func (c *Client) RoomFile(roomID, fileID string, params map[string]string) (File, error) {
	ret, err := c.Get("/rooms/"+roomID+"/files/"+fileID, params)
	var file File
	json.Unmarshal(ret, &file)
	return file, err
}

// RateLimit model
type RateLimit struct {
	Limit     int
	Remaining int
	ResetTime int64
}

// ResetDate time.Time representation of ResetTime
func (r RateLimit) ResetDate() time.Time {
	return time.Unix(r.ResetTime, 0)
}

func (c *Client) RateLimit() *RateLimit {
	if c.latestRateLimit == nil {
		// When API is not called even once, call API and get RateLimit in response header
		c.Me()
	}
	return c.latestRateLimit
}
