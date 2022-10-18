package abstractfactory

type Notebook struct {
	memory    int
	processor string
	screen    string
}

func (c *Notebook) setMemory(memory int) {
	c.memory = memory
}

func (c *Notebook) setProcessor(processor string) {
	c.processor = processor
}

func (c *Notebook) setScreen(screen string) {
	c.screen = screen
}

func (c *Notebook) getMemory() int {
	return c.memory
}

func (c *Notebook) getProcessor() string {
	return c.processor
}

func (c *Notebook) getScreen() string {
	return c.screen
}
