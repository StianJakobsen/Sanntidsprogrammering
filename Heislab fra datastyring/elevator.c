#include "elevator.h"

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <signal.h>


void elevator_terminate(int __attribute__((unused)) v){
    elev_set_speed(0);
    exit(0);
}

void __attribute__((constructor)) elevator_init(void){
    signal(SIGINT, elevator_terminate);
}




Elevator newElevator(ElevatorType et, int speed, CFG_ClearOrders co, CFG_NewOrderAtCurrentFloor noc){
    int ok = elev_init(et);
    if(!ok){
        printf("Hardware init failed!\n");
        exit(1);
    }
    
    Elevator e = {0};
    
    e.state = init;
    e.floor = -1;
    e.dirn = stop;
    
    e.config.speed = speed;
    e.config.clearOrders = co;
    e.config.newOrderAtCurrentFloor = noc;
    
    setMotorDirn(&e, 0);
    
    return e;
}


int equals(Elevator const * const first, Elevator const * const second){
    if( first->floor  != second->floor  ||
        first->dirn   != second->dirn   ||
        first->state  != second->state
    ){
        return 0;
    }
    for(int floor = 0; floor < N_FLOORS; floor++){
        for(int btn = 0; btn < 3; btn++){
            if(first->orders[floor][btn] != second->orders[floor][btn]){
                return 0;
            }
        }
    }
    return 1;
}


char* toString(ElevatorState es){
    return
        es == init ?          "Init" :
        es == idle ?          "Idle" :
        es == moving ?        "Moving" :
        es == doorOpen ?      "Door open" :
                              "UNDEFINED";
}


void printElevatorState(Elevator const * const e){
    printf(
        "+---"      "+-------+\n"
        "| %-10s"           "|\n"
        "| %s "     "| u d c |\n"
        "+---"      "+-------+\n",
        toString(e->state),
        e->dirn == down ? "v" : 
        e->dirn == up   ? "^" : 
                          "-"
    );
    for(int floor = N_FLOORS-1; floor >= 0; floor--){
        printf("|%s%d | ",
            floor == e->floor ? ">" : " ",
            floor
        );
        for(int btn = 0; btn < 3; btn++){
            if( (btn == BUTTON_CALL_UP      && floor == N_FLOORS-1) ||
                (btn == BUTTON_CALL_DOWN    && floor == 0))
            {
                printf("  ");
            } else {
                printf(
                    "%s ", 
                    e->orders[floor][btn] ? "*" : "-"
                );
            }
        }
        printf("|\n");
    }
    printf(
        "+---+-------+\n"
    );
}



int getButtonPressEvent(ButtonPressEvent * const bpe){
    static int currButtons[N_FLOORS][3];
    static int prevButtons[N_FLOORS][3];
    
    for(int floor = 0; floor < N_FLOORS; floor++){
        for(int btn = 0; btn < 3; btn++){
            if( (btn == BUTTON_CALL_UP      && floor == N_FLOORS-1) ||
                (btn == BUTTON_CALL_DOWN    && floor == 0))
            {
                continue;
            }
            
            prevButtons[floor][btn] = currButtons[floor][btn];
            currButtons[floor][btn] = elev_get_button_signal(btn, floor);
            
            if(currButtons[floor][btn] != prevButtons[floor][btn]  &&  currButtons[floor][btn]){
                *bpe = (ButtonPressEvent){.floor = floor, .button = btn};
                return 1;
            }
        }
    }
    return 0;
}

int getFloorArrivalEvent(FloorArrivalEvent * const fae){
    static int currFloor = -1;
    static int prevFloor = -1;

    prevFloor = currFloor;
    currFloor = elev_get_floor_sensor_signal();
    if(currFloor != prevFloor && currFloor != -1){
        *fae = (FloorArrivalEvent){.floor = currFloor};
        return 1;
    } else {
        return 0;
    }
}



void chooseDirn(Elevator * const e){
    if(e->state == moving){
        printf("\7Chosing a new direction when moving makes no sense!\n"
               "  No action taken\n");
        return;
    }
    if(!hasOrders(e)){
        e->dirn = stop;
        return;
    }
    switch(e->dirn){
    case down:
        if(ordersBelow(e)  &&  e->floor != 0){
            e->dirn = down;
        } else {
            e->dirn = up;
        }
        break;
        
    case up:
        if(ordersAbove(e)  &&  e->floor != N_FLOORS-1){
            e->dirn = up;
        } else {
            e->dirn = down;
        }
        break;
        
    case stop:
        if(ordersAbove(e)){
            e->dirn = up;
        } else if(ordersBelow(e)) {
            e->dirn = down;
        } else {
            e->dirn = stop;
        }
        break;
        
    default:
        printf("  WAT at %s:%d: Elevator has unexpected direction %d!\n", __FUNCTION__, __LINE__, e->dirn);
        e->dirn = stop;
        break;
    }
}

void clearOrdersAtCurrentFloor(Elevator * const e){
    if(e->state == moving){
        printf("\7Clearing orders when moving makes no sense!\n"
               "  No action taken\n");
        return;
    }
    
    e->orders[e->floor][BUTTON_COMMAND] = 0;
    
    switch(e->config.clearOrders){
    case CO_InDirnOnly:
        switch(e->dirn){
        case down:
            e->orders[e->floor][BUTTON_CALL_DOWN] = 0;
            if(!ordersBelow(e)){
                e->orders[e->floor][BUTTON_CALL_UP] = 0;
            }
            break;
            
        case up:
            e->orders[e->floor][BUTTON_CALL_UP] = 0;
            if(!ordersAbove(e)){
                e->orders[e->floor][BUTTON_CALL_DOWN] = 0;
            }
            break;
            
        case stop:
            e->orders[e->floor][BUTTON_CALL_UP] = 0;
            e->orders[e->floor][BUTTON_CALL_DOWN] = 0;
            break;
            
        default:
            printf("  WAT at %s:%d: Elevator has unexpected direction %d!\n", __FUNCTION__, __LINE__, e->dirn); 
            break;
        }
        break;
        
    case CO_InBothDirns:
        e->orders[e->floor][BUTTON_CALL_UP] = 0;
        e->orders[e->floor][BUTTON_CALL_DOWN] = 0;
        break;
    }
}





void setButtonLights(Elevator const * const e){
    for(int floor = 0; floor < N_FLOORS; floor++){
        for(int btn = 0; btn < 3; btn++){
            if( (btn == BUTTON_CALL_UP      && floor == N_FLOORS-1) ||
                (btn == BUTTON_CALL_DOWN    && floor == 0))
            {
                continue;
            }
            elev_set_button_lamp(btn, floor, e->orders[floor][btn]);
        }
    }
}



void setMotorDirn(Elevator const * const e, Dirn dirn){
    elev_set_speed(e->config.speed * dirn);
}





int hasOrders(Elevator const * const e){
    for(int floor = 0; floor < N_FLOORS; floor++){
        for(int btn = 0; btn < 3; btn++){
            if(e->orders[floor][btn]){
                return 1;
            }
        }
    }
    return 0;
}

int ordersAbove(Elevator const * const e){
    for(int floor = e->floor+1; floor < N_FLOORS; floor++){
        for(int btn = 0; btn < 3; btn++){
            if(e->orders[floor][btn]){
                return 1;
            }
        }
    }
    return 0;
}

int ordersBelow(Elevator const * const e){
    for(int floor = 0; floor < e->floor; floor++){
        for(int btn = 0; btn < 3; btn++){
            if(e->orders[floor][btn]){
                return 1;
            }
        }
    }
    return 0;
}



int shouldStop(Elevator const * const e){
    switch(e->dirn){
    case -1:
        return
            e->orders[e->floor][BUTTON_CALL_DOWN]   ||
            e->orders[e->floor][BUTTON_COMMAND]     ||
            e->floor == 0                           ||
            !ordersBelow(e);
    case 1:
        return
            e->orders[e->floor][BUTTON_CALL_UP]     ||
            e->orders[e->floor][BUTTON_COMMAND]     ||
            e->floor == N_FLOORS-1                  ||
            !ordersAbove(e);
    case 0:
        return
            e->orders[e->floor][BUTTON_CALL_DOWN]   ||
            e->orders[e->floor][BUTTON_CALL_UP]     ||
            e->orders[e->floor][BUTTON_COMMAND];
    default:
        printf("  WAT at %s:%d: Elevator has unexpected direction %d!\n", __FUNCTION__, __LINE__, e->dirn);
        return 0;    
    }
}

void removeAllOrders(Elevator * const e){
    for(int floor = 0; floor < N_FLOORS; floor++){
        for(int btn = 0; btn < 3; btn++){
            e->orders[floor][btn] = 0;
        }
    }
}


















