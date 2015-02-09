// go run udpRecieve.go
package main
import (. "fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	"runtime"
	"time"
	."net"
)

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err) //err.Error()
		return //os.exit(1)
	}
	
}

func read() {
	buffer := make([]byte, 1024)
	udpAddr, err := ResolveUDPAddr("udp4", ":20001")
	conn, err := ListenUDP("udp4", udpAddr)
	checkError(err)
	for {
		n,err := conn.Read(buffer)
		checkError(err)
		Printf("Rcv %d bytes: %s\n",n, buffer)
		
	}	
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// recvSock = 
	//buffer := make([]byte, 1024)
	udpAddr, err := ResolveUDPAddr("udp4", "129.241.187.145:31111")
	checkError(err)
	Println("1")
	conn, err := DialUDP("udp4", nil, udpAddr)
	checkError(err)
	Println("2")
	
	go read()
	
	//Println(conn2)
	//os.Exit(1)
	for {
		//buffer = nil
		time.Sleep(1000*time.Millisecond)
		Println("Hei!")
		
		// WRITE
		Println("Er du der server??")
		_, err := conn.Write([]byte("Eg er her\n")) // \x00
		//conn.Write([]byte("Eg er her\n"))
		checkError(err)
		
				
		// READ
		//n,err := conn.Read(buffer)
		//checkError(err)
		//Printf("Rcv %d bytes: %s\n",n, buffer)
		

	}
	
	
	


}
