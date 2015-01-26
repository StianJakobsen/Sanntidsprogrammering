// go run oving2_go.go
package main
import (. "fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	"runtime"
	"time"
	"net" 
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// recvSock = 
	buffer := make([]byte, 1024)
	conn, err := Dial("udp", "serverip:34933") // Mulig net. trengs foran Dial
	if err != nil {
		Printf("Noe gikk galt %v", err)
		return
	}
	
	
	
time.Sleep(100*time.Millisecond)

}
