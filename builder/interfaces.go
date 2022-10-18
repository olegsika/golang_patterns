package builder

type Builder interface {
	setRam()
	setGPU()
	setSSD()
	setCPU()
	getComputer() Computer
}
