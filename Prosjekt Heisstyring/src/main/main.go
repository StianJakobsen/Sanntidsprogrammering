package main

import ( 
	"fmt"
	//"udp"
	"driver"
)

func main() {
	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
		return
	}
	
	fmt.Println("Press STOP button to stop elevator and exit program.")
	
	driver.SetMotorDirection(driver.DIRN_UP)
	
	for {
		if driver.GetFloorSensorSignal() == driver.N_FLOORS - 1 {
			driver.SetMotorDirection(driver.DIRN_DOWN)
		} else if driver.GetFloorSensorSignal() == 0 {
			driver.SetMotorDirection(driver.DIRN_UP)
		}
		if driver.GetStopSignal() != 0 {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	driver.SetFloorIndicator(driver.GetFloorSensorSignal())
	}
}		 
