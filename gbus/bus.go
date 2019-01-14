package gbus

import "time"

//Bus interface provides the majority of functionality to Send, Reply and Publish messages to the Bus
type Bus interface {
	HandlerRegister
	BusSwitch
	Messaging
	SagaRegister
}

//Messaging interface to send and publish messages to the bus
type Messaging interface {
	/*
		Send a command or a command response to a specific service
		one-to-one semantics
	*/
	Send(toService string, command BusMessage) error

	/*
		Publish and event, one-to-many semantics
	*/
	Publish(exchange, topic string, event BusMessage) error
}

//BusSwitch starts and shutdowns the bus
type BusSwitch interface {
	/*
		Start starts the bus, once the bus is started messages get consiumed from the queue
		and handlers get invoced.
		Register all handlers prior to calling GBus.Start()
	*/
	Start() error
	/*
		Shutdown the bus and close connection to the underlying broker
	*/
	Shutdown()
}

//HandlerRegister registers message handlers to specific messages and events
type HandlerRegister interface {
	/*
		HandleMessage registers a handler to a specific message type
		Use this methof to register handlers for commands and reply messages
		Use the HandleEvent method to subscribe on events and registr a handler
	*/
	HandleMessage(message interface{}, handler MessageHandler) error
	/*
		HandleEvent registers a handler for a specific message type published
		to an exchange with a specific topic
	*/
	HandleEvent(exchange, topic string, event interface{}, handler MessageHandler) error
}

//MessageHandler signature for all command handlers
type MessageHandler func(invocation Invocation, message *BusMessage)

//Saga is the base interface for all Sagas.
type Saga interface {
	//StartedBy returns the messages that when received should create a new saga instance
	StartedBy() []interface{}
	/*
		RegisterAllHandlers passes in the HandlerRegister so that the saga can register
		the messages that it handles
	*/
	RegisterAllHandlers(register HandlerRegister)
}

//Timeout is the interface a saga needs to implement to get timeout servicess
type Timeout interface {
	RequestTimeout() time.Duration
	Timeout(invocation Invocation, message *BusMessage)
}

//SagaRegister registers sagas to the bus
type SagaRegister interface {
	RegisterSaga(saga Saga) error
}
