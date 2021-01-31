package main

import (
	"fmt"
	"time"
)

/*
	count sheep and fish side by side
*/
func main() {
	c := make(chan string)
	go count("sheep", c)

	for msg := range c {

		// **  msg := range c did this job
		// msg, open := <-c
		// if !open {
		// 	break
		// }
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
