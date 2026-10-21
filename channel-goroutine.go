package main

import "C"
import (
	"time"
)

//export asyncCompute
func asyncCompute(n C.int) C.int {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second) // 模拟耗时计算
		ch <- int(n) * 2
	}()

	result := <-ch
	return C.int(result)
}

func main5() {}