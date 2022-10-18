package decorator

import "fmt"

func Example() {
	computer := &Computer{}

	computerWithGPU := &AddGPU{computer: computer}
	computerWithSSD := &AddSSD{computer: computerWithGPU}

	fmt.Printf("The price of computer with ssd and gpu %d, the power %d \n", computerWithSSD.getPrice(), computerWithSSD.getPower())
}
