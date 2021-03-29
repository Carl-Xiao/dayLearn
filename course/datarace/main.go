package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup
var Counter int

func Rune(id int) {
	for count := 0; count < 2; count++ {
		// value := Counter
		// value++
		// Counter = value

		// Counter++

		//使用原子操作 Counter = Counter + 1

	}
	wait.Done()
}

func main() {
	for r := 1; r <= 2; r++ {
		wait.Add(1)
		go Rune(r)
	}
	wait.Wait()
	fmt.Println(Counter)
}
