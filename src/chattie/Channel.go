package main

type Channel interface {
	Close()

	Join() error
	Leave() error

	Send(message *Message) error
	Receive() (*Message, error)
}
