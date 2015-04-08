package main

import ( 
	"fmt"
	"udp"
	"driver"
	"control"
	"runtime" 
	//"net"
	//"os"
)

func main() {
	
	fmt.Println(udp.GetID()*10)
	runtime.GOMAXPROCS(runtime.NumCPU())
	floorChan := make(chan int)
	var Status udp.Status
	
	udp.UdpInit(30169, 30170, 1024, &Status)
	//Status.ID = udp.GetID()	
	fmt.Println("Getfloor", driver.GetFloorSensorSignal())	
	PrintStatus(Status)
	
	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
		return
	}
		

	fmt.Println("Press STOP button to stop elevator and exit program.")
	
	go control.GoToFloor(2,floorChan,&Status)
	
	for {
		_, temp := control.GetCommand()
		floorChan<- temp
		PrintStatus(Status)
		
		if driver.GetStopSignal() != 0 {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	
	}
}		 

func PrintStatus(Status udp.Status) {
	fmt.Println("Running: ", Status.Running)
	fmt.Println("CurrentFloor: ", Status.CurrentFloor)
	fmt.Println("NextFloor: ", Status.NextFloor)
	fmt.Println("Primary: ", Status.Primary)
	fmt.Println("ID: ", Status.ID)
}
