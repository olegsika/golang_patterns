package mediatorchatexample

import "errors"

type chat struct {
	Users []User
}

func NewChat() Chat {
	return &chat{}
}

func (c *chat) Add(user User) {
	if user.ID == 0 {
		return
	}

	c.Users = append(c.Users, user)
}

func (c *chat) Say(userID int, message string) error {
	user, i := c.GetUser(userID)

	if user == nil {
		return errors.New("the user not in chat \n")
	}

	c.Users[i].Messages = append(c.Users[i].Messages, message)

	return nil
}

func (c *chat) GetUser(userID int) (*User, int) {
	for i, u := range c.Users {
		if u.ID == userID {
			return &u, i
		}
	}

	return nil, -1
}

func (c *chat) SayAll(message string) {
	for _, user := range c.Users {
		user.Messages = append(user.Messages, message)
	}
}

func (c *chat) GetUsers() []User {
	return c.Users
}
