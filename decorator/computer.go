package decorator

type Computer struct {
}

func (a *Computer) getPrice() int {
	return 10
}

func (a *Computer) getPower() int {
	return 10
}
