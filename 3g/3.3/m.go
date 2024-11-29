package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	time.Sleep(1 * time.Second)

	val := <-ch
	fmt.Println("горутина получила значение", val)
}
