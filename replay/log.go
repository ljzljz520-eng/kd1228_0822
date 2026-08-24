package replay

import (
	"fmt"
	"jellyfield/model"
	"sort"
)

type Log struct {
	gestures []model.GestureRecord
	controls []model.ControlEvent
}

func NewLog() *Log { return &Log{gestures: []model.GestureRecord{}, controls: []model.ControlEvent{}} }
func (l *Log) AddGesture(record model.GestureRecord) error {
	if record.ID == "" {
		return fmt.Errorf("gesture id required")
	}
	l.gestures = append(l.gestures, record)
	return nil
}
func (l *Log) AddControl(event model.ControlEvent) error {
	if event.ID == "" {
		return fmt.Errorf("control id required")
	}
	l.controls = append(l.controls, event)
	return nil
}
func (l *Log) Gestures() []model.GestureRecord {
	result := append([]model.GestureRecord(nil), l.gestures...)
	sort.Slice(result, func(i, j int) bool { return result[i].Sequence < result[j].Sequence })
	return result
}
func (l *Log) Controls() []model.ControlEvent {
	result := append([]model.ControlEvent(nil), l.controls...)
	sort.Slice(result, func(i, j int) bool { return result[i].Sequence < result[j].Sequence })
	return result
}
func (l *Log) LastSequence() uint64 {
	last := uint64(0)
	for _, record := range l.gestures {
		if record.Sequence > last {
			last = record.Sequence
		}
	}
	for _, event := range l.controls {
		if event.Sequence > last {
			last = event.Sequence
		}
	}
	return last
}
func (l *Log) ByTarget(id string) []model.ControlEvent {
	result := []model.ControlEvent{}
	for _, event := range l.controls {
		if event.TargetID == id {
			result = append(result, event)
		}
	}
	return result
}
