package mediatorchatexample

type Chat interface {
	Add(User)
	Say(int, string) error
	SayAll(string)
	GetUsers() []User
	GetUser(int) (*User, int)
}
