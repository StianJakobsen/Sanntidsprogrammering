// listenTestUDP.go
package main

import (
	"fmt"
	"runtime"
	"encoding/json"
	"net"
)

type TellerStruct struct{
	Teller int
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	buffer := make([]byte, 8)
	udpAddr, err := net.ResolveUDPAddr("udp", ":30169")
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("ERROR, %v", err)
		return
	}
	
	currentStruct := TellerStruct{Teller: 1}
	
	for {
		n, _ := conn.Read(buffer)
		err = json.Unmarshal(buffer[:n], &currentStruct)
		if err != nil {
			fmt.Printf("Noe gikk galt %v", err)
			return
		} 
		fmt.Printf("Rcv %d bytes: %d\n", n, currentStruct.Teller)
	}

	//fmt.Println("Hello World!")
}
