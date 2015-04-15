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
	
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(udp.GetID())	
	floorChan := make(chan int)
	//var Status udp.Status
	var Data udp.Data	
	//Data := make(map[int]udp.Status)
	//var PrimaryQ [3]string
	temp:=make([]int,5)
	temp[0] =3
	temp[1] =2
	temp[2] =1
	temp[3] =2
	temp[4] =1
	fmt.Println(temp)
	temp = append(temp,temp...)
	fmt.Println(temp)
	udp.UdpInit(30169, 39998, 1024, &Data)
	//Status.ID = udp.GetID()	
	fmt.Println("Getfloor", driver.GetFloorSensorSignal())	


	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
		return
	}
		
	


	fmt.Println("Press STOP button to stop elevator and exit program.")
	
	//if Status.Primary == true {
	//	go udp.Send()
	//} else {
	//	go udp.Listen()
	//}	
		
	go control.GoToFloor(2,floorChan,&Data)
	
	for {
		_, temp := control.GetCommand()
		floorChan<- temp
		//PrintStatus(Data.Status)
		fmt.Println("Stop signal pressed ", driver.GetStopSignal())
		if driver.GetStopSignal() != 0 {
			fmt.Println("Stop signal pressed ", driver.GetStopSignal())			
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
