package structs_test

import (
	"testing"

	"github.com/rssharma75/learn_go/structs"
)

func assert(got, want float64, shape structs.Shape, t *testing.T) {
	if got != want {
		t.Errorf("test Failed for %#v,  expected:%v, got:%v", shape, want, got)
	}
}
func TestStructs(t *testing.T) {
	t.Run("Test Peremeter", func(t *testing.T) {
		rectangle := structs.Rectangle{Length: 10.0, Width: 10.0}
		got := rectangle.Perimeter()

		want := 40.0

		assert(got, want, rectangle, t)

	})

	t.Run("test Area", func(t *testing.T) {
		r := structs.Rectangle{Length: 4.0, Width: 6.0}
		got := r.Area()

		want := 24.0

		assert(got, want, r, t)

	})

}

func TestCircles(t *testing.T) {
	t.Run("Area", func(t *testing.T) {
		circle := structs.Circle{Radius: 10.0}
		got := circle.Area()
		want := 314.15999999999997

		assert(got, want, circle, t)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape structs.Shape
		want  float64
	}{
		{structs.Rectangle{6.0, 4.0}, 24.0},
		{structs.Circle{10.0}, 314.15999999999997},
		{structs.Triangle{Base: 10.0, Height: 10.0}, 50.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.want
		assert(got, want, tt.shape, t)
	}
}
