// go run oving2_go.go
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

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// recvSock = 
	buffer := make([]byte, 1024)
	tcpAddr, err := ResolveTCPAddr("tcp", ":30000") // Mulig net. trengs foran Dial
	checkError(err)
	Println(1)
	listener, err := ListenTCP("tcp", tcpAddr)
	checkError(err)
	Println(2)
	//conn, err := listener.Accept()
	//checkError(err)
	//Println(3)
	for {
		conn, err := listener.AcceptTCP()
		checkError(err)
		Println(3)
		time.Sleep(1000*time.Millisecond)
		//Println("Hei!")
		n,err := conn.Read(buffer)
		checkError(err)
		Println(4)
		Printf("Rcv %d bytes: %s\n",n, buffer)
	}
	

}
