#pragma once

typedef struct Timer Timer;
struct Timer {
    unsigned long   timerEndTime;
    int             timerActive;
};

void timer_start(Timer* t, unsigned long duration);
void timer_abort(Timer* t);
int timer_hasTimedOut(Timer* t);

