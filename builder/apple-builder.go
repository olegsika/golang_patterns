package builder

type AppleBuilder struct {
	ram int
	gpu string
	ssd int
	cpu float64
}

func newAppleBuilder() *AppleBuilder {
	return &AppleBuilder{}
}

func (a *AppleBuilder) setRam() {
	a.ram = 2048
}
func (a *AppleBuilder) setGPU() {
	a.gpu = "nvidia GeForce"
}
func (a *AppleBuilder) setSSD() {
	a.ssd = 2048
}
func (a *AppleBuilder) setCPU() {
	a.cpu = 2.2
}

func (a *AppleBuilder) getComputer() Computer {
	return Computer{
		RAM: a.ram,
		GPU: a.gpu,
		SSD: a.ssd,
		CPU: a.cpu,
	}
}
