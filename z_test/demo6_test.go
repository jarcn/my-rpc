package ztest

import (
	"runtime"
	"sync"
	"testing"
)

func TestGoRt(t *testing.T) {
	//允许程序更改调度器可以使用的逻辑处理器数量。
	//函数参数指的是通知调度器只能为该程序使用指定数量的逻辑处理函数。
	runtime.GOMAXPROCS(2)
	//计数信号量，用来记录并维护运行的goroutine。
	var wg sync.WaitGroup
	//大于0，wait方法会阻塞。
	wg.Add(2)
	t.Log("start goroutines")
	go func() {
		//通知wg程序执行完成。
		//defer会修改函数调用时机，在正在执行的函数返回时才真正调用defer声明的函数。
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				t.Logf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				t.Logf("%c", char)
			}
		}
	}()

	t.Log("waiting to finish")
	wg.Wait()

	t.Log("terminating program")
}
