package decorator

type AddSSD struct {
	computer IComputer
}

func (a *AddSSD) getPrice() int {
	computerPrice := a.computer.getPrice()
	return computerPrice + 20
}

func (a *AddSSD) getPower() int {
	computerPower := a.computer.getPower()
	return computerPower + 20
}
