// go run oving2_go.go
package main
import (
. "fmt" // Using '.' to avoid prefixing functions with their package names
// This is probably not a good idea for large projects...
"runtime"
"time"
)

var i = 0
func opp(message chan int, doneChan chan bool) {
	var k int;
	for (k) =0;k<1000000;k++{
		//i = i+1;
		message <- 1
	}
	doneChan <- True
}

func ned(message chan int, doneChan chan bool) {
	var k int;
	for k =1000000;k>0;k--{
		//i = i-1;
		message <- 0
	}
	doneChan <- True
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	messages := make(chan int)
	doneChan := make(chan bool)
	//messages2 := make(chan int)
	//var i = 0;
	// Try doing the exercise both with and without it!
	go ned(messages);
	go opp(messages);

for{
	select{
		case c := <- messages:
			if (c == 1){
				i = i + 1;	
			}else if (c == 0){
				i = i - 1;
			}
			<- messages
		//case d := <- messages2:
			//fmt.Println(d)
				

	}
}




 // This spawns someGoroutine() as a goroutine
// We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
// We'll come back to using channels in Exercise 2. For now: Sleep.
time.Sleep(100*time.Millisecond)
Println("Her er i: ");
Println(i);

}
