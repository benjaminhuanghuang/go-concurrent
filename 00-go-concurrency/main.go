package main

import (
	"fmt"
	"time"
)

/*
	count sheep and fish side by side
*/
func main() {
	go count("sheep")
	go count("fish")

	// Blocking the main funciton from exiting
	fmt.Scanln()
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
