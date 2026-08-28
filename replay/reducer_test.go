package replay

import (
	"jellyfield/app"
	"jellyfield/model"
	"testing"
)

func TestReducerReplaysControls(t *testing.T) {
	state := NewState(app.DemoScene())
	events := []model.ControlEvent{{ID: "c1", TargetID: "jelly-a", Field: "speed", Value: 3.1, Sequence: 1}, {ID: "c2", TargetID: "jelly-b", Field: "pulse", Value: .2, Sequence: 2}}
	if state.ApplyAll(events) != 2 {
		t.Fatal("events")
	}
	speed, _ := state.Speed("jelly-a")
	if speed != 3.1 {
		t.Fatal(speed)
	}
	pulse, _ := state.Pulse("jelly-b")
	if pulse != .2 {
		t.Fatal(pulse)
	}
}
