package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello() {
	fmt.Println("in go routine")
	wg.Done()
}

func main1() {
	go hello()
	fmt.Println("exit main")
	time.Sleep(time.Second) // bad practice
}

func main() {

	wg.Add(1)
	go hello()
	fmt.Println("exit main")
	wg.Wait()
}
