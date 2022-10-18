package abstractfactory

type Samsung struct {
}

func (a *Samsung) makeComputer() IComputer {
	return &SamsungComputer{Computer{
		memory:    1024,
		processor: "Samsung CPU",
	}}
}

func (a *Samsung) makeNotebook() INotebook {
	return &SamsungNotebook{Notebook{
		memory:    2048,
		processor: "Samsung CPU New",
		screen:    "15",
	}}
}
