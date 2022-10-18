package abstractfactory

type Computer struct {
	memory    int
	processor string
}

func (c *Computer) setMemory(memory int) {
	c.memory = memory
}

func (c *Computer) setProcessor(processor string) {
	c.processor = processor
}

func (c *Computer) getMemory() int {
	return c.memory
}

func (c *Computer) getProcessor() string {
	return c.processor
}
