package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(prefix string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("%s %d\n", prefix, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printNumbers("Routine1", &wg)
	go printNumbers("Routine2", &wg)

	wg.Wait()
	fmt.Println("Done")
}
