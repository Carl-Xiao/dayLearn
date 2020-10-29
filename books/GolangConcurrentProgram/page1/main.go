package main

import (
	"fmt"
	"time"
)

func main() {
	var data int

	go func() {
		data++
	}()
	if data == 0 {
		fmt.Println("Hello World", data)
	}
	time.Sleep(1 * time.Second)
}
