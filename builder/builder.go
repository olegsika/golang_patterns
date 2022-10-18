package builder

import (
	"errors"
	"fmt"
)

func Example() {
	fmt.Printf("Builder Start \n")
	appleBuilder, _ := getBuilder(AppleBuilderName)
	samsungBuilder, _ := getBuilder(SamsungBuilderName)

	director := newDirector(appleBuilder)
	appleComputer := director.buildComputer()

	printComputer(AppleBuilderName, appleComputer)

	director.setBuilder(samsungBuilder)
	samsungComputer := director.buildComputer()

	printComputer(SamsungBuilderName, samsungComputer)
	fmt.Printf("Builder Finish \n")
}

func getBuilder(builderType string) (Builder, error) {
	switch builderType {
	case AppleBuilderName:
		return newAppleBuilder(), nil
	case SamsungBuilderName:
		return newSamsungBuilder(), nil
	}

	return nil, errors.New("wrong builder passed")
}

func printComputer(builderName string, computer Computer) {
	fmt.Printf("Builder %s, Ram %d, GPU %s, SSD %d, CPU %2.f \n", builderName, computer.RAM, computer.GPU, computer.SSD, computer.CPU)
}
