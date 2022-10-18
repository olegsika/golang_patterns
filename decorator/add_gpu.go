package decorator

type AddGPU struct {
	computer IComputer
}

func (a *AddGPU) getPrice() int {
	computerPrice := a.computer.getPrice()
	return computerPrice + 10
}

func (a *AddGPU) getPower() int {
	computerPower := a.computer.getPower()
	return computerPower + 20
}
