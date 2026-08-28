package model

import "fmt"

type GestureKind string

const (
	GestureSelect   GestureKind = "select"
	GestureSpeed    GestureKind = "speed"
	GestureColor    GestureKind = "color"
	GestureExpand   GestureKind = "expand"
	GestureContract GestureKind = "contract"
)

type Gesture struct {
	Kind     GestureKind
	TargetID string
	X, Y     float64
	Amount   float64
	Color    Color
}

func (g Gesture) Validate() error {
	if g.TargetID == "" {
		return fmt.Errorf("gesture target is required")
	}
	switch g.Kind {
	case GestureSelect, GestureSpeed, GestureColor, GestureExpand, GestureContract:
		return nil
	default:
		return fmt.Errorf("unsupported gesture %q", g.Kind)
	}
}

type ControlCommand struct {
	TargetID string
	Speed    *float64
	Color    *Color
	Pulse    *float64
	Select   bool
}

func (c ControlCommand) Validate() error {
	if c.TargetID == "" {
		return fmt.Errorf("control target is required")
	}
	if c.Speed != nil && (*c.Speed < 0.1 || *c.Speed > 8) {
		return fmt.Errorf("speed outside range")
	}
	if c.Pulse != nil && (*c.Pulse < 0 || *c.Pulse > 1) {
		return fmt.Errorf("pulse outside range")
	}
	return nil
}

type Query struct {
	SceneID     string
	JellyfishID string
	MinLuma     float64
	MaxSpeed    float64
}

func (q Query) Matches(j Jellyfish) bool {
	if q.JellyfishID != "" && q.JellyfishID != j.ID {
		return false
	}
	if q.MinLuma > 0 && j.BaseColor.Luma() < q.MinLuma {
		return false
	}
	if q.MaxSpeed > 0 && j.Speed > q.MaxSpeed {
		return false
	}
	return true
}
