package geometry

import (
	"jellyfield/model"
	"math"
)

type Field struct {
	Width, Height float64
	Margin        float64
}

func NewField(width, height float64) Field {
	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}
	return Field{Width: width, Height: height, Margin: 24}
}

func (f Field) Min() model.Vector { return model.Vector{X: f.Margin, Y: f.Margin} }
func (f Field) Max() model.Vector { return model.Vector{X: f.Width - f.Margin, Y: f.Height - f.Margin} }
func (f Field) Contains(point model.Vector) bool {
	min, max := f.Min(), f.Max()
	return point.X >= min.X && point.X <= max.X && point.Y >= min.Y && point.Y <= max.Y
}
func (f Field) Clamp(point model.Vector) model.Vector { return point.Clamp(f.Min(), f.Max()) }

func (f Field) Bounce(position, velocity model.Vector, radius float64) (model.Vector, model.Vector) {
	min, max := model.Vector{X: f.Margin + radius, Y: f.Margin + radius}, model.Vector{X: f.Width - f.Margin - radius, Y: f.Height - f.Margin - radius}
	if min.X > max.X {
		min.X, max.X = f.Margin, f.Width-f.Margin
	}
	if min.Y > max.Y {
		min.Y, max.Y = f.Margin, f.Height-f.Margin
	}
	if position.X < min.X {
		position.X = min.X
		velocity.X = math.Abs(velocity.X)
	}
	if position.X > max.X {
		position.X = max.X
		velocity.X = -math.Abs(velocity.X)
	}
	if position.Y < min.Y {
		position.Y = min.Y
		velocity.Y = math.Abs(velocity.Y)
	}
	if position.Y > max.Y {
		position.Y = max.Y
		velocity.Y = -math.Abs(velocity.Y)
	}
	return position, velocity
}

func Normalize(value, low, high float64) float64 {
	if high <= low {
		return 0
	}
	result := (value - low) / (high - low)
	if result < 0 {
		return 0
	}
	if result > 1 {
		return 1
	}
	return result
}

func Denormalize(value, low, high float64) float64 {
	if value < 0 {
		value = 0
	}
	if value > 1 {
		value = 1
	}
	return low + (high-low)*value
}
