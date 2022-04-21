package ztest

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	shoudown int64
	wg2      sync.WaitGroup
)

func TestAtomic(t *testing.T) {
	wg2.Add(2)
	go dowork("A")
	go dowork("B")
	time.Sleep(1 * time.Second)
	t.Log("shutdown now")
	atomic.StoreInt64(&shoudown, 1)
	wg2.Wait()
}

func dowork(name string) {
	defer wg2.Done()
	for {
		fmt.Printf("doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shoudown) == 1 {
			fmt.Printf("shutting %s down \n", name)
			break
		}
	}
}
