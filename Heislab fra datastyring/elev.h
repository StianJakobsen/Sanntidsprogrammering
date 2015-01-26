#ifndef __INCLUDE_DRIVER_H__
#define __INCLUDE_DRIVER_H__

#include "simulator/io.h"

#define N_FLOORS 4

typedef enum tag_elev_lamp_type { 
    BUTTON_CALL_UP      = 0,
    BUTTON_CALL_DOWN    = 1,
    BUTTON_COMMAND      = 2
} elev_button_type_t;


int     elev_init(ElevatorType et);

void    elev_set_speed(int speed);

void    elev_set_door_open_lamp(int value);

int     elev_get_obstruction_signal(void);

int     elev_get_stop_signal(void);
void    elev_set_stop_lamp(int value);

int     elev_get_floor_sensor_signal(void);
void    elev_set_floor_indicator(int floor);

int     elev_get_button_signal(elev_button_type_t button, int floor);
void    elev_set_button_lamp(elev_button_type_t button, int floor, int value);

#endif
