

Getting started
===============

Download the simulator from the project page 
  (https://github.com/TTK4145/Project), and put it in this folder.

If using the hardware at the lab, go to main.c and change 
  "newElevator(ET_simulation, ...)" to "newElevator(ET_comedi, ...)" (on line 19)

If you want to use a custom config for the simulator, you will need to copy the 
  config file (simulator.con) into the same directory as the executable (called
  "lift" when using the makefile).


Implementation notes
====================

It is implemented as a state machine. The state enum is in elevator.h


The outer loop "switches" on events ("get{...}Event(&..)"), and each 
  event-handler switches on state

The most interesting (from an algorithmic point of view) functions are
    chooseDirn()
    clearOrdersAtCurrentFloor()
    shouldStop()
    hasOrders(), ordersAbove() and ordersBelow()

Two config options, for different implementations of "unspecified" behaviour:
    - CFG_ClearOrders:
        (Context: Both "up" and "down" buttons are pressed when the elevator 
          arrives at that floor)
          
        - CO_InDirnOnly:
            Clears only the order in the current direction of travel
        
        - CO_InBothDirns: 
            Clears both orders, with the assumption that people will get on the 
              elevator even though it is going in the wrong direction
                          
    - CFG_NewOrderAtCurrentFloor:
        (Context: Someone pushes a button on the currrent floor when the door 
          is still open)
          
        - NOC_ExtendDoorOpenTime:
            Keeps the door open longer. This means that an obnoxious user can 
              effectively prevent the elevator from moving! (Consider the 
              real-time consequences of this!)
        
        - NOC_IgnoreNewOrder:
            Ignores the button press entirely.                    
                      

elev.c and elev.h are used with minor modifications from the driver (mostly 
  just passing through the ElevatorType)



Other
=====

Is this code "good code"? It was written my me, alone, by myself, without code 
  review. You know what that means.






