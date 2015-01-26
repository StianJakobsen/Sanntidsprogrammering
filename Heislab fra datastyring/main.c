
// Avoids annoying warning about "implicit declaration of function usleep"
#define _BSD_SOURCE 

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "elevator.h"


static const unsigned long  doorOpenTime = 3;


int main(){
    printf("started\n");
    
    Elevator            elevator    = newElevator(ET_simulation, 300, CO_InBothDirns, NOC_IgnoreNewOrder);
    ButtonPressEvent    bpe;
    FloorArrivalEvent   fae;
    
    
    if(elev_get_floor_sensor_signal() == -1){
        setMotorDirn(&elevator, down);
    }
    
    printElevatorState(&elevator);
    
    while(1){
    
        if(getButtonPressEvent(&bpe)){
            printf("The %s button was pressed on floor %d\n",
                bpe.button == BUTTON_CALL_UP    ? "  UP   " :
                bpe.button == BUTTON_CALL_DOWN  ? " DOWN  " :
                bpe.button == BUTTON_COMMAND    ? "COMMAND" : 
                                                  "(\7WAT!) ",
                bpe.floor
            );
            
            printElevatorState(&elevator);
            
            switch(elevator.state){
            case init:
                break;
                
            case idle:
                elevator.orders[bpe.floor][bpe.button] = 1;
                
                clearOrdersAtCurrentFloor(&elevator);
                chooseDirn(&elevator);
                
                if(elevator.dirn != stop){
                    setMotorDirn(&elevator, elevator.dirn);
                    
                    elevator.state = moving;
                } else {
                    elev_set_door_open_lamp(1);
                    timer_start(&elevator.timer, doorOpenTime);
                    
                    elevator.state = doorOpen;
                }
                break;
                
            case moving:
                elevator.orders[bpe.floor][bpe.button] = 1;
                break;
                
            case doorOpen:
                switch(elevator.config.newOrderAtCurrentFloor){
                case NOC_ExtendDoorOpenTime:
                    elevator.orders[bpe.floor][bpe.button] = 1;                    
                    if(shouldStop(&elevator)){
                        timer_start(&elevator.timer, doorOpenTime);
                        clearOrdersAtCurrentFloor(&elevator);
                    }
                    break;
                case NOC_IgnoreNewOrder:
                    if(bpe.floor != elevator.floor){
                        elevator.orders[bpe.floor][bpe.button] = 1;
                    }
                    break;
                }
                break;
                
            default:
                printf( "  \7WAT at %s:%d: "
                        "Unexpected state in event \"getButtonPressEvent{button:%d, floor:%d}\": %s\n",
                        __FUNCTION__, __LINE__, bpe.button, bpe.floor, toString(elevator.state));
                break;
            }
            
            setButtonLights(&elevator);
            
            printElevatorState(&elevator);
        }
        
        if(getFloorArrivalEvent(&fae)){
            printf("Elevator arrived at floor %d\n", fae.floor);
            elevator.floor = fae.floor;

            printElevatorState(&elevator);

            elev_set_floor_indicator(elevator.floor);
            
            switch(elevator.state){
            case init:
                elevator.floor = fae.floor;
                setMotorDirn(&elevator, stop);
                
                elevator.state = idle;
                break;
                
            case idle:
                printf( "  \7WAT at %s:%d: "
                        "Unexpected state in event \"getFloorArrivalEvent{floor:%d}\": %s\n",
                        __FUNCTION__, __LINE__, fae.floor, toString(elevator.state));
                break;
                
            case moving:
                if(shouldStop(&elevator)){
                    setMotorDirn(&elevator, stop);
                    
                    elevator.state = doorOpen;  // must be set *before* call to clearOrdersAtCurrentFloor
                    
                    clearOrdersAtCurrentFloor(&elevator);
                    setButtonLights(&elevator);
                    timer_start(&elevator.timer, doorOpenTime);
                    elev_set_door_open_lamp(1);
                }
                break;
                            
            default:
                printf( "  \7WAT at %s:%d: "
                        "Unexpected state in event \"getFloorArrivalEvent{floor:%d}\": %s\n",
                        __FUNCTION__, __LINE__, fae.floor, toString(elevator.state));
                break;
            }
            
            printElevatorState(&elevator);
        }
        
        if(timer_hasTimedOut(&elevator.timer)){
            printf("Elevator timer has timed out\n");
            
            printElevatorState(&elevator);
            
            switch(elevator.state){
            case init:
            case idle:
            case moving:
                break;
                
            case doorOpen:
                chooseDirn(&elevator);
                elev_set_door_open_lamp(0);
                if(elevator.dirn){
                    setMotorDirn(&elevator, elevator.dirn);
                    
                    elevator.state = moving;
                } else {
                    elevator.state = idle;
                }
                break;
                
                
            default:
                printf( "  \7WAT at %s:%d: "
                        "Unexpected state in event \"timer_hasTimedOut{}\": %s\n",
                        __FUNCTION__, __LINE__, toString(elevator.state));
                break;
            }
            
            printElevatorState(&elevator);
        }
    
        usleep(1000*10);
    }
    
}


