#include <time.h>

#include "timer.h"

void timer_start(Timer* t, unsigned long duration){
    t->timerEndTime = time(0) + duration;
    t->timerActive  = 1;
}

void timer_abort(Timer* t){
    t->timerActive = 0;
}

int timer_hasTimedOut(Timer* t){ //impure
    if(t->timerActive  &&  time(0) > t->timerEndTime){    
        t->timerActive = 0;
        return 1;
    } else {
        return 0;
    }
}

