#include<stdio.h>

int main(){
  int ant;
  int i;

  printf("Hvor mange ganger vil du at 'Hello World!' skal skrives ut?\n");
  scanf("%d", &ant);
  for(i = 0; i < ant; ++i){
    printf("Hello World\n");
  }
}
