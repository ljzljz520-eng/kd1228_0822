package model

import "testing"

func TestColorMixAndValidation(t *testing.T) {
	color := NewColor(10, 20, 30)
	if color.Hex() != "#0a141e" {
		t.Fatal(color.Hex())
	}
	if color.Mix(NewColor(110, 120, 130), 0.5) != (Color{R: 60, G: 70, B: 80}) {
		t.Fatal("mix")
	}
	j := NewJellyfish("j", "test", Vector{X: 2, Y: 3}, color, 1, 4)
	if err := j.Validate(); err != nil {
		t.Fatal(err)
	}
	scene := NewScene("s", 100, 100)
	scene.Jellyfish = []Jellyfish{j}
	if err := scene.Validate(); err != nil {
		t.Fatal(err)
	}
}
