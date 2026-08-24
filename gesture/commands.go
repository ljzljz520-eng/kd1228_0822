package gesture

import "jellyfield/model"

type Translator struct{ next uint64 }

func NewTranslator() *Translator { return &Translator{} }
func (t *Translator) Translate(g model.Gesture) (model.ControlCommand, model.GestureRecord, error) {
	if err := g.Validate(); err != nil {
		return model.ControlCommand{}, model.GestureRecord{}, err
	}
	t.next++
	record := model.GestureRecord{ID: formatID(t.next), TargetID: g.TargetID, Kind: string(g.Kind), Amount: g.Amount, Color: g.Color, Sequence: t.next}
	command := model.ControlCommand{TargetID: g.TargetID}
	switch g.Kind {
	case model.GestureSpeed:
		command.Speed = &g.Amount
	case model.GestureColor:
		command.Color = &g.Color
	case model.GestureSelect:
		command.Select = true
	case model.GestureExpand:
		pulse := 1.0
		command.Pulse = &pulse
	case model.GestureContract:
		pulse := 0.15
		command.Pulse = &pulse
	}
	return command, record, nil
}
func formatID(sequence uint64) string {
	digits := ""
	if sequence == 0 {
		return "gesture-0"
	}
	for sequence > 0 {
		digits = string(byte('0'+sequence%10)) + digits
		sequence /= 10
	}
	return "gesture-" + digits
}
