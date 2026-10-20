package main

// // 头文件搜索路径

/*
#cgo CFLAGS: -I./c_src/
#cgo LDFLAGS: -L./c_src/ -lmylib
#include "mylib.h"
*/
import "C"
import "fmt"

func main2() {
    a, b := C.int(10), C.int(20)
    res := C.add(a, b)
    fmt.Printf("add result: %d\n", int(res))
}
