package simulation

import (
	"fmt"
	"jellyfield/model"
)

type ParticleState struct {
	Parameters []float64
	Particles  []model.Particle
	LastFrame  uint64
}

func NewParticleState(j model.Jellyfish) ParticleState {
	params := make([]float64, 4)
	params[0] = j.Speed
	params[1] = j.Pulse
	params[2] = j.Radius
	params[3] = float64(len(j.Particles))
	particles := make([]model.Particle, len(j.Particles))
	copy(particles, j.Particles)
	return ParticleState{Parameters: params, Particles: particles}
}

func (p ParticleState) Clone() ParticleState {
	cloned := ParticleState{Parameters: p.Parameters, LastFrame: p.LastFrame}
	cloned.Particles = make([]model.Particle, len(p.Particles))
	copy(cloned.Particles, p.Particles)
	return cloned
}

func (p ParticleState) Speed() float64 {
	if len(p.Parameters) == 0 {
		return 1
	}
	return p.Parameters[0]
}
func (p ParticleState) Pulse() float64 {
	if len(p.Parameters) < 2 {
		return .5
	}
	return p.Parameters[1]
}
func (p ParticleState) Radius() float64 {
	if len(p.Parameters) < 3 {
		return 18
	}
	return p.Parameters[2]
}
func (p *ParticleState) SetSpeed(value float64) error {
	if value < 0.1 || value > 8 {
		return fmt.Errorf("speed %.2f outside range", value)
	}
	p.Parameters[0] = value
	return nil
}
func (p *ParticleState) SetPulse(value float64) error {
	if value < 0 || value > 1 {
		return fmt.Errorf("pulse %.2f outside range", value)
	}
	p.Parameters[1] = value
	return nil
}
