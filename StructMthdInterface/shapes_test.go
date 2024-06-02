package main

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.g want %g", got, want)
		}
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 62.83185307179586

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
