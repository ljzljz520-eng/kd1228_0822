package layout

import (
	"jellyfield/app"
	"jellyfield/model"
	"testing"
)

func TestArrangeNearestAndNormalization(t *testing.T) {
	layout := Arrange(app.DemoScene())
	id, ok := layout.Nearest(model.Vector{X: 220, Y: 210})
	if !ok || id != "jelly-a" {
		t.Fatal(id, ok)
	}
	point := layout.Normalize(model.Vector{X: 400, Y: 250})
	if point.X != .5 || point.Y != .5 {
		t.Fatal(point)
	}
	if !layout.InBounds(model.Vector{X: 1, Y: 1}) {
		t.Fatal("bounds")
	}
}
