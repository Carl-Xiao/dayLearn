package main

import (
	"fmt"
	"sync"
)

type config struct {
	a []int
}

//使用mutex处理并发的问题
var lock sync.Mutex

func main() {
	a := config{}

	go func() {
		i := 0
		for {
			i++
			lock.Lock()
			a.a = []int{i, i + 1, i + 2, i + 3, i + 4}
			lock.Unlock()
		}
	}()
	var wg sync.WaitGroup

	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				lock.Lock()
				fmt.Println(a.a)
				lock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
