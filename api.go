package gochatwork

func (c *Client) Me() string {
        return c.Get("/me", map[string]string{})
}

func (c *Client) MyStatus() string {
        return c.Get("/my/status", map[string]string{})
}

// params keys
//  - assigned_by_account_id
//  - status: [open, done]
func (c *Client) MyTasks(params map[string]string) string {
        return c.Get("/my/tasks", params)
}

func (c *Client) Contacts() string {
        return c.Get("/contacts", map[string]string{})
}

func (c *Client) Rooms() string {
        return c.Get("/rooms", map[string]string{})
}

func (c *Client) Room(roomId string) string {
        return c.Get("/rooms/" + roomId, map[string]string{})
}

func (c *Client) RoomMembers(roomId string) string {
        return c.Get("/rooms/" + roomId + "/members", map[string]string{})
}

// XXX: Not yet implement
func (c *Client) RoomMessages(roomId string) string {
        return c.Get("/rooms/" + roomId + "/messages", map[string]string{})
}

func (c *Client) RoomMessage(roomId, messageId string) string {
        return c.Get("/rooms/" + roomId + "/messages/" + messageId, map[string]string{})
}

func (c *Client) RoomTasks(roomId string) string {
        return c.Get("/rooms/" + roomId + "/tasks", map[string]string{})
}

func (c *Client) RoomTask(roomId, taskId string) string {
        return c.Get("/rooms/" + roomId + "/tasks/" + taskId, map[string]string{})
}

// params key
//   - account_id
func (c *Client) RoomFiles(roomId string, params map[string]string) string {
        return c.Get("/rooms/" + roomId + "/files", params)
}

func (c *Client) RoomFile(roomId, fileId string) string {
        return c.Get("/rooms/" + roomId + "/files/" + fileId, map[string]string{})
}
