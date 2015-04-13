package control

import ( 
	"fmt"
	"udp"
	//"runtime"
	"time"
	"driver"
	//"control"
	//"os"
)

func GoToFloor(button int, floorChan chan int,data *udp.Data) {
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
		
			} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 {
			udp.SetStatus(data,2, floor)
			driver.SetMotorDirection(driver.DIRN_UP) 
		
			} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 {
			udp.SetStatus(data ,1, floor)
			driver.SetMotorDirection(driver.DIRN_DOWN)
			}

		}
	
	}
}

func GetDestination() (int,int) {
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
			for button := 0; button < 2; button++ {
				 if(driver.GetButtonSignal(button,floor) == 1) {
					return button,floor
				}
			}
		}
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	}
return -1,-1
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

/*
func CostFunction(StatusList *Status, OrderList []int) {
	for all elevators {
		if elevator == same floor as command {
			do nothing/open door
		}
	}
	
*/




