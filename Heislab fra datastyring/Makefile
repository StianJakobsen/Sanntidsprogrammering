
CC = gcc

CFLAGS = -std=c11 -g -Wall

LDFLAGS = -lpthread -lcomedi -lm

ELEVSRC = elev.c main.c elevator.c timer.c
ELEVOBJ = $(ELEVSRC:.c=.o)

DRIVERSRC = simulator/io.c simulator/simulation_elevator.a simulator/libphobos2.a



TARGET = lift

$(TARGET): $(ELEVOBJ)
	$(CC) $(CFLAGS) -o $(TARGET) $(ELEVOBJ) $(DRIVERSRC) $(LDFLAGS)
	
%.o : %.c
	$(CC) $(CFLAGS) -c $< -o $@ 
	
clean:
	rm -f $(TARGET) $(ELEVOBJ)


.PHONY: clean

