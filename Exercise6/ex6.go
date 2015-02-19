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
)


type TellerStruct struct {
		teller int
	}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	//buffer := make([]byte, 1024)
	currentStruct := TellerStruct{0}

	

	//current := 0
	tellerChan := make(chan int)
	primaryChan := make(chan int) // kanskje bruke til å passe på at eg er primary
	backupChan := make(chan int)



	//tellerChan <- currentStruct.teller
	//Println(<-tellerChan)
	//backupChan <- 1
	//primaryChan <- 1
	
	go listenForPrimary(tellerChan, backupChan, primaryChan, &currentStruct)

	for {
		//Println("forloopmain")
		select {
			case <-primaryChan:
				Println("Start Alive Broadcast")
				//tellerChan<- currentStruct.teller
				go imAlive(currentStruct.teller)
			case <-backupChan:
				//Println("hei fra primary state")
				// Vente x antall sekund
				// Hvis listen for primary ikkje får melding
				// --> I AM PRIMARY!
				//currentStruct.teller = <-tellerChan
				Println("Siste tall motatt fra primary: ", currentStruct.teller)
		}
	}
}


func imAlive(teller int) { // Bare sende siste tal for simplicity
	udpAddr, err := ResolveUDPAddr("udp", "192.168.145.255:30169") // Broadcast (rektig ip?)
	checkError(err)
	conn, err := DialUDP("udp", nil, udpAddr)
	checkError(err)
	currentStruct := TellerStruct{teller}
	//Println("imalive?")
	//currentStruct.teller = <-tellerChan
	for {
		b,_ := json.Marshal(currentStruct)
		conn.Write(b)	
		Println("Sent: ",currentStruct.teller) 		
		currentStruct.teller = currentStruct.teller + 1
		time.Sleep(1*time.Second)
	}
}

func listenForPrimary(tellerChan chan int, backupChan chan int, primaryChan chan int, 
currentStruct *TellerStruct) {
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
			//Println((*currentStruct).teller)
			//tellerChan <- (*currentStruct).teller
			primaryChan<- 1
			//time.Sleep(1*time.Second)
			break
		}
		err = json.Unmarshal(buffer[0:n], currentStruct)
		checkError(err)
		//tellerChan<- (*currentStruct).teller
		backupChan<- 1
	}

}

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err) //err.Error()
		return                           //os.exit(1)
	}
}
