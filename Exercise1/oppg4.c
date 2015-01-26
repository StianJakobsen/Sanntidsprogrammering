#include <pthread.h>
#include <stdio.h>

// Note the return type: void
void* countDownFromMillion(){
	for(int i = 1000000; i > 0; i--){
		// DoNothing
	}
	return NULL;
}

void* countUpToMillion(){
	for(int i = 0; i < 1000000; i++){}
	return NULL;
}



int main(){
	pthread_t nedThread;
	pthread_t oppThread;
	pthread_create(&nedThread, NULL, countDownFromMillion, NULL);
	pthread_create(&oppThread, NULL, countUpToMillion, NULL);

