// listenTestUDP.go
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	buffer := make([]byte, 1024)
	udpAddr, err := ResolveUDPAddr("udp", ":30169")
	conn, err := ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("ERROR, %v", err)
		return
	}

	for {
		n, _ := conn.Read(buffer)
		fmt.Printf("Rcv %d bytes: %d\n", n, buffer)
	}

	fmt.Println("Hello World!")
}
