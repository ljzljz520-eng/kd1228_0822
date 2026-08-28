package control_test

import (
	"jellyfield/app"
	"jellyfield/control"
	"jellyfield/model"
	"jellyfield/simulation"
	"testing"
)

func TestPanelSelectionAndApply(t *testing.T) {
	engine, err := simulation.NewEngine(app.DemoScene())
	if err != nil {
		t.Fatal(err)
	}
	panel := control.NewPanel(engine)
	speed := 2.5
	event, err := panel.Apply(model.ControlCommand{TargetID: "jelly-b", Speed: &speed, Select: true})
	if err != nil {
		t.Fatal(err)
	}
	if panel.Selected() != "jelly-b" || event.Field != "speed" {
		t.Fatalf("panel=%s event=%+v", panel.Selected(), event)
	}
}
