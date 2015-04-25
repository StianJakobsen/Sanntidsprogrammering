//Sanntidsprogrammering!!
package main

import ( 
	"fmt"
	"udp"
	"driver"
	"control"
	"runtime" 
	//"net"
	//"os"
	//"sort"
	"functions"
)

func main() {
	fmt.Println("FINN ET BEDRE STED FOR RUNNING=0 I GÅR")
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(udp.GetID())
		
	//floorChan := make(chan int)
	var Data udp.Data
	
	dataIn, dataOut := make(chan *udp.Data), make(chan *udp.Data)
	//statusIn, statusOut := make(chan *udp.Status), make(chan *udp.Status)
	PrimaryChan := make(chan int)
	SlaveChan := make(chan int)
	SortChan := make(chan int)
	
	dataIn <- Data
	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
	return
	}
	udp.UdpInit(30169, 39998, 1024, Data, dataIn, dataOut, PrimaryChan,SlaveChan, SortChan)
	
	if(Data.Statuses[udp.GetIndex(udp.GetID(),Data)].CurrentFloor == -1){
		control.GoToFloor(0,Data.Statuses[udp.GetIndex(udp.GetID(),Data)],0)	
	}
	fmt.Println("Ferdig med å initialisere")
	fmt.Println("Currentfloor: ", Data.Statuses[udp.GetIndex(udp.GetID(),&Data)].CurrentFloor)
	fmt.Println("GetINDEX: ",udp.GetIndex(udp.GetID(),Data))
	//fmt.Println("Currentfloor: ", Data.Statuses[udp.GetIndex(udp.GetID(),&Data)].CurrentFloor)
	fmt.Println("Test: ", udp.GetIndex(udp.GetID(),Data))
	fmt.Println("Currentfloor[0]: ", Data.Statuses[0].CurrentFloor)
	//Status.ID = udp.GetID()	
	fmt.Println("Getfloor", driver.GetFloorSensorSignal())	


		
	//PrimaryChan<- 1
	//SlaveChan<-1
	fmt.Println("MIN INDEX ER: ", udp.GetIndex(udp.GetID(),Data))
	go control.GetDestination(Data.Statuses[udp.GetIndex(udp.GetID(),Data)])
	go control.ElevatorControl(Data.Statuses[udp.GetIndex(udp.GetID(), Data)]) //statusIn, statusOut)
	fmt.Println("index fra main: ", udp.GetIndex(udp.GetID(), Data))
	if Data.Statuses[udp.GetIndex(udp.GetID(), Data)].Primary {
		dataIn<- Data
		fmt.Println("Setter igang PrimaryListen og Costfunction")
		go udp.PrimaryListen(dataIn, dataOut, SortChan)
		go control.CostFunction(dataIn, dataOut)
	}

	for {
		fmt.Println("for loop")
		select {
			case <-PrimaryChan:
				Data.Statuses[udp.GetIndex(udp.GetID(), Data)].Primary = true
				go control.CostFunction(dataIn, dataOut) 
			case <-SlaveChan:
				
			case <-SortChan:
				if len(Data.PrimaryQ)  > 1{
					temp := functions.SortUp(Data.PrimaryQ[1:])
					Data.PrimaryQ = Data.PrimaryQ[:1]
					Data.PrimaryQ = append(Data.PrimaryQ, temp...)
					fmt.Println(Data.PrimaryQ)
				}
			case temp := <-dataOut:
				fmt.Println("Er i main og har tatt imot fra????")
				dataIn<- temp
				//statusIn<- &data.Statuses[udp.GetIndex(udp.GetID(), &Data)]
				//dataIn<-
			//case <-statusOut	
			//default:
				//fmt.Println("default case")
		}
	}
	

	
	
	
	
	fmt.Println("Press STOP button to stop elevator and exit program.")
	
	//if Status.Primary == true {
	//	go udp.Send()
	//} else {
	//	go udp.Listen()
	//}	
		
	//go control.GoToFloor(2,floorChan,&Data)
	
	/*
	for {
		//_, temp := control.GetCommand()
		//floorChan<- temp
		//PrintStatus(Data.Status)
		fmt.Println("Stop signal pressed ", driver.GetStopSignal())
		if driver.GetStopSignal() != 0 {
			fmt.Println("Stop signal pressed ", driver.GetStopSignal())			
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	
	}
	*/
}		 

func PrintStatus(Status udp.Status) {
	fmt.Println("Running: ", Status.Running)
	fmt.Println("CurrentFloor: ", Status.CurrentFloor)
	fmt.Println("NextFloor: ", Status.NextFloor)
	fmt.Println("Primary: ", Status.Primary)
	fmt.Println("ID: ", Status.ID)
	fmt.Println("OrderList: ", Status.OrderList)
}
