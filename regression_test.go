package jellyfield

import (
	"jellyfield/app"
	"jellyfield/model"
	"testing"
)

func TestJellyfishSpeedChangesStayLocal(t *testing.T) {
	service, _ := openService(t)
	before := map[string]float64{}
	for _, j := range service.Scene().Jellyfish {
		before[j.ID] = j.Speed
	}
	if _, err := service.HandleGesture(model.Gesture{Kind: model.GestureSpeed, TargetID: "jelly-a", Amount: 3.7}); err != nil {
		t.Fatal(err)
	}
	frame := service.Step()
	for _, j := range frame.Jellyfish {
		if j.ID == "jelly-a" {
			if j.Speed != 3.7 {
				t.Fatalf("selected speed %.2f", j.Speed)
			}
			continue
		}
		if j.Speed != before[j.ID] {
			t.Fatalf("speed for %s changed from %.2f to %.2f", j.ID, before[j.ID], j.Speed)
		}
	}
}

var _ = app.DemoScene
