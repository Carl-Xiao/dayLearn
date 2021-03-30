package main

import (
	"fmt"
	"sync"
)

func sumMutexLock1(s sync.Mutex) {
	s.Lock()
	fmt.Printf("sumMutexLock1, s: %p\n", &s)
	defer s.Unlock()
}

func sumMutexLock2(s sync.Mutex) {
	s.Lock()
	fmt.Printf("sumMutexLock2, s: %p\n", &s)
	defer s.Unlock()
}

func main() {
	mutex := sync.Mutex{}
	fmt.Printf("TestMutex21, s: %p\n", &mutex)
	sumMutexLock1(mutex)
	sumMutexLock2(mutex)
}
