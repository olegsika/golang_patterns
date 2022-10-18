package abstractfactory

type Apple struct {
}

func (a *Apple) makeComputer() IComputer {
	return &AppleComputer{Computer{
		memory:    1024,
		processor: "M1",
	}}
}

func (a *Apple) makeNotebook() INotebook {
	return &AppleNotebook{Notebook{
		memory:    2048,
		processor: "M2",
		screen:    "13",
	}}
}
