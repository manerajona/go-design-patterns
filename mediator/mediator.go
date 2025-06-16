package mediator

import "fmt"

type User struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

func NewUser(username string) *User {
	return &User{Name: username}
}

func (user *User) Receive(sender, message string) {
	entry := fmt.Sprintf("%s: '%s'", sender, message)
	user.chatLog = append(user.chatLog, entry)
}

func (user *User) Announce(message string) {
	if user.Room != nil {
		user.Room.Broadcast(user.Name, message)
	}
}

func (user *User) DirectMessage(username, message string) {
	if user.Room != nil {
		user.Room.Message(user.Name, username, message)
	}
}

type ChatRoom struct {
	users []*User
}

func (c *ChatRoom) Broadcast(username, message string) {
	for _, user := range c.users {
		if user.Name != username {
			user.Receive(username, message)
		}
	}
}

func (c *ChatRoom) Join(user *User) {
	joinMessage := user.Name + " joins the chat"
	c.Broadcast("Room", joinMessage)

	user.Room = c
	c.users = append(c.users, user)
}

func (c *ChatRoom) Message(fromUsername, toUsername, message string) {
	for _, user := range c.users {
		if user.Name == toUsername {
			user.Receive(fromUsername, message)
		}
	}
}
