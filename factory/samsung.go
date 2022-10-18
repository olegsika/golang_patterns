package factory

type Samsung struct {
	Computer
}

func newSamsungComputer() IComputer {
	return &Samsung{
		Computer: Computer{
			memory:    2048,
			processor: "Samsung A1",
		}}
}
