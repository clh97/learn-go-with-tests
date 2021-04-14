package structs_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rect := Rectangle{Width: 12.0, Height: 6.0}
	got := Area(rect)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
