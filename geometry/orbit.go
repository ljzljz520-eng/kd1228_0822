package geometry

import (
	"jellyfield/model"
	"math"
)

type Orbit struct {
	Center       model.Vector
	Radius       float64
	Phase        float64
	AngularSpeed float64
}

func (o Orbit) PositionAt(frame uint64) model.Vector {
	angle := o.Phase + float64(frame)*o.AngularSpeed
	return model.Vector{X: o.Center.X + math.Cos(angle)*o.Radius, Y: o.Center.Y + math.Sin(angle)*o.Radius}
}

func (o Orbit) TangentAt(frame uint64) model.Vector {
	angle := o.Phase + float64(frame)*o.AngularSpeed
	return model.Vector{X: -math.Sin(angle) * o.AngularSpeed * o.Radius, Y: math.Cos(angle) * o.AngularSpeed * o.Radius}
}

func Interpolate(a, b model.Vector, amount float64) model.Vector {
	if amount < 0 {
		amount = 0
	}
	if amount > 1 {
		amount = 1
	}
	return model.Vector{X: a.X + (b.X-a.X)*amount, Y: a.Y + (b.Y-a.Y)*amount}
}
func Spiral(point, center model.Vector, turns float64) model.Vector {
	dx, dy := point.X-center.X, point.Y-center.Y
	angle := turns * math.Pi * 2
	c, s := math.Cos(angle), math.Sin(angle)
	return center.Add(model.Vector{X: dx*c - dy*s, Y: dx*s + dy*c})
}
