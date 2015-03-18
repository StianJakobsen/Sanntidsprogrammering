package control

import ( 
	"fmt"
	//"udp"
	//"runtime"
	"time"
	"driver"
	//"control"
)

func GoToFloor(button int, floorChan chan int) {
	floor := <-floorChan
	
	var done int
	temp:= floor	
	if(driver.GetFloorSensorSignal() == -1) {
		driver.SetMotorDirection(driver.DIRN_DOWN)
		}
	for {
		select {
		
		case temp = <-floorChan:
			//fmt.Println("Her er temp: %d", temp)
			//fmt.Println("Her er DONE: %d", done)
			//if done == 1{
									
			//	floor = temp
			//	done = 0
			//}

		default:			
						
			driver.SetFloorIndicator(driver.GetFloorSensorSignal())	
			if done == 1{
				fmt.Printf("GAA IN EHFE")				
				floor = temp
				done = 0
			}	
			fmt.Printf("Hva er done? %d",done)
			driver.SetButtonLamp(button,floor,1)
			fmt.Println("Her er flooooooooor: %d", floor)			
			if floor == driver.GetFloorSensorSignal()  {
				driver.SetDoorOpenLamp(true)				
				driver.SetMotorDirection(driver.DIRN_STOP)
				time.Sleep(2*time.Second)
				driver.SetDoorOpenLamp(false)
				driver.SetFloorIndicator(floor)
				driver.SetButtonLamp(button,floor,0)
				done = 1
				//temp = -1
				//driver.SetDoorOpenLamp(false)	
				fmt.Println("Done: %d", done)
				break
		
			} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 {
			driver.SetMotorDirection(driver.DIRN_UP) 
		
			} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 {
			driver.SetMotorDirection(driver.DIRN_DOWN)
			}
			if driver.GetStopSignal() != 0 {
				driver.SetMotorDirection(driver.DIRN_STOP)
				break
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
return 666,666
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
return 666,666
}








