package ztest

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestGoroutine(t *testing.T) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	t.Log("create goroutines")
	go printPrime("A")
	go printPrime("B")
	t.Log("wating to finish")

	wg.Wait()
	t.Log("terminating program")

}

func printPrime(name string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", name, outer)
	}
	fmt.Println("completed", name)

}
