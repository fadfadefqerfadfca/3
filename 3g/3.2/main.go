package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	numChannel := make(chan int, 10)
	strChannel := make(chan string, 10)
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		numChannel <- i
	}
	close(numChannel)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range numChannel {
				strChannel <- strconv.Itoa(num)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(strChannel)
	}()

	for str := range strChannel {
		fmt.Println(str)
	}
}
