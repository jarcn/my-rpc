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
