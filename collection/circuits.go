package collection

import "net"

type Connection struct {
	Conn  net.Conn
	Auth  ConnectionAuth
	Exits Exit
}

type ConnectionAuth struct {
	AuthHost string
	AuthPort string
	AuthType string
	AuthPass string
}

type Exit struct {
	V4 string
	V6 string
}
