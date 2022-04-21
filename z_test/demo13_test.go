package ztest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	numberGoroutine = 4  //要使用的goroutine的数量
	taskLoad        = 10 //要处理的工作数量
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestChan2(t *testing.T) {
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutine)
	for gr := 1; gr < numberGoroutine; gr++ {
		go worker(tasks, gr)
	}
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("task:%d", post)
	}
	//当通道关闭后，goroutine依旧可以从通道接收数据，但是不能再向通道里发送数据。
	//能够从已经关闭的通道中接收数据这一点非常重要，因为这允许通道关闭后依旧能取出其中缓冲的全部值，而不会有数据丢失。
	//从一个已经关闭且没有数据的通道里获取数据，总会立刻返回，并返回一个通道类型的零值。
	//如果在获取通道时还加入了可选的标志，就能得到通道的状态信息。
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("worker: %d : shutting down\n", worker)
			return
		}
		fmt.Printf("worker: %d : started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("worker:%d : completed %s\n", worker, task)
	}
}
