package mediatorchatexample

import (
	"errors"
	"patterns/tools"
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {
	tCases := []struct {
		name string
		user User
	}{
		{
			name: "Should be added to chat",
			user: User{
				ID:   1,
				Name: "Test",
			},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			c := NewChat()
			c.Add(tc.user)

			users := c.GetUsers()

			if !UserInChat(users, tc.user) {
				tt.Errorf("The users does not added to chat")
				return
			}
		})
	}
}

func TestSay(t *testing.T) {
	tCases := []struct {
		name             string
		expectedErr      error
		user             User
		message          string
		expectedMessages []string
	}{
		{
			name: "should be success",
			user: User{
				ID:   1,
				Name: "Test",
			},
			message:          "Hello, I'm test",
			expectedMessages: []string{"Hello, I'm test"},
		},
		{
			name:        "should be error. The user not added",
			expectedErr: errors.New("the user not in chat"),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			c := NewChat()

			c.Add(tc.user)

			err := c.Say(tc.user.ID, tc.message)

			if !tools.IsEqual(err, tc.expectedErr) {
				tt.Errorf("expected err = %v but actual err = %v", tc.expectedErr, err)
				return
			}

			user, _ := c.GetUser(tc.user.ID)
			if user != nil {
				if !reflect.DeepEqual(user.Messages, tc.expectedMessages) {
					tt.Errorf("expected messages = %v but actual messages = %v", user.Messages, tc.expectedMessages)
					return
				}
			}
		})
	}
}

func TestSayAll(t *testing.T) {
	tCases := []struct {
		name             string
		users            []User
		message          string
		expectedMessages []string
	}{
		{
			name: "should be success",
			users: []User{
				{
					ID:   1,
					Name: "test 1",
				},
				{
					ID:   2,
					Name: "test 2",
				},
			},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			c := NewChat()

			for _, u := range tc.users {
				c.Add(u)
			}

			c.SayAll(tc.message)

			for _, u := range tc.users {
				user, _ := c.GetUser(u.ID)

				if user != nil {
					if !reflect.DeepEqual(user.Messages, tc.expectedMessages) {
						tt.Errorf("expected messages = %v but actual messages = %v", user.Messages, tc.expectedMessages)
						return
					}
				}
			}
		})
	}
}

func UserInChat(users []User, user User) bool {
	for _, u := range users {
		if u.ID == user.ID {
			return true
		}
	}

	return false
}
