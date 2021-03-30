package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type config struct {
	a []int
}

func main() {
	var v atomic.Value
	go func() {
		i := 0
		for {
			i++
			a := config{
				a: []int{i, i + 1, i + 2, i + 3, i + 4},
			}
			v.Store(a)
		}
	}()
	var wg sync.WaitGroup

	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				a := v.Load().(config)
				fmt.Println(a.a)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
