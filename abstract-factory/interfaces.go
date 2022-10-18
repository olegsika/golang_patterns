package abstractfactory

type IComputer interface {
	setMemory(memory int)
	setProcessor(processor string)
	getMemory() int
	getProcessor() string
}

type INotebook interface {
	setMemory(memory int)
	setProcessor(processor string)
	setScreen(screen string)
	getMemory() int
	getProcessor() string
	getScreen() string
}

type ITechFactory interface {
	makeComputer() IComputer
	makeNotebook() INotebook
}
