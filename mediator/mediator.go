package mediator

import "fmt"

func Example() {
	fmt.Printf("Mediator 2 Start \n")
	stationManager := newStationManger()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
	fmt.Printf("Mediator 2 Finish \n")
}
