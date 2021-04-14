package structs_and_interfaces

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{Width: 12.0, Height: 6.0}
		want := 72.0

		checkArea(t, rect, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{Radius: 10.0}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})

	t.Run("tables", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
			{shape: Circle{Radius: 10}, want: 314.1592653589793},
			{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
		}

		for _, tt := range areaTests {
			checkArea(t, tt.shape, tt.want)
		}
	})
}

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{Width: 10.0, Height: 10.0}
		want := 40.0

		checkPerimeter(t, rect, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{Radius: 10.0}
		want := 62.83185307179586

		checkPerimeter(t, circle, want)
	})

	t.Run("triangles", func(t *testing.T) {
		triangle := Triangle{Base: 12.0, Height: 6.0}
		want := 24.0

		checkPerimeter(t, triangle, want)
	})

}
