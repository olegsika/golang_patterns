package main

import (
	abstractFactory "patterns/abstract-factory"
	"patterns/bridge"
	"patterns/builder"
	"patterns/decorator"
	"patterns/factory"
	"patterns/mediator"
	mediatorChatExample "patterns/mediator_chat_example"
)

func main() {
	abstractFactory.Example()
	bridge.Example()
	builder.Example()
	decorator.Example()
	factory.Example()
	mediator.Example()
	mediatorChatExample.Example()

}
