package main

import ( 
	"fmt"
	"udp"
	"driver"
)

func main() {
	if !initElevator()) {
		fmt.Println("Unable to initialize elevator hardware!")
		return
	}
	
	fmt.Println("Press STOP button to stop elevator and exit program.")
	
	driver.SetMotorDirection(DIRN_UP)
	
	for {
		if driver.GetFloorSensorSignal() == N_FLOORS - 1 {
			driver.SetMotorDirection(DIRN_DOWN)
		} else if driver.GetFloorSensorSignal() == 0 {
			driver.SetMotorDirection(DIRN_UP)
		}
		if driver.GetStopSignal() {
			driver.SetMotorDirection(DIRN_STOP)
			break
		}
	}
}		 
