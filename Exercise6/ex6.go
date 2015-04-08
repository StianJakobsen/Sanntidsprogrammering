// ex6
// go run ex6.go
package main

import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
	// This is probably not a good idea for large projects...
	. "net"
	"runtime"
	"time"
	"encoding/json"
	//"strconv"
)


type TellerStruct struct {
		Teller int
	}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...

	currentStruct := TellerStruct{0}
	primaryChan := make(chan int)
	backupChan := make(chan int)
	
	go listenForPrimary(backupChan, primaryChan, &currentStruct)

	for {
		select {
			case <-primaryChan:
				Println("Start Alive Broadcast")
				go imAlive(currentStruct.Teller)
			case <-backupChan:
				Println("Siste tall motatt fra primary: ", currentStruct.Teller)
		}
	}
}


func imAlive(teller int) { // Bare sende siste tal for simplicity
	udpAddr, err := ResolveUDPAddr("udp", "129.241.187.255:30169") // Broadcast (endre ip nettverket du sitter på)
	checkError(err)
	conn, err := DialUDP("udp", nil, udpAddr)
	checkError(err)
	currentStruct := TellerStruct{teller}

	for {
		b,_ := json.Marshal(currentStruct)
		conn.Write(b)	
		Println("Sent: ",currentStruct.Teller) 		
		currentStruct.Teller = currentStruct.Teller + 1
		time.Sleep(1*time.Second)
	}
}

func listenForPrimary(backupChan chan int, primaryChan chan int, currentStruct *TellerStruct) {
	buffer := make([]byte, 1024)
	udpAddr, err := ResolveUDPAddr("udp", ":30169")
	checkError(err)
	conn, err := ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		conn.SetReadDeadline(time.Now().Add(3*time.Second))
		n, err := conn.Read(buffer)
		if err != nil{
			Println("Tar over som primary!")
			primaryChan<- 1
			break
		}

		err = json.Unmarshal(buffer[0:n], currentStruct)
		checkError(err)
		backupChan<- 1
	}

}

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err)
		return
	}
}
