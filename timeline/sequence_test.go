package timeline

import (
	"jellyfield/model"
	"testing"
)

func TestClockAndSequence(t *testing.T) {
	clock := New(2)
	if clock.Advance() != 2 {
		t.Fatal(clock.Frame())
	}
	clock.Toggle()
	if clock.Advance() != 2 {
		t.Fatal("paused")
	}
	clock.Toggle()
	if err := clock.Label("intro"); err != nil {
		t.Fatal(err)
	}
	sequence := NewSequence()
	sequence.Add(Keyframe{Frame: 0, Position: model.Vector{X: 0, Y: 0}, Speed: 1})
	sequence.Add(Keyframe{Frame: 10, Position: model.Vector{X: 100, Y: 50}, Speed: 3})
	key, ok := sequence.At(5)
	if !ok || key.Position.X != 50 || key.Speed != 2 {
		t.Fatal(key, ok)
	}
}
