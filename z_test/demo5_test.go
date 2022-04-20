package ztest

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

type user struct {
	name  string
	email string
}

//使用值接收者实现了一个方法
//如果使用值接收者声明方法，调用时会使用这个值的一个副本来执行。
func (u user) notify() {
	fmt.Printf("sending user email to %s<%s> \n", u.name, u.email)
}

//使用指针接收者实现了一个方法
//这个方法使用指针接收者声明，这个接收者的类型是指向user类型值的指针，而不是user类型的值。
//当调用使用指针接收者声明的方法时，这个方法会共享调用方法时接收者所指向的值。
func (u *user) changeEmail(email string) {
	u.email = email
}

func TestUser(t *testing.T) {
	bill := user{"Bill", "bill@email.com"}
	bill.notify() //使用值接收者定义的方法，执行notiy方法的是 user 结构体的副本 bill

	lisa := &user{"Lisa", "lisa@email.com"}
	// lisa.notify() 等价于 (*lisa).notify()
	lisa.notify() // go 自动做了优化调用

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}

func TestTrim(t *testing.T) {

	s := "   a b c   "
	t.Log(s)

	//这里传递给 trim 方法的是 s 的副本
	//对副本所作的修改不会影响到原始值
	ns := strings.Trim(s, " ")
	t.Log(s)
	t.Log(ns)
}

func TestCurl(t *testing.T) {
	url := "http://www.baidu.com"
	r, err := http.Get(url)
	if err != nil {
		t.Log(err)
		return
	}
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		t.Log(err)
	}
}

type notifier interface {
	notice()
}

func (u *user) notice() {
	fmt.Printf("sending user email to %s<%s> \n", u.name, u.email)
}

func TestNotice(t *testing.T) {
	bill := user{"Bill", "bill@email.com"}
	sendNotice(&bill)
}

func sendNotice(n notifier) {
	n.notice()
}

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("duration: %d", *d)
}

//不是总能获取一个值的地址，所以值的方法集只包含了使用值接收者实现的方法。
func TestDuration(t *testing.T) {
	d := duration(42)
	t.Log(d.pretty())
}
