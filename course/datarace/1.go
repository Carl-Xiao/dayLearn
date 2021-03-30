package main

import (
	"fmt"
	"sync"
)

type config struct {
	a []int
}

func main() {
	a := config{}

	go func() {
		i := 0
		for {
			i++
			a.a = []int{i, i + 1, i + 2, i + 3, i + 4}
		}
	}()
	var wg sync.WaitGroup

	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				fmt.Println(a.a)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
