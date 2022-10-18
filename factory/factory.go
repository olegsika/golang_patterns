package factory

import "fmt"

func Example() {
	fmt.Printf("Factory Start \n")
	apple, _ := getComputer(AppleCompanyName)
	samsung, _ := getComputer(SamsungCompanyName)

	printComputer(apple)
	printComputer(samsung)
	fmt.Printf("Factory Finish \n")

}

func printComputer(c IComputer) {
	fmt.Printf("Memory: %d, Processor: %s \n", c.getMemory(), c.getProcessor())
}

func getComputer(company string) (IComputer, error) {
	switch company {
	case AppleCompanyName:
		return newAppleComputer(), nil
	case SamsungCompanyName:
		return newSamsungComputer(), nil
	}

	return nil, fmt.Errorf("wrong factory passed")
}
