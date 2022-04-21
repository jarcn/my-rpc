package ztest

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	counter int64
	wg1     sync.WaitGroup
)

func TestG(t *testing.T) {
	wg1.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg1.Wait()
	t.Log("final counter:", counter)
}

func incCounter(id int) {
	defer wg1.Done()
	for count := 0; count < 2; count++ {
		//捕获counter的值
		// value := counter
		//原子操作，确保同一时刻只有一个goroutine操作共享资源
		atomic.AddInt64(&counter, 1)
		//当前goroutine从线程退出，并放回到队列
		runtime.Gosched()
		// 	//副本值增加
		// 	value++
		// 	//副本值写回到共享变量
		// 	counter = value
	}
}
