package control

import ( 
	"fmt"
	"udp"
	//"runtime"
	"time"
	"driver"
	//"control"
	//"os"
	"sort"
	
)

func GoToFloor(button int,  floorChan chan int,data *udp.Data) {
	floor := <-floorChan
	if driver.GetFloorSensorSignal() == -1 {
		driver.SetMotorDirection(driver.DIRN_DOWN)
	}
	var done int
	temp:= floor	
	//polse:
	for {	/*
		if driver.GetStopSignal() != 0 {
			driver.SetMotorDirection(driver.DIRN_STOP)
			fmt.Println("Stop button pressed")			
			os.Exit(1)	
			}*/	
		select {
		
		case temp = <-floorChan:
			//fmt.Println("Her er temp: %d", temp)
			//fmt.Println("Her er DONE: %d", done)
			//if done == 1{
									
			//	floor = temp
			//	done = 0
			//}

		default:			
			/*if driver.GetStopSignal() != 0 {
				driver.SetMotorDirection(driver.DIRN_STOP)
				os.Exit(1)	
			}		*/
			driver.SetFloorIndicator(driver.GetFloorSensorSignal())	
			if done == 1{
				//fmt.Printf("GAA IN EHFE")				
				floor = temp
				done = 0
				
				
			}	
			//fmt.Printf("Hva er done? %d\n",done)
			driver.SetButtonLamp(button,floor,1)
			//fmt.Printf("Her er flooooooooor: %d\n", floor)
				
			if floor == driver.GetFloorSensorSignal()  {
				
				fmt.Println("Framme på:", floor)
				udp.SetStatus(data,0,floor)

				driver.SetDoorOpenLamp(true)				
				driver.SetMotorDirection(driver.DIRN_STOP)
				time.Sleep(1*time.Second)
				driver.SetDoorOpenLamp(false)
				driver.SetFloorIndicator(floor)
				driver.SetButtonLamp(button,floor,0)
				done = 1
				
				//temp = -1
				//driver.SetDoorOpenLamp(false)	
				//fmt.Println("Done: %d", done)
				break
		
			} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1 {   
			udp.SetStatus(data,2, floor)
			driver.SetMotorDirection(driver.DIRN_UP) 
		
			} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1{
			udp.SetStatus(data ,1, floor)
			driver.SetMotorDirection(driver.DIRN_DOWN)
			}

		}
	
	}
}

func GetDestination(status *udp.Status) int { //returnerer bare button, orderlist oppdateres
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
				 if(driver.GetButtonSignal(0,floor) == 1 && status.UpList[len(status.UpList)-1] != floor) {
					status.UpList = append(status.UpList,floor)
					status.PriList = append(status.PriList,0)
					return 0
				} else if driver.GetButtonSignal(1,floor) == 1  && status.DownList[len(status.DownList)-1] != floor {
					status.DownList = append(status.DownList,floor)
					status.PriList = append(status.PriList,1)
					return 1 
				} else if driver.GetButtonSignal(2,floor) == 1  && status.CommandList[len(status.CommandList)-1] != floor {
					status.CommandList = append(status.CommandList,floor)
					status.PriList = append(status.PriList,2)
					return 2
				}
			
		}
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	}
return -1
}

func GetCommand() (int,int) {
	button := 2	
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
			if(driver.GetButtonSignal(button,floor) == 1) {
				return button,floor
			}
			
		}
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	}
return -1,-1
}


func CostFunction(data *udp.Data) {
	DownList := data.Statuses[data.PrimaryQ[0]].DownList
	for k := 1; k < len(data.PrimaryQ);k++ {
		DownList = append(DownList,data.Statuses[data.PrimaryQ[k]].DownList...)
	}
	UPList := data.Statuses[data.PrimaryQ[0]].UpList
	for l := 1; l < len(data.PrimaryQ);l++ {
		DownList = append(UpList,data.Statuses[data.PrimaryQ[l]].UpList...)
	}
	UpList = SortUp(UpList)
	DownList = SortUp(DownList)
	DownList = sort.Reverse(DownList)
	
	for i := 0; i < len(data.PrimaryQ); i++ {
		for down:=0; down<len(DownList);down++ {
			if DownList[down] == data.Statuses[data.PrimaryQ[i]].CurrentFloor {
				data.Statuses[data.PrimaryQ[i]].OrderList = DownList[down:]
				DownList = UpdateList(DownList,down)
				//pluss noe mer, som å åpne døra
			} else if data.Statuses[data.PrimaryQ[i]].CurrentFloor == DownList[down]+1 && data.Statuses[data.PrimaryQ[i]].Running == 1 {
				//StopAtsenddosomething....
			} else if data.Statuses[data.PrimaryQ[i]].CurrentFloor == DownList[down]+1 && data.Statuses[data.PrimaryQ[i]].Running == 0 {
				// Send the fucking elevator
			} else if data.Statuses[data.PrimaryQ[i]].CurrentFloor > DownList[down] && data.Statuses[data.PrimaryQ[i]].Running == 1 {
				// gs
			} else if data.Statuses[data.PrimaryQ[i]].CurrentFloor > DownList[down] && data.Statuses[data.PrimaryQ[i]].Running == 0 {
				// SNART I MÅL? :P
			} else {
				// ....
			}
		}
		for up:=0; up<len(UpList);up++ {
			if UpList[up] == data.Statuses[data.PrimaryQ[i]].CurrentFloor {
				UpList = UpdateList(UpList,up)
				//pluss noe mer, som å åpne døra
			} else if data.Statuses[data.PrimaryQ[i]].CurrentFloor == UpList[up]-1 && data.Statuses[data.PrimaryQ[i]].Running == 1 {
				
				//StopAtsenddosomething....
			} else if data.Statuses[data.PrimaryQ[i]].CurrentFloor == UpList[up]-1 && data.Statuses[data.PrimaryQ[i]].Running == 0 {
				// Send the fucking elevator
			} else if data.Statuses[data.PrimaryQ[i]].CurrentFloor < UpList[up] && data.Statuses[data.PrimaryQ[i]].Running == 1 {
				// gs
			} else if data.Statuses[data.PrimaryQ[i]].CurrentFloor < UpList[up] && data.Statuses[data.PrimaryQ[i]].Running == 0 {
				// SNART I MÅL? :P
			} else {
				// SEND THE CLOSEST FUCKING ELEVATOR
			}
		}
	}
		
		

}



func UpdateList(OrderList []int, j int) []int {
	temp := make([]int, len(OrderList)-1)
	for i:= 0; i<len(OrderList);i++ {
		if i<j {
			temp[i] = OrderList[i]
		} else if i>j {
			temp[i-1] = OrderList[i]
		}
	}
	return temp
}

func SortUp(UpList []int)  []int{
	sort.Ints(UpList)
	temp := make([]int,1)
	temp[0] = UpList[0]
	counter := 0
	for i:= 1;i<len(UpList); i++ {
		if UpList[i] > temp[counter] {
			counter ++
			temp = append(temp,UpList[i])
		}
	}
	return temp
}	
/*
func SortDown(DownList []int)  []int{
	sort.Ints(DownList)
	temp := make([]int,1)
	temp[0] = DownList[len(DownList)-1]
	counter := 0
	for i:= (len(DownList)-1); i>=0; i-- {
		
		if DownList[i] < temp[counter] {
			counter ++
			temp = append(temp,DownList[i])
		}
	}
	return temp
} */
	





