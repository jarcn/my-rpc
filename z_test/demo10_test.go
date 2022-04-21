package ztest

import (
	"runtime"
	"sync"
	"testing"
)

var (
	counter10 int
	wg10      sync.WaitGroup
	mutex10   sync.Mutex
)

func TestMutex(t *testing.T) {
	wg10.Add(2)

	go incCounter10(1)
	go incCounter10(2)

	wg10.Wait()

	t.Logf("final counter: %d", counter10)
}

func incCounter10(id int) {
	defer wg10.Done()
	for count := 0; count < 2; count++ {
		//对共享资源加锁，保证同一时刻只有一个goroutine能操作共享资源
		mutex10.Lock()
		{
			value := counter10
			runtime.Gosched()
			value++
			counter10 = value
		}
		mutex10.Unlock()
	}
}
