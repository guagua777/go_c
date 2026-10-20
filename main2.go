package main

/*
#include <stdio.h>
extern int goCallback(int x);

void run_cb() {
    int r = goCallback(100);
    printf("C receive callback ret: %d\n", r);
}
*/
import "C"

func main() {
    C.run_cb()
}
