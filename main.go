package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	hostname := "192.168.1.1"
	protocol := "tcp"
	var b bool
	for i := 1; i < 1025; i++ {
		b = Scanport(hostname, protocol, i)
		if b {
			fmt.Println("Port ", i, " is open")
		} else {
			fmt.Println("Port ", i, " is closed")
		}
	}
	fmt.Println("Time taken to scan 1024 ports: ", time.Since(now))

}

func Scanport(hostname, protocol string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	con, err := net.DialTimeout(protocol, address, 10*time.Second)
	defer con.Close()
	if err != nil {
		return true
	}
	return false
}
