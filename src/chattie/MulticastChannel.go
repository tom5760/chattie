package main

import (
	"encoding/json"
	"net"
)

type MulticastChannel struct {
	conn  *net.UDPConn
	group net.Addr

	decoder *json.Decoder
}

func NewMulticastChannel(address string) (*MulticastChannel, error) {
	var c MulticastChannel
	var err error

	udpAddress, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	c.conn, err = net.ListenMulticastUDP("udp", nil, udpAddress)
	if err != nil {
		return nil, err
	}

	c.group = udpAddress
	c.decoder = json.NewDecoder(c.conn)

	return &c, nil
}

func (c *MulticastChannel) Close() {
	c.conn.Close()
}

func (c *MulticastChannel) Join() error {
	return nil
}

func (c *MulticastChannel) Leave() error {
	return nil
}

func (c *MulticastChannel) Send(message *Message) error {
	// Can't use json.Encoder because we need to use PacketConn.WriteTo().
	buffer, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = c.conn.WriteTo(buffer, c.group)
	if err != nil {
		return err
	}

	return nil
}

func (c *MulticastChannel) Receive() (*Message, error) {
	var m Message
	err := c.decoder.Decode(&m)
	return &m, err
}
