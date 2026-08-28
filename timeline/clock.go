package timeline

import (
	"fmt"
	"jellyfield/model"
)

type Clock struct {
	frame  uint64
	step   uint64
	paused bool
	labels map[uint64]string
}

func New(step uint64) *Clock {
	if step == 0 {
		step = 1
	}
	return &Clock{step: step, labels: map[uint64]string{}}
}
func (c *Clock) Frame() uint64        { return c.frame }
func (c *Clock) Paused() bool         { return c.paused }
func (c *Clock) Toggle()              { c.paused = !c.paused }
func (c *Clock) SetPaused(value bool) { c.paused = value }
func (c *Clock) Advance() uint64 {
	if !c.paused {
		c.frame += c.step
	}
	return c.frame
}
func (c *Clock) AdvanceBy(count uint64) uint64 {
	for i := uint64(0); i < count; i++ {
		c.Advance()
	}
	return c.frame
}
func (c *Clock) Label(label string) error {
	if label == "" {
		return fmt.Errorf("label is empty")
	}
	c.labels[c.frame] = label
	return nil
}
func (c *Clock) LabelAt(frame uint64, label string) error {
	if label == "" {
		return fmt.Errorf("label is empty")
	}
	c.labels[frame] = label
	return nil
}
func (c *Clock) LabelFor(frame uint64) (string, bool) { label, ok := c.labels[frame]; return label, ok }
func (c *Clock) Reset()                               { c.frame = 0; c.paused = false; c.labels = map[uint64]string{} }
func (c *Clock) StepSize() uint64                     { return c.step }
func (c *Clock) SetStep(step uint64) {
	if step == 0 {
		step = 1
	}
	c.step = step
}
func (c *Clock) FramesUntil(target uint64) uint64 {
	if target <= c.frame {
		return 0
	}
	remaining := target - c.frame
	return (remaining + c.step - 1) / c.step
}

type Keyframe struct {
	Frame    uint64
	Position model.Vector
	Speed    float64
	Color    model.Color
}

func InterpolateKeyframes(a, b Keyframe, frame uint64) Keyframe {
	if b.Frame <= a.Frame {
		return a
	}
	amount := float64(frame-a.Frame) / float64(b.Frame-a.Frame)
	if amount < 0 {
		amount = 0
	}
	if amount > 1 {
		amount = 1
	}
	return Keyframe{Frame: frame, Position: model.Vector{X: a.Position.X + (b.Position.X-a.Position.X)*amount, Y: a.Position.Y + (b.Position.Y-a.Position.Y)*amount}, Speed: a.Speed + (b.Speed-a.Speed)*amount, Color: a.Color.Mix(b.Color, amount)}
}
func FindKeyframe(keyframes []Keyframe, frame uint64) (Keyframe, bool) {
	if len(keyframes) == 0 {
		return Keyframe{}, false
	}
	best := keyframes[0]
	for _, keyframe := range keyframes {
		if keyframe.Frame <= frame && keyframe.Frame >= best.Frame {
			best = keyframe
		}
	}
	return best, true
}
