package ztest

import (
	"fmt"
	"testing"
)

//这里打印出来的结果是最后一次循环结果
// func farr1() []func() {
// 	var s []func()
// 	for i := 0; i < 3; i++ {
// 		//将多个匿名函数添加到列表
// 		s = append(s, func() { fmt.Println(&i, i) })
// 	}
// 	return s //返回匿名函数列表
// }

func farr() []func() {
	var s []func()
	for i := 0; i < 3; i++ {
		x := i
		//将多个匿名函数添加到列表
		s = append(s, func() { fmt.Println(&i, x) })
	}
	return s //返回匿名函数列表
}

func TestDemo3(t *testing.T) {
	for _, f := range farr() { //执行所有匿名函数
		f()
	}
}

var x int = 1

//闭包修改全局变量
func TestMain(t *testing.T) {
	y := func() int {
		x += 1
		return x
	}()
	fmt.Println("main:", x, y)
}

/**
闭包总结:
从形式上看，匿名函数都属于闭包.闭包使用非常灵活。
首先检查函数的参数，声明可以接收参数的匿名函数，这些类型的闭包问题也就引刃而解了。
*/
