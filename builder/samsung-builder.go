package builder

type SamsungBuilder struct {
	ram int
	gpu string
	ssd int
	cpu float64
}

func newSamsungBuilder() *SamsungBuilder {
	return &SamsungBuilder{}
}

func (a *SamsungBuilder) setRam() {
	a.ram = 1024
}
func (a *SamsungBuilder) setGPU() {
	a.gpu = "AMD radeon"
}
func (a *SamsungBuilder) setSSD() {
	a.ssd = 1024
}
func (a *SamsungBuilder) setCPU() {
	a.cpu = 1.1
}

func (a *SamsungBuilder) getComputer() Computer {
	return Computer{
		RAM: a.ram,
		GPU: a.gpu,
		SSD: a.ssd,
		CPU: a.cpu,
	}
}
