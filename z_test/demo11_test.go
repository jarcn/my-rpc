package ztest

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var wg11 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestChan(t *testing.T) {
	court := make(chan int)
	wg11.Add(2)
	go player("nadal", court)
	go player("djokovic", court)
	//开始让两个goroutine准备工作
	court <- 1
	wg11.Wait()
}

func player(name string, court chan int) {
	defer wg11.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s missed \n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("player %s missed \n", name)
			close(court)
			return
		}
		fmt.Printf("player %s hit %d \n", name, ball)
		ball++
		//通知另一个goroutine开始操作
		court <- ball
	}
}
