package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("in go routine", i)
	wg.Done()
}

func main1() {
	go hello(1)
	fmt.Println("exit main")
	time.Sleep(time.Second) // bad practice
}

/*
error caused by closure
*/
func main_wrong() {

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println("in go routine", i)
			wg.Done()
		}()
	}

	fmt.Println("exit main")
	wg.Wait()
}
func main() {

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println("in go routine", i)
			wg.Done()
		}(i)
	}

	fmt.Println("exit main")
	wg.Wait()
}
