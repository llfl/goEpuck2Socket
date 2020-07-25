package main

import (
	driver "github.com/llfl/E-Puck2-Golang/driverUtils"
	"fmt"
	"net"
	"time"
)

const (
	listenPort = "8888"
)

var epuck *driver.EPuckHandle

func main()  {

	
	server, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		return
	}
	epuck = driver.NewEPuckHandle()
	defer epuck.Device.Close()
	defer epuck.Stop()

	for{
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Fail to connect, %s\n Retry in 5 sec", err)
			continue
		}
		connHandler(conn)
	}
}

func connHandler(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 1024)
	for{
		c.Write([]byte("r"))
		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			continue
		}
		msg := string(buf[0:cnt])
		c.Write([]byte("ok"))
		switch msg {
		case "spin_right":
			epuck.FreeSpin(64)
		case "spin_left":
			epuck.FreeSpin(-64)
		case "move_forward":
			epuck.Forward(64)
		case "move_backward":
			epuck.Forward(-64)
		case "stop":
			epuck.Stop()
		case "bye":
			break;
		}
		time.Sleep(200 * time.Millisecond)
	}
}