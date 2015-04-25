// driver

package driver // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.c and driver.go
/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
//import "C"

const N_FLOORS = 4

//Lights
var upLights = [N_FLOORS]int{LIGHT_UP1, LIGHT_UP2, LIGHT_UP3, LIGHT_UP4}
var downLights = [N_FLOORS]int{LIGHT_DOWN1, LIGHT_DOWN2, LIGHT_DOWN3, LIGHT_DOWN4}
var cmdLights = [N_FLOORS]int{LIGHT_COMMAND1,LIGHT_COMMAND2, LIGHT_COMMAND3, LIGHT_COMMAND4}

//Buttons
var upButtons = [N_FLOORS]int{BUTTON_UP1, BUTTON_UP2, BUTTON_UP3, BUTTON_UP4}
var downButtons = [N_FLOORS]int{BUTTON_DOWN1, BUTTON_DOWN2, BUTTON_DOWN3, BUTTON_DOWN4}
var CmdButtons = [N_FLOORS]int{BUTTON_COMMAND1, BUTTON_COMMAND2, BUTTON_COMMAND3, BUTTON_COMMAND4}

const (
BUTTON_CALL_UP int = iota
BUTTON_CALL_DOWN int= iota
BUTTON_COMMAND int = iota 
)
const(
DIRN_DOWN int = iota -1
DIRN_STOP int = iota -1
DIRN_UP int = iota -1
)
func errorFloor(floor int) bool{
	if floor > N_FLOORS || floor < 0 {
		return true 
	} else {
		return false
	} 
}

func InitElevator() int { // sjekke denna
	if io_init()== 0 {
		return 0
	}
	for i := 0; i < N_FLOORS; i++ {
		if i != 0{
			SetButtonLamp(BUTTON_CALL_DOWN, i, 0)
		} 		
		if i != N_FLOORS - 1 {
			SetButtonLamp(BUTTON_CALL_UP, i, 0)
		}
		SetButtonLamp(BUTTON_COMMAND, i, 0)
	}
	SetStopLamp(false)
	SetMotorDirection(DIRN_STOP)
	SetDoorOpenLamp(false)
	if GetFloorSensorSignal() != -1 {
		SetFloorIndicator(GetFloorSensorSignal())
	} 
	// + noko greiar herat
	return 1
}

func SetMotorDirection(dir int) {
	if dir == 0 {
		io_write_analog(MOTOR, 0)	
	} else if dir > 0 {
		io_clear_bit(MOTORDIR)
		io_write_analog(MOTOR, 2800)
	} else if dir < 0 {
		io_set_bit(MOTORDIR)
		io_write_analog(MOTOR, 2800)
	}
	
}

func SetDoorOpenLamp(value bool) {
	if value {
		io_set_bit(LIGHT_DOOR_OPEN)
	} else {
		io_clear_bit(LIGHT_DOOR_OPEN)      
	}
}

func GetObstructionSignal() int {
	return io_read_bit(OBSTRUCTION)
}

func GetStopSignal() int {
	return io_read_bit(STOP)
}

func SetStopLamp(value bool) {
	if value {
		io_set_bit(LIGHT_STOP);
	} else {
		io_clear_bit(LIGHT_STOP);
	} 
}

func GetFloorSensorSignal() int {
	if io_read_bit(SENSOR_FLOOR1) == 1 {
		return 0 // ground floor
	} else if io_read_bit(SENSOR_FLOOR2) == 1 {
		return 1
	} else if io_read_bit(SENSOR_FLOOR3) == 1 {
		return 2
	} else if io_read_bit(SENSOR_FLOOR4) == 1 {
		return 3
	} else {
		return -1; // between floors
	}
}

func SetFloorIndicator(floor int) { // one light must always be on
	// sjekke om rektig input?
	if errorFloor(floor) == false {
	byteFloor := byte(floor) 
		if (byteFloor & 0x02) != 0 {
			io_set_bit(LIGHT_FLOOR_IND1)
		} else {
			io_clear_bit(LIGHT_FLOOR_IND1)
		}
		if (byteFloor & 0x01) != 0 {
			io_set_bit(LIGHT_FLOOR_IND2)
		} else {
			io_clear_bit(LIGHT_FLOOR_IND2)
		}
}
}

func GetButtonSignal(button int, floor int) int{
	if !(errorFloor(floor)){
		if((button == BUTTON_CALL_UP && floor == N_FLOORS-1) == false) {
			if((button == BUTTON_CALL_DOWN && floor == 0) == false) {
				if(button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND){
					if button == BUTTON_CALL_UP {
						if io_read_bit(upButtons[floor]) != 0 {
							return 1
						} else {
							return 0
						}
					} else if (button == BUTTON_CALL_DOWN) {
						if io_read_bit(downButtons[floor]) != 0 {
							return 1
						} else {
							return 0
						}
					} else {
						if io_read_bit(CmdButtons[floor]) != 0 {
							return 1
						} else {
							return 0
						}
					}
}
}}}
return 0
}

func SetButtonLamp(button int, floor int, value int) {
	if errorFloor(floor) == false {
		if(button == BUTTON_CALL_UP && floor == N_FLOORS-1) == false {
			if(button == BUTTON_CALL_DOWN && floor == 0) == false {
				if(button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND){
					if button == BUTTON_CALL_UP {
						if value != 0 {
							io_set_bit(upLights[floor])
						} else {
							io_clear_bit(upLights[floor])
						}
					} else if button == BUTTON_CALL_DOWN {
						if value != 0 {
							io_set_bit(downLights[floor])
						} else {
							io_clear_bit(downLights[floor])
						}
					} else {
						if value != 0 {
							io_set_bit(cmdLights[floor])
						} else {
							io_clear_bit(cmdLights[floor])
						}
					}					
}
}}}}
 




////////////////////////////////////////////////////////////////////




//func ResetLights() {
//	
//}



