package builder

type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildComputer() Computer {
	d.builder.setRam()
	d.builder.setGPU()
	d.builder.setSSD()
	d.builder.setCPU()
	return d.builder.getComputer()
}
