package input

import (
	"fmt"
	"jellyfield/model"
)

type Pointer struct {
	ID       int
	Position model.Vector
	Pressed  bool
	Pressure float64
}
type GestureBuilder struct {
	last   Pointer
	active bool
	points []model.Vector
}

func NewGestureBuilder() *GestureBuilder { return &GestureBuilder{points: []model.Vector{}} }
func (b *GestureBuilder) Begin(pointer Pointer) error {
	if b.active {
		return fmt.Errorf("gesture already active")
	}
	b.active = true
	b.last = pointer
	b.points = []model.Vector{pointer.Position}
	return nil
}
func (b *GestureBuilder) Move(pointer Pointer) error {
	if !b.active {
		return fmt.Errorf("gesture not active")
	}
	b.last = pointer
	b.points = append(b.points, pointer.Position)
	return nil
}
func (b *GestureBuilder) End(pointer Pointer) (Trace, error) {
	if !b.active {
		return Trace{}, fmt.Errorf("gesture not active")
	}
	b.last = pointer
	b.points = append(b.points, pointer.Position)
	trace := Trace{Start: b.points[0], End: pointer.Position, Points: append([]model.Vector(nil), b.points...), Pressure: pointer.Pressure}
	b.active = false
	return trace, nil
}
func (b *GestureBuilder) Cancel()      { b.active = false; b.points = nil }
func (b *GestureBuilder) Active() bool { return b.active }

type Trace struct {
	Start    model.Vector
	End      model.Vector
	Points   []model.Vector
	Pressure float64
}

func (t Trace) Delta() model.Vector {
	return model.Vector{X: t.End.X - t.Start.X, Y: t.End.Y - t.Start.Y}
}
func (t Trace) Distance() float64 { return t.Start.Distance(t.End) }
func (t Trace) Direction() string {
	delta := t.Delta()
	if abs(delta.X) >= abs(delta.Y) {
		if delta.X >= 0 {
			return "right"
		}
		return "left"
	}
	if delta.Y >= 0 {
		return "down"
	}
	return "up"
}
func (t Trace) Smooth() Trace {
	if len(t.Points) < 3 {
		return t
	}
	result := Trace{Start: t.Start, End: t.End, Pressure: t.Pressure, Points: make([]model.Vector, 0, len(t.Points))}
	for i := range t.Points {
		start := i - 1
		if start < 0 {
			start = 0
		}
		end := i + 1
		if end >= len(t.Points) {
			end = len(t.Points) - 1
		}
		sum := model.Vector{}
		for j := start; j <= end; j++ {
			sum = sum.Add(t.Points[j])
		}
		result.Points = append(result.Points, sum.Scale(1/float64(end-start+1)))
	}
	return result
}

func (t Trace) Centroid() model.Vector {
	if len(t.Points) == 0 {
		return t.Start
	}
	sum := model.Vector{}
	for _, point := range t.Points {
		sum = sum.Add(point)
	}
	return sum.Scale(1 / float64(len(t.Points)))
}

func (t Trace) IsTap(maxDistance float64) bool {
	if maxDistance < 0 {
		maxDistance = 0
	}
	return t.Distance() <= maxDistance && len(t.Points) <= 3
}

func (t Trace) IsSwipe(minDistance float64) bool {
	if minDistance < 0 {
		minDistance = 0
	}
	return t.Distance() >= minDistance && len(t.Points) >= 2
}

func (t Trace) Progress(point model.Vector) float64 {
	distance := t.Start.Distance(t.End)
	if distance == 0 {
		return 0
	}
	covered := t.Start.Distance(point)
	if covered < 0 {
		covered = 0
	}
	if covered > distance {
		covered = distance
	}
	return covered / distance
}
func abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}
