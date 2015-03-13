package control

import ( 
	//"fmt"
	//"udp"
	"driver"
	//"control"
)

func GoToFloor(button int, floor int) {
	driver.SetButtonLamp(button,floor,1)
	if(driver.GetFloorSensorSignal() == -1) {
		driver.SetMotorDirection(driver.DIRN_UP)
		}
	for {
		driver.SetFloorIndicator(driver.GetFloorSensorSignal())		
		if floor == driver.GetFloorSensorSignal() {
			driver.SetMotorDirection(driver.DIRN_STOP)
			driver.SetFloorIndicator(floor)
			driver.SetButtonLamp(button,floor,0)
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

func GetDestination() (int,int) {
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
			for button := 0; button <= 2; button++ {
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
