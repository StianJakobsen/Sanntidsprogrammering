//Sanntidsprogrammering!!
package main

import ( 
	"fmt"
	"udp"
	"driver"
	"control"
	"runtime" 
	"time"
)



func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(udp.GetID())
	var data udp.Data
	costIn, costOut := make(chan *udp.Data), make(chan *udp.Data)
	primListenIn, primListenOut := make(chan *udp.Data), make(chan *udp.Data)
	slaveUpdateIn, slaveUpdateOut := make(chan *udp.Data), make(chan *udp.Data)
	slaveListenIn, slaveListenOut := make(chan *udp.Data), make(chan *udp.Data)
	PrimaryChan := make(chan int)
	SlaveChan := make(chan int)
	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
	return
	}
	udp.UdpInit(39998, &data, slaveListenIn, slaveListenOut, PrimaryChan,SlaveChan)
	fmt.Println("Ferdig med å initialisere")	
	time.Sleep(1000*time.Millisecond)
	go control.GetDestination(&data)
	go control.ElevatorControl(&data) 
	go control.LampControl(&data)
	
	
	if data.Statuses[udp.GetIndex(udp.GetID(), &data)].Primary {
		fmt.Println("Setter igang PrimaryListen og Costfunction")
		go udp.PrimaryListen(primListenIn, primListenOut)
		go control.CostFunction(costIn, costOut)
		costIn <- &data
	}


	for {
		select {
			case <-PrimaryChan:
				data.Statuses[udp.GetIndex(udp.GetID(), &data)].Primary = true
				go control.CostFunction(costIn, costOut) 
				go udp.PrimaryListen(primListenIn,primListenOut)
				costIn <- &data
			case <-SlaveChan:
				go udp.SlaveUpdate(slaveUpdateIn, slaveUpdateOut)
				slaveUpdateIn <- &data

			case <-costOut:
				
				primListenIn <- &data
			
			case <-primListenOut:
				
				costIn <- &data
				
			case <-slaveListenOut:
				slaveUpdateIn <- &data 
			
			case <- slaveUpdateOut:
				slaveListenIn <- &data
				

		}
	}
	

	
	
	
	

}		 


