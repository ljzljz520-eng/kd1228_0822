package layout

import (
	"jellyfield/model"
	"math"
	"sort"
)

type Layout struct {
	Width     float64
	Height    float64
	Positions map[string]model.Vector
}

func Arrange(scene model.SceneState) Layout {
	layout := Layout{Width: scene.Width, Height: scene.Height, Positions: map[string]model.Vector{}}
	for _, j := range scene.Jellyfish {
		layout.Positions[j.ID] = j.Position
	}
	return layout
}
func (l Layout) Move(id string, position model.Vector) bool {
	if _, ok := l.Positions[id]; !ok {
		return false
	}
	l.Positions[id] = position
	return true
}
func (l Layout) Center() model.Vector { return model.Vector{X: l.Width / 2, Y: l.Height / 2} }
func (l Layout) DistanceToCenter(id string) float64 {
	point, ok := l.Positions[id]
	if !ok {
		return math.Inf(1)
	}
	return point.Distance(l.Center())
}
func (l Layout) Nearest(point model.Vector) (string, bool) {
	ids := make([]string, 0, len(l.Positions))
	for id := range l.Positions {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	best := ""
	distance := math.Inf(1)
	for _, id := range ids {
		current := l.Positions[id].Distance(point)
		if current < distance {
			distance = current
			best = id
		}
	}
	return best, best != ""
}
func (l Layout) Ordered() []string {
	ids := make([]string, 0, len(l.Positions))
	for id := range l.Positions {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
func (l Layout) InBounds(point model.Vector) bool {
	return point.X >= 0 && point.Y >= 0 && point.X <= l.Width && point.Y <= l.Height
}
func (l Layout) Normalize(point model.Vector) model.Vector {
	if l.Width <= 0 || l.Height <= 0 {
		return model.Vector{}
	}
	return model.Vector{X: point.X / l.Width, Y: point.Y / l.Height}
}
func (l Layout) Denormalize(point model.Vector) model.Vector {
	return model.Vector{X: point.X * l.Width, Y: point.Y * l.Height}
}
