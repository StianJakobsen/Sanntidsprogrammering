// listenTestUDP.go
package main

import (
	"fmt"
	"runtime"
	"encoding/json"
	"net"
)

type TellerStruct struct{
	teller int
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	buffer := make([]byte, 1024)
	udpAddr, err := net.ResolveUDPAddr("udp", ":30169")
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("ERROR, %v", err)
		return
	}
	
	currentStruct := TellerStruct{0}
	
	for {
		n, _ := conn.Read(buffer)
		err = json.Unmarshal(buffer[0:n],&currentStruct)
		if err != nil {
			fmt.Printf("Noe gikk galt %v", err)
			return
		} 
		fmt.Printf("Rcv %d bytes: %d\n", n, currentStruct.teller)
	}

	//fmt.Println("Hello World!")
}
