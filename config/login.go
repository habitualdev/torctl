package config

import (
	"fmt"
	"net"
	"torctl/collection"
)

func PasswordAuth(t collection.Connection) (net.Conn, string) {
	data := make([]byte, 1024)
	conn, err := net.Dial("tcp", t.Auth.AuthHost+":"+t.Auth.AuthPort)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, ""
	}
	_, err = conn.Write([]byte("AUTHENTICATE \"" + t.Auth.AuthPass + "\"\n"))
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, ""
	}
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, ""
	}

	return conn, string(data)
}

func CookieAuth(t collection.Connection) (net.Conn, string) {
	data := make([]byte, 1024)
	conn, err := net.Dial("tcp", t.Auth.AuthHost+":"+t.Auth.AuthPort)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, ""
	}
	_, err = conn.Write([]byte("AUTHENTICATE " + t.Auth.AuthPass + "\n"))
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, ""
	}
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, ""
	}

	return conn, string(data)
}
