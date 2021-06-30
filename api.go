package chatwork

import (
	"encoding/json"
)

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

func (c *Client) Me() (*Me, error) {
	ret, err := c.Get("/me", &Params{})
	if err != nil {
		return nil, err
	}
	var me Me
	err = json.Unmarshal(ret, &me)
	return &me, err
}

type Status struct {
	UnreadRoomNum  int `json:"unread_room_num"`
	MentionRoomNum int `json:"mention_room_num"`
	MytaskRoomNum  int `json:"mytask_room_num"`
	UnreadNum      int `json:"unread_num"`
	MentionNum     int `json:"mention_num"`
	MyTaskNum      int `json:"mytask_num"`
}

func (c *Client) MyStatus() (*Status, error) {
	ret, err := c.Get("/my/status", &Params{})
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(ret, &status)
	return &status, err
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
func (c *Client) MyTasks(params *Params) ([]MyTask, error) {
	ret, err := c.Get("/my/tasks", params)
	if err != nil {
		return nil, err
	}
	var tasks []MyTask
	err = json.Unmarshal(ret, &tasks)
	return tasks, err
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

func (c *Client) Contacts() ([]Contact, error) {
	ret, err := c.Get("/contacts", &Params{})
	if err != nil {
		return nil, err
	}
	var contacts []Contact
	err = json.Unmarshal(ret, &contacts)
	return contacts, err
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

func (c *Client) Rooms() ([]Room, error) {
	ret, err := c.Get("/rooms", &Params{})
	if err != nil {
		return nil, err
	}
	var rooms []Room
	err = json.Unmarshal(ret, &rooms)
	return rooms, err
}

func (c *Client) Room(roomId string) (*Room, error) {
	ret, err := c.Get("/rooms/"+roomId, &Params{})
	if err != nil {
		return nil, err
	}
	var room Room
	err = json.Unmarshal(ret, &room)
	return &room, err
}

// params keys
//   * name
//   * members_admin_ids
//   - description
//   - icon_preset
//   - members_member_ids
//   - members_readonly_ids
func (c *Client) CreateRoom(params *Params) ([]byte, error) {
	return c.Post("/rooms", params)
}

// params keys
//   - description
//   - icon_preset
//   - name
func (c *Client) UpdateRoom(roomId string, params *Params) ([]byte, error) {
	return c.Put("/rooms/"+roomId, params)
}

// params key
//   * action_type: [leave, delete]
func (c *Client) DeleteRoom(roomId string, params *Params) ([]byte, error) {
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

func (c *Client) RoomMembers(roomId string) ([]Member, error) {
	ret, err := c.Get("/rooms/"+roomId+"/members", &Params{})
	if err != nil {
		return nil, err
	}
	var members []Member
	err = json.Unmarshal(ret, &members)
	return members, err
}

// params keys
//   * members_admin_ids
//   - members_member_ids
//   - members_readonly_ids
func (c *Client) UpdateRoomMembers(roomId string, params *Params) ([]byte, error) {
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

func (c *Client) RoomMessages(roomId string) ([]Message, error) {
	ret, err := c.Get("/rooms/"+roomId+"/messages", &Params{})
	if err != nil {
		return nil, err
	}
	var messages []Message
	err = json.Unmarshal(ret, &messages)
	return messages, err
}

func (c *Client) PostRoomMessage(roomId string, body string) ([]byte, error) {
	return c.Post("/rooms/"+roomId+"/messages", &Params{"body": body})
}

func (c *Client) RoomMessage(roomId, messageId string) (*Message, error) {
	ret, err := c.Get("/rooms/"+roomId+"/messages/"+messageId, &Params{})
	if err != nil {
		return nil, err
	}
	var message Message
	err = json.Unmarshal(ret, &message)
	return &message, err
}

type Task struct {
	TaskId            int     `json:"task_id"`
	Account           Account `json:"account"`
	AssignedByAccount Account `json:"assigned_by_account"`
	MessageId         string  `json:"message_id"`
	Body              string  `json:"body"`
	LimitTime         int64   `json:"limit_time"`
	Status            string  `json:"status"`
}

func (c *Client) RoomTasks(roomId string) ([]Task, error) {
	ret, err := c.Get("/rooms/"+roomId+"/tasks", &Params{})
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(ret, &tasks)
	return tasks, err
}

// params keys
//   * body
//   * to_ids
//   - limit
func (c *Client) PostRoomTask(roomId string, params *Params) ([]byte, error) {
	return c.Post("/rooms/"+roomId+"/tasks", params)
}

func (c *Client) RoomTask(roomId, taskId string) (*Task, error) {
	ret, err := c.Get("/rooms/"+roomId+"/tasks/"+taskId, &Params{})
	if err != nil {
		return nil, err
	}
	var task Task
	err = json.Unmarshal(ret, &task)
	return &task, err
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
func (c *Client) RoomFiles(roomId string, params *Params) ([]File, error) {
	ret, err := c.Get("/rooms/"+roomId+"/files", params)
	if err != nil {
		return nil, err
	}
	var files []File
	err = json.Unmarshal(ret, &files)
	return files, err
}

func (c *Client) RoomFile(roomId, fileId string) (*File, error) {
	ret, err := c.Get("/rooms/"+roomId+"/files/"+fileId, &Params{})
	if err != nil {
		return nil, err
	}
	var file File
	err = json.Unmarshal(ret, &file)
	return &file, err
}
