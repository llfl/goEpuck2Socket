package main

import (
	driver "github.com/llfl/E-Puck2-Golang/driverUtils"
	"fmt"
	"net"
	"time"
)

const (
	ipAddr = "192.168.1.123"
	destPort = "8888"
)

func main()  {

	url := ipAddr + ":" + destPort
	for{
		conn, err := net.Dial("tcp", url)
		if err != nil {
			fmt.Printf("Fail to connect, %s\n Retry in 5 sec", err)
			time.Sleep(5 * time.Second)
			continue
		}
		connHandler(conn)
	}
}

func connHandler(c net.Conn) {
	defer c.Close()
	epuck := driver.NewEPuckHandle()
	defer epuck.Device.Close()
	defer epuck.Stop()
	c.Write([]byte("10"))
	buf := make([]byte, 1024)
	for{
		c.Write([]byte("epuck"))
		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			continue
		}
		msg := string(buf[0:cnt])
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
		}
		time.Sleep(200 * time.Millisecond)
	}
}