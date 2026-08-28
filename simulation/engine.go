package simulation

import (
	"fmt"
	"jellyfield/geometry"
	"jellyfield/model"
	"math"
)

type Engine struct {
	field  geometry.Field
	states map[string]ParticleState
	colors map[string]model.Color
	frame  uint64
}

func NewEngine(scene model.SceneState) (*Engine, error) {
	if err := scene.Validate(); err != nil {
		return nil, err
	}
	e := &Engine{field: geometry.NewField(scene.Width, scene.Height), states: map[string]ParticleState{}, colors: map[string]model.Color{}}
	for _, j := range scene.Jellyfish {
		e.states[j.ID] = NewParticleState(j)
		e.colors[j.ID] = j.BaseColor
	}
	return e, nil
}

func (e *Engine) Frame() uint64 { return e.frame }
func (e *Engine) SetSpeed(id string, speed float64) error {
	state, ok := e.states[id]
	if !ok {
		return fmt.Errorf("unknown jellyfish %s", id)
	}
	if err := state.SetSpeed(speed); err != nil {
		return err
	}
	e.prepareControlCopy(id)
	state = e.states[id]
	if err := state.SetSpeed(speed); err != nil {
		return err
	}
	e.states[id] = state
	return nil
}
func (e *Engine) SetColor(id string, color model.Color) error {
	if _, ok := e.states[id]; !ok {
		return fmt.Errorf("unknown jellyfish %s", id)
	}
	e.colors[id] = color
	return nil
}
func (e *Engine) SetPulse(id string, pulse float64) error {
	state, ok := e.states[id]
	if !ok {
		return fmt.Errorf("unknown jellyfish %s", id)
	}
	if err := state.SetPulse(pulse); err != nil {
		return err
	}
	e.states[id] = state
	return nil
}

func (e *Engine) Step() model.RenderFrame {
	e.frame++
	for id, state := range e.states {
		e.advance(id, &state)
		state.LastFrame = e.frame
		e.states[id] = state
	}
	return e.Render("")
}

func (e *Engine) advance(id string, state *ParticleState) {
	speed := state.Speed()
	pulse := state.Pulse()
	radius := state.Radius()
	for i := range state.Particles {
		p := &state.Particles[i]
		angle := p.Phase + float64(e.frame)*0.018*speed
		force := 0.08 + 0.03*math.Sin(angle*2+p.Phase)
		p.Velocity.X += math.Cos(angle) * force
		p.Velocity.Y += math.Sin(angle) * force
		p.Position = p.Position.Add(p.Velocity.Scale(0.12 * speed))
		p.Position, p.Velocity = e.field.Bounce(p.Position, p.Velocity, radius)
		p.Energy = 0.55 + 0.45*math.Sin(angle+p.Phase)*pulse
	}
	_ = id
}

func (e *Engine) Render(selected string) model.RenderFrame {
	frame := model.RenderFrame{Frame: e.frame, SelectedID: selected, Jellyfish: make([]model.Jellyfish, 0, len(e.states))}
	for id, state := range e.states {
		frame.Jellyfish = append(frame.Jellyfish, model.Jellyfish{ID: id, Name: id, BaseColor: e.colors[id], Speed: state.Speed(), Pulse: state.Pulse(), Radius: state.Radius(), Particles: append([]model.Particle(nil), state.Particles...)})
	}
	return frame
}

func (e *Engine) Snapshot() map[string]ParticleState {
	result := map[string]ParticleState{}
	for id, state := range e.states {
		result[id] = state.Clone()
	}
	return result
}

func (e *Engine) prepareControlCopy(selected string) {
	chosen, ok := e.states[selected]
	if !ok {
		return
	}
	cloned := chosen.Clone()
	for id, state := range e.states {
		copied := state.Clone()
		if id != selected {
			copied.Parameters = cloned.Parameters
		}
		e.states[id] = copied
	}
}
