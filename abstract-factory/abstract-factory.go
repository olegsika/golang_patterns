package abstractfactory

import "fmt"

func Example() {
	fmt.Printf("Abstract Factory Start \n")
	appleFactory, _ := GetTechFactory(AppleCompanyName)
	samsungFactory, _ := GetTechFactory(SamsungCompanyName)

	appleComputer := appleFactory.makeComputer()
	appleNotebook := appleFactory.makeNotebook()

	samsungComputer := samsungFactory.makeComputer()
	samsungNotebook := samsungFactory.makeNotebook()

	printComputer(appleComputer)
	printComputer(samsungComputer)

	printNotebook(appleNotebook)
	printNotebook(samsungNotebook)
	fmt.Printf("Abstract Factory Finish \n")
}

func printComputer(c IComputer) {
	fmt.Printf("Memory: %d, Processor: %s \n", c.getMemory(), c.getProcessor())
}

func printNotebook(n INotebook) {
	fmt.Printf("Memory: %d, Processor: %s, Screen: %s \n", n.getMemory(), n.getProcessor(), n.getScreen())
}
