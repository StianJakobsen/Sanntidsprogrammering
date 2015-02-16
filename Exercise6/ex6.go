// ex6
// go run ex6.go
package main

import ( 
	."fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	"runtime"
	"time"
	."net"
	"bufio"
	"os"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does... 
	//buffer := make([]byte, 1024)
	current := 0
	tellerChan := make(chan int)
	primaryChan := make(chan int) // kanskje bruke til å passe på at eg er primary
	backupChan := make(chan int)
	
	backupChan<- 1
	
	// go send(ip)
	// go listen()
	
	for {
		if primary() {
			select {
				case 
				
			}	
		}
		else {
			select {
				case
				
			}
		}
	}
	
	// eller:
	
	for {
		select {
			case <-primaryChan:
				Println(current)
				current++
				primaryChan<- 1
			
			case <-backupChan:
				
			
		}
		
	}
	
	
	
	// deadChan := make(chan int) //trurkje eg treng detta nå
	// <-deadChan
	

	
		
	fmt.Println("Hello World!")
}

func counter(tellerChan chan<- int) {
	for i := 0; ; i++ {
		Println(i)
		
	}
}


func imAlive() {
	udpAddr, err := ResolveUDPAddr("udp", "129.241.187.255:30169") // Broadcast (rektig ip?)
	checkError(err)
	conn, err := DialUDP("udp", nil, udpAddr)
	checkError(err)		
	for {
		time.Sleep(1*time.Second)
		conn.Write([]byte("I am alive\n")) // sende current tall å kanskje?
	}
}


func listenForPrimary() {
	buffer := make([]byte, 1024)
	udpAddr, err := ResolveUDPAddr("udp", ":32222")
	conn, err := ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		n, err := conn.Read(buffer)
		checkError(err)
		
		//Printf("Rcv %d bytes: %s\n",n, buffer)
	}	
}


func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err) //err.Error()
		return //os.exit(1)
	}
	
}