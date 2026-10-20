package main

import "C"
import "fmt"

//export goCallback
func goCallback(x C.int) C.int {
	fmt.Printf("Go callback got: %d\n", int(x))
	return x + 1
}
