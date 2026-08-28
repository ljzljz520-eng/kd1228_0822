package control

import (
	"fmt"
	"jellyfield/model"
	"jellyfield/simulation"
)

type Panel struct {
	engine   *simulation.Engine
	selected string
	sequence uint64
}

func NewPanel(engine *simulation.Engine) *Panel { return &Panel{engine: engine} }
func (p *Panel) Selected() string               { return p.selected }
func (p *Panel) Apply(command model.ControlCommand) (model.ControlEvent, error) {
	if err := command.Validate(); err != nil {
		return model.ControlEvent{}, err
	}
	if command.Select {
		p.selected = command.TargetID
	}
	if p.selected == "" {
		p.selected = command.TargetID
	}
	if p.selected != command.TargetID {
		return model.ControlEvent{}, fmt.Errorf("target %s is not selected", command.TargetID)
	}
	p.sequence++
	event := model.ControlEvent{ID: fmt.Sprintf("control-%d", p.sequence), TargetID: command.TargetID, Sequence: p.sequence}
	if command.Speed != nil {
		if err := p.engine.SetSpeed(command.TargetID, *command.Speed); err != nil {
			return model.ControlEvent{}, err
		}
		event.Field = "speed"
		event.Value = *command.Speed
	}
	if command.Color != nil {
		if err := p.engine.SetColor(command.TargetID, *command.Color); err != nil {
			return model.ControlEvent{}, err
		}
		event.Field = "color"
		event.Color = *command.Color
	}
	if command.Pulse != nil {
		if err := p.engine.SetPulse(command.TargetID, *command.Pulse); err != nil {
			return model.ControlEvent{}, err
		}
		event.Field = "pulse"
		event.Value = *command.Pulse
	}
	return event, nil
}

func (p *Panel) ResetSelection() { p.selected = "" }
func (p *Panel) SetSelected(id string) error {
	if id == "" {
		return fmt.Errorf("selection cannot be empty")
	}
	p.selected = id
	return nil
}
