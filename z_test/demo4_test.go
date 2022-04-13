package ztest

import (
	"fmt"
	"sync"
	"testing"
)

//使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母
//最终效果如下： 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func TestGo(t *testing.T) {
	str, num := make(chan bool), make(chan bool)
	w := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-num:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				str <- true
			}
		}
	}()
	w.Add(1)
	go func(w *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-str:
				if i >= 'Z' {
					w.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				num <- true
			}
		}
	}(&w)
	num <- true
	w.Wait()
}
