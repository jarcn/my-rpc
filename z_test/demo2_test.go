package ztest

import (
	"fmt"
	"testing"
)

func TestDemo2(t *testing.T) {

	m := make(map[int]int, 10)

	for i := 1; i <= 10; i++ {
		m[i] = i
	}

	//协程泄漏,必包引用外部数据
	//在没有将变量 v 的拷贝值传进匿名函数之前，只能获取最后一次循环的值,这是新手最容易遇到的坑。
	/* for k1, v1 := range m {
		go func() {
			fmt.Println("k1:", k1, "v1:", v1)
		}()
	} */

	//通过传递参数来解决
	for k, v := range m {
		go func(k, v int) {
			fmt.Println("k ->", k, "v ->", v)
		}(k, v)
	}
}
