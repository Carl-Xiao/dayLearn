package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, value := range []string{"hello", "greeting", "world"} {
		wg.Add(1)
		go func(value string) {
			defer wg.Done()
			// loop variable value captured by func literal
			ip := &value
			fmt.Println(&ip)
			fmt.Printf("%s原始指针的内存地址是：%p\n", value, &value)
		}(value)
	}
	wg.Wait()
}
