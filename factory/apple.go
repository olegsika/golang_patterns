package factory

type Apple struct {
	Computer
}

func newAppleComputer() IComputer {
	return &Apple{
		Computer: Computer{
			memory:    512,
			processor: "Apple M1",
		}}
}
