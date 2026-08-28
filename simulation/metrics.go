package simulation

import "jellyfield/model"

type Metrics struct {
	Frame         uint64
	AverageSpeed  float64
	ParticleCount int
	Brightness    float64
}

func (e *Engine) Measure() Metrics {
	result := Metrics{Frame: e.frame}
	if len(e.states) == 0 {
		return result
	}
	for _, state := range e.states {
		result.AverageSpeed += state.Speed()
		result.ParticleCount += len(state.Particles)
		for _, p := range state.Particles {
			result.Brightness += p.Energy
		}
	}
	result.AverageSpeed /= float64(len(e.states))
	if result.ParticleCount > 0 {
		result.Brightness /= float64(result.ParticleCount)
	}
	return result
}

func (e *Engine) ApplyJellyfish(j model.Jellyfish) error {
	if err := j.Validate(); err != nil {
		return err
	}
	if _, ok := e.states[j.ID]; !ok {
		e.states[j.ID] = NewParticleState(j)
	}
	e.colors[j.ID] = j.BaseColor
	return nil
}
