package mediator

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
