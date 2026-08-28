package replay

import "jellyfield/model"

type State struct {
	Speeds   map[string]float64
	Colors   map[string]model.Color
	Pulses   map[string]float64
	Selected string
}

func NewState(scene model.SceneState) State {
	state := State{Speeds: map[string]float64{}, Colors: map[string]model.Color{}, Pulses: map[string]float64{}}
	for _, j := range scene.Jellyfish {
		state.Speeds[j.ID] = j.Speed
		state.Colors[j.ID] = j.BaseColor
		state.Pulses[j.ID] = j.Pulse
	}
	state.Selected = scene.SelectedID
	return state
}
func (s *State) Apply(event model.ControlEvent) bool {
	switch event.Field {
	case "speed":
		s.Speeds[event.TargetID] = event.Value
	case "color":
		s.Colors[event.TargetID] = event.Color
	case "pulse":
		s.Pulses[event.TargetID] = event.Value
	default:
		return false
	}
	return true
}
func (s State) Speed(id string) (float64, bool)     { value, ok := s.Speeds[id]; return value, ok }
func (s State) Color(id string) (model.Color, bool) { value, ok := s.Colors[id]; return value, ok }
func (s State) Pulse(id string) (float64, bool)     { value, ok := s.Pulses[id]; return value, ok }
func (s State) ApplyAll(events []model.ControlEvent) int {
	applied := 0
	for _, event := range events {
		if s.Apply(event) {
			applied++
		}
	}
	return applied
}
func (s State) Snapshot() State {
	copyState := State{Speeds: map[string]float64{}, Colors: map[string]model.Color{}, Pulses: map[string]float64{}, Selected: s.Selected}
	for id, value := range s.Speeds {
		copyState.Speeds[id] = value
	}
	for id, value := range s.Colors {
		copyState.Colors[id] = value
	}
	for id, value := range s.Pulses {
		copyState.Pulses[id] = value
	}
	return copyState
}
