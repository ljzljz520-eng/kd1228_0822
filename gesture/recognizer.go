package gesture

import (
	"fmt"
	"jellyfield/geometry"
	"jellyfield/model"
)

type Recognizer struct {
	field     geometry.Field
	threshold float64
}

func NewRecognizer(width, height float64) Recognizer {
	return Recognizer{field: geometry.NewField(width, height), threshold: 32}
}
func (r Recognizer) Hit(scene model.SceneState, point model.Vector) (string, bool) {
	best := ""
	distance := r.threshold
	for _, j := range scene.Jellyfish {
		d := j.Position.Distance(point)
		if d <= distance {
			best = j.ID
			distance = d
		}
	}
	return best, best != ""
}
func (r Recognizer) Select(scene model.SceneState, point model.Vector) (model.Gesture, error) {
	id, ok := r.Hit(scene, point)
	if !ok {
		return model.Gesture{}, fmt.Errorf("no jellyfish at %.1f,%.1f", point.X, point.Y)
	}
	return model.Gesture{Kind: model.GestureSelect, TargetID: id, X: point.X, Y: point.Y}, nil
}
func (r Recognizer) Speed(scene model.SceneState, point model.Vector, amount float64) (model.Gesture, error) {
	id, ok := r.Hit(scene, point)
	if !ok {
		return model.Gesture{}, fmt.Errorf("speed gesture has no target")
	}
	return model.Gesture{Kind: model.GestureSpeed, TargetID: id, Amount: geometry.Denormalize(amount, 0.1, 8)}, nil
}
func (r Recognizer) Color(scene model.SceneState, point model.Vector, color model.Color) (model.Gesture, error) {
	id, ok := r.Hit(scene, point)
	if !ok {
		return model.Gesture{}, fmt.Errorf("color gesture has no target")
	}
	return model.Gesture{Kind: model.GestureColor, TargetID: id, Color: color}, nil
}
