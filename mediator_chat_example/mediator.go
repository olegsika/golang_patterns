package mediatorchatexample

import "fmt"

func Example() {
	fmt.Printf("Mediator Start \n")

	userOne := User{
		ID:   1,
		Name: "Oleh",
	}
	userTwo := User{
		ID:   2,
		Name: "John",
	}
	userThree := User{
		ID:   3,
		Name: "Bob",
	}

	c := NewChat()

	c.Add(userOne)
	c.Add(userTwo)

	if err := c.Say(userThree.ID, "Hello World"); err != nil {
		fmt.Printf("Error when the user three say. Err - %+v", err)
	}

	if err := c.Say(userTwo.ID, "I'm  John"); err != nil {
		fmt.Printf("Error when the user two say. Err - %+v", err)
	}

	c.SayAll("Hello All")

	fmt.Printf("Mediator Finish \n")
}
