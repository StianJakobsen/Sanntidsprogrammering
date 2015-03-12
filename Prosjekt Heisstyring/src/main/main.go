package main

import ( 
	"fmt"
	//"udp"
	"driver"
	"control"
)

func main() {
	if driver.InitElevator() == 0 {
		fmt.Println("Unable to initialize elevator hardware!")
		return
	}
	
	fmt.Println("Press STOP button to stop elevator and exit program.")
	

	
	for {
		test1, test2 := control.GetCommand()
		control.GoToFloor(test1,test2)
	
		
		if driver.GetStopSignal() != 0 {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	
	}
}		 
