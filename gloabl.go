package main

/*
#cgo CFLAGS: -I./c_src/
#cgo LDFLAGS: -L./c_src/ -lgloabl
#include "gloabl.h"
*/
import "C"
import "fmt"

func main() {
    C.increment()
    C.getCount()
    fmt.Printf("finish\n")
}