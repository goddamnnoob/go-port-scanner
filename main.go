package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	now := time.Now()
	wg.Add(1024)
	hostname := "192.168.1.1"
	protocol := "tcp"
	for i := 1; i < 1025; i++ {
		go Scanport(hostname, protocol, i, &wg)
	}
	wg.Wait()
	fmt.Println("Time taken to scan 1024 ports: ", time.Since(now))

}

func Scanport(hostname, protocol string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := hostname + ":" + strconv.Itoa(port)
	con, err := net.Dial(protocol, address)
	if con != nil {
		defer con.Close()
	}
	if err != nil {
		return
	}
	fmt.Println("Port ", port, " is open")
	return
}
