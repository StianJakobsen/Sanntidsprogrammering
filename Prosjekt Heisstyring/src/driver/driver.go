// driver

package driver // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.c and driver.go
/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"

const N_FLOORS = 4

//Lights
const upLights := [N_FLOORS]int{LIGHT_UP1, LIGHT_UP2, LIGHT_UP3, LIGHT_UP4}
const downLights := [N_FLOORS]int{LIGHT_DOWN1, LIGHT_DOWN2, LIGHT_DOWN3, LIGHT_DOWN4}
const cmdLights := [N_FLOORS]int{LIGHT_COMMAND1,LIGHT_COMMAND2, LIGHT_COMMAND3, LIGHT_COMMAND4}

//Buttons
const upButtons := [N_FLOORS]int{BUTTON_UP1, BUTTON_UP2, BUTTON_UP3, BUTTON_UP4}
const downButtons := [N_FLOORS]int{BUTTON_DOWN1, BUTTON_DOWN2, BUTTON_DOWN3, BUTTON_DOWN4}
const cmdButtons := [N_FLOORS]int{BUTTON_COMMAND1, BUTTON_COMMAND2, BUTTON_COMMAND3, BUTTON_COMMAND4}

func initElevator() {
	if !C.io_init(){
		return 0
	}
	for i := 0; i < N_FLOORS; i++ {
		if i != 0{
			C.io_clear_bit(downlights[i])
		} 		
		if i != N_FLOORS - 1 {
			C.io_clear_bit(upLights[i])
		}
		C.io_clear_bit(cmdLights[i])
	} 
	// + noko greier her
}

func SetDoorOpenLamp(value bool) {

	if value {
		C.io_set_bit(LIGHT_DOOR_OPEN);
	else 
		C.io_clear_bit(LIGHT_DOOR_OPEN);      
		}

}

func SetStopLamp(value bool) {
	if value {
		C.io_set_bit(LIGHT_STOP);
	else
		C.io_clear_bit(LIGHT_STOP);
		}
}
func SetButtonLamp(

func GetFloor() {

}

func FloorIndicator(floor, on/off) { // one light must always be on

}

func ResetLights() {
	
}

func ElevDirection(dir int) {
	
}

func ElevDoorLamp


