package factory

type IComputer interface {
	setMemory(memory int)
	setProcessor(processor string)
	getMemory() int
	getProcessor() string
}
