#include "elev.h"
#include "timer.h"

typedef enum ElevatorState ElevatorState;
enum ElevatorState {
    init,
    idle,
    moving,
    doorOpen
};

typedef enum {
    down    = -1,
    stop    = 0,
    up      = 1
} Dirn;

typedef enum {
    NOC_ExtendDoorOpenTime,
    NOC_IgnoreNewOrder
} CFG_NewOrderAtCurrentFloor;

typedef enum {
    CO_InDirnOnly,
    CO_InBothDirns
} CFG_ClearOrders;

typedef struct Elevator Elevator;
struct Elevator {
    ElevatorState   state;
    int             floor;
    Dirn            dirn;
    
    int             orders[N_FLOORS][3];
    
    Timer           timer;
    
    struct {
        int                         speed;
        CFG_NewOrderAtCurrentFloor  newOrderAtCurrentFloor;
        CFG_ClearOrders             clearOrders;
    } config;
};

typedef struct ButtonPressEvent ButtonPressEvent;
struct ButtonPressEvent {
    int             floor;
    int             button;
};

typedef struct FloorArrivalEvent FloorArrivalEvent;
struct FloorArrivalEvent {
    int             floor;
};


// Utilities
Elevator newElevator(ElevatorType et, int speed, CFG_ClearOrders co, CFG_NewOrderAtCurrentFloor noc);
int equals(Elevator const * const first, Elevator const * const second);
char* toString(ElevatorState es);
void printElevatorState(Elevator const * const e);

// Get events from hardware. Returns 1 if event has occured (pointer modified)
int getButtonPressEvent(ButtonPressEvent * const bpe);
int getFloorArrivalEvent(FloorArrivalEvent * const fae);

// Things that can be done when standing still
void chooseDirn(Elevator * const e);
void clearOrdersAtCurrentFloor(Elevator * const e);

// Actuation
void setButtonLights(Elevator const * const e);
void setMotorDirn(Elevator const * const e, Dirn dirn);

// State in(tro)spection
int hasOrders(Elevator const * const e);
int ordersAbove(Elevator const * const e);
int ordersBelow(Elevator const * const e);
int shouldStop(Elevator const * const e);

// 
void removeAllOrders(Elevator * const e);













