package main

/*
#include <stdio.h>
#include <stdlib.h>

void hello(const char* msg) {
    printf("C says: %s\n", msg);
}
*/
import "C"

import (
	"unsafe"
)

func main1() {
    msg := "Hello from Go"
    // Go string -> C string
    cMsg := C.CString(msg)
    defer C.free(unsafe.Pointer(cMsg)) // 必须释放！CString在C堆分配

    C.hello(cMsg)
}





// import "fmt"

// func main() {
// 	fmt.Println("Hello, go-c!")
// }