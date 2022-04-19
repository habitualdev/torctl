package circuit

import (
	"fmt"
	"net"
	"strings"
	"torctl/collection"
)

func UpdateCircuit(c collection.Connection) string {
	temp := make([]byte, 1024)
	var response string

	_, err := c.Conn.Write([]byte("SIGNAL ACTIVE\r\n"))
	if err != nil {
		return "Error: " + err.Error()
	}
	_, err = c.Conn.Read(temp)
	if err != nil {
		return ""
	}

	response = string(temp) + "\n"

	_, err = c.Conn.Write([]byte("SIGNAL NEWNYM\r\n"))
	if err != nil {
		return "Error: " + err.Error()
	}
	_, err = c.Conn.Read(temp)
	if err != nil {
		return ""
	}

	response += string(temp) + "\n"

	_, err = c.Conn.Write([]byte("getinfo circuit-status\r\n"))
	if err != nil {
		return "Error: " + err.Error()
	}
	_, err = c.Conn.Read(temp)
	if err != nil {
		return ""
	}

	response += string(temp) + "\n"

	return response
}

func GetStatus(c collection.Connection) string {
	temp := make([]byte, 1024)
	var response string

	_, err := c.Conn.Write([]byte("getinfo circuit-status\r\n"))
	if err != nil {
		return "Error: " + err.Error()
	}
	_, err = c.Conn.Read(temp)
	if err != nil {
		return ""
	}

	response = string(temp) + "\n"

	return response
}

func GetConfig(conn net.Conn) string {
	data := make([]byte, 1024)
	conn.Write([]byte("getinfo config-text\n"))
	_, err := conn.Read(data)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return string(data)
}

func GetTraffic(conn net.Conn) string {
	data := make([]byte, 1024)
	var response string
	conn.Write([]byte("getinfo traffic/read\n"))
	_, err := conn.Read(data)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	response = "bytes read/bytes written: " + strings.Split(strings.Split(string(data), "=")[1], "\n")[0] + "/"

	conn.Write([]byte("getinfo traffic/written\n"))
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	response += strings.Split(strings.Split(string(data), "=")[1], "\n")[0] + "\n"

	return response
}
