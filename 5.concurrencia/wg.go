/**
https://golang.org/pkg/sync/#WaitGroup

*/

package main

import (
	"fmt"
	"sync"
	"time"
)


func __main() {
	var wg sync.WaitGroup
	for i :=0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}
	wg.Wait()
}


func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Done %d\n", i)
}