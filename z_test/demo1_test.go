package ztest

import "testing"

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func TestPerimeter(t *testing.T) {
	r := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(r)
	want := 100.00
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

/* func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func Perimeter(a, b float64) float64 {
	return a + b
} */

//切片共享底层数组存储地址
func TestSlice1(t *testing.T) {

	// slice := []int{1, 2, 3, 4, 5}
	// ns := slice[0:3]
	// for k, v := range ns {
	// 	ns[k] = v + 1
	// }
	// for k, v := range slice {
	// 	t.Log(k, v)
	// }

	m := map[string]string{"a": "A"}
	m["a"] = "123123"
	t.Log(m["a"])

}
