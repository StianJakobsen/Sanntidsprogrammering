// go run networkUDP.go
package main
import (. "fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	"runtime"
	"time"
	."net"
	"bufio"
	"os"
)

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err) //err.Error()
		return //os.exit(1)
	}
	
}


//func broadcast(data []byte, conn *UDPConn) { // "129.241.187.255:
//	udpAddr, err := ResolveUDPAddr("udp", "129.241.187.255:20001")
//	checkError(err)
//	conn, err := DialUDP("udp", nil, udpAddr)
//	checkError(err)
//	for {
//		conn.Write(data) // conn.Write([]byte("Hvem er der?\n"))
//	
//	}
//}

func listen() {
	buffer := make([]byte, 1024)
	udpAddr, err := ResolveUDPAddr("udp", ":32222")
	conn, err := ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		Println("Hører")
		n,err := conn.Read(buffer)
		checkError(err)
		Printf("Rcv %d bytes: %s\n",n, buffer)
	}	
}


func send(ip []byte) { // data []byte
	udpAddr, err := ResolveUDPAddr("udp", string(ip[:21]))
	checkError(err)
	conn, err := DialUDP("udp", nil, udpAddr)
	checkError(err)
	for {
		Println("SENDER")
		//buffer = nil
		time.Sleep(1000*time.Millisecond)
		
		// WRITE
		//Println("Er du der server??")
		_, err := conn.Write([]byte("fetbmwazz\n")) // \x00
		checkError(err)
	}

}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does... 
	//buffer := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)
	Print("IP til mottaker:")
	ip, _ := reader.ReadBytes('\n')
	go send(ip)
	go listen()
	
	deadChan := make(chan int)
	<-deadChan
	

		for {
		
		//Print(ip)
		//Print("Melding:")
		//msg, _ := reader.ReadString('\n')
		


	}












}
