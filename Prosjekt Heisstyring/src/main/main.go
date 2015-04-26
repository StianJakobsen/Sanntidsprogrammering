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
	//"functions"
"time"
)



func main() {
	fmt.Println("FINN ET BEDRE STED FOR RUNNING=0 I GÅR")
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(udp.GetID())
	
	//floorChan := make(chan int)
	var data udp.Data
	costIn, costOut := make(chan *udp.Data), make(chan *udp.Data)
	primListenIn, primListenOut := make(chan *udp.Data), make(chan *udp.Data)
	slaveUpdateIn, slaveUpdateOut := make(chan *udp.Data), make(chan *udp.Data)
	slaveListenIn, slaveListenOut := make(chan *udp.Data), make(chan *udp.Data)
	//statusIn, statusOut := make(chan *udp.Status), make(chan *udp.Status)
	PrimaryChan := make(chan int)
	SlaveChan := make(chan int)
	//SortChan := make(chan int)
	
	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
	return
	}
	udp.UdpInit(30169, 39998, 1024, &data, slaveListenIn, slaveListenOut, PrimaryChan,SlaveChan)
	
	/*if(data.Statuses[udp.GetIndex(udp.GetID(), &data)].CurrentFloor == -1){
		control.GoToFloor(0,&data)	
	}*/
	fmt.Println("Ferdig med å initialisere")	

	fmt.Println("MIN INDEX ER: ", udp.GetIndex(udp.GetID(), &data))
	time.Sleep(1000*time.Millisecond)
	go control.GetDestination(&data)
	go control.ElevatorControl(&data) //statusIn, statusOut)
	go control.LampControl(&data)
	
	
	if data.Statuses[udp.GetIndex(udp.GetID(), &data)].Primary {
		fmt.Println("Setter igang PrimaryListen og Costfunction")
		go udp.PrimaryListen(primListenIn, primListenOut)
		go control.CostFunction(costIn, costOut)
		costIn <- &data
	}

	for {
		//fmt.Println("Uplist?: ", data.Statuses[0].UpList)
		select {
			case <-PrimaryChan:
				data.Statuses[udp.GetIndex(udp.GetID(), &data)].Primary = true
				go control.CostFunction(costIn, costOut) 
			//	go udp.PrimaryBroadcast(broadcastPort,&data)
				go udp.PrimaryListen(primListenIn,primListenOut)
				costIn <- &data
			case <-SlaveChan:
				go udp.SlaveUpdate(slaveUpdateIn, slaveUpdateOut)
				slaveUpdateIn <- &data
			/*	
			case <-SortChan: // passe på å omsortere Statuses og
				if len(data.PrimaryQ)  > 1{
					temp := functions.SortUp(data.PrimaryQ[1:])
					data.PrimaryQ = data.PrimaryQ[:1]
					data.PrimaryQ = append(data.PrimaryQ, temp...)
					fmt.Println(data.PrimaryQ)
				}
				fmt.Println("La til nytt element i PrimaryQ: ", data.PrimaryQ)
				//primListenIn <- data
			*/
			case <-costOut:
				//fmt.Println("COSTOUT")
				primListenIn <- &data
			
			case <-primListenOut:
				//fmt.Println("PRIMLISTENOUT")
				costIn <- &data
				
			case <-slaveListenOut:
				slaveUpdateIn <- &data 
			
			case <- slaveUpdateOut:
				slaveListenIn <- &data
			//case dataIn := <-dataOut:
			//	fmt.Println("Er i main og har tatt imot fra????")
				//dataIn<- temp
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
