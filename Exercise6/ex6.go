// ex6
// go run ex6.go
package main

import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
	// This is probably not a good idea for large projects...
	"bufio"
	"encoding/binary"
	. "net"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	//buffer := make([]byte, 1024)
	var current int64 = 0
	//current := 0
	tellerChan := make(chan int64)
	primaryChan := make(chan int) // kanskje bruke til å passe på at eg er primary
	backupChan := make(chan int)

	tellerChan <- current
	backupChan <- 1

	go listenForPrimary(tellerChan, backupChan)

	for {
		select {
		case <-primaryChan:
			Println(current)
			current++
			tellerChan <- current
			primaryChan <- 1
			time.Sleep(1 * time.Second)

		case <-backupChan:
			// Vente x antall sekund
			// Hvis listen for primary ikkje får melding
			// --> I AM PRIMARY!

		}

	}

	// deadChan := make(chan int) //trurkje eg treng detta nå
	// <-deadChan

}

//func counter(tellerChan chan int64) {
//	for i := 0; ; i++ {
//		Println(i)
//	}
//}

func imAlive(tellerChan <-chan int64) { // Bare sende siste tal for simplicity
	udpAddr, err := ResolveUDPAddr("udp", "129.241.187.255:30169") // Broadcast (rektig ip?)
	checkError(err)
	conn, err := DialUDP("udp", nil, udpAddr)
	checkError(err)
	for {
		//time.Sleep(1*time.Second)
		conn.Write([]byte(<-tellerChan))
		//conn.Write([]byte("I am alive\n")) // sende current tall å kanskje?
	}
}

func listenForPrimary(tellerChan chan<- int64, backupChan chan<- int) {
	buffer := make([]byte, 8)
	udpAddr, err := ResolveUDPAddr("udp", ":30169")
	conn, err := ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		n, err := conn.Read(buffer)
		checkError(err)
		current, _ := binary.Varint(buffer)
		tellerChan <- current
		backupChan <- 1
		//Printf("Rcv %d bytes: %s\n",n, buffer)
	}
}

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err) //err.Error()
		return                           //os.exit(1)
	}
}
