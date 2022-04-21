package ztest

import (
	"fmt"
	"testing"
	"time"
)

func TestDemo12(t *testing.T) {
	baton := make(chan int)
	wg.Add(1)
	go Runner(baton)
	baton <- 1 //开始比赛信号
	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton
	fmt.Printf("runner %d running with baton \n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("runner %d to the line \n", runner)
		go Runner(baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Printf("runner %d finished, race over \n", runner)
		wg.Done()
		return
	}
	fmt.Printf("runner %d exchange with runner %d\n", runner, newRunner)
	baton <- newRunner
}
