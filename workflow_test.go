package jellyfield

import (
	"jellyfield/app"
	"jellyfield/model"
	"jellyfield/persistence"
	"testing"
)

func openService(t *testing.T) (*app.Service, *persistence.Store) {
	t.Helper()
	store, err := persistence.Open(t.TempDir() + "/scene.db")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { store.Close() })
	service, err := app.New(app.DemoScene(), store)
	if err != nil {
		t.Fatal(err)
	}
	return service, store
}

func TestWorkflowOne(t *testing.T) {
	service, store := openService(t)
	if err := service.Save(); err != nil {
		t.Fatal(err)
	}
	if err := service.SelectAt(model.Vector{X: 220, Y: 210}); err != nil {
		t.Fatal(err)
	}
	if err := service.SetSpeedAt(model.Vector{X: 220, Y: 210}, 0.5); err != nil {
		t.Fatal(err)
	}
	frame := service.Step()
	if frame.Frame != 1 {
		t.Fatalf("frame=%d", frame.Frame)
	}
	if _, err := store.LoadScene("reef-demo"); err != nil {
		t.Fatal(err)
	}
}

func TestWorkflowTwo(t *testing.T) {
	service, _ := openService(t)
	if err := service.SetColorAt(model.Vector{X: 470, Y: 270}, model.NewColor(255, 30, 120)); err != nil {
		t.Fatal(err)
	}
	service.Step()
	rows := service.Query(model.Query{MinLuma: 100})
	if len(rows) < 1 {
		t.Fatal("expected bright jellyfish")
	}
	if service.Metrics().ParticleCount != 44 {
		t.Fatalf("particles=%d", service.Metrics().ParticleCount)
	}
}

func TestWorkflowThree(t *testing.T) {
	service, _ := openService(t)
	if err := service.SelectAt(model.Vector{X: 640, Y: 160}); err != nil {
		t.Fatal(err)
	}
	event, err := service.HandleGesture(model.Gesture{Kind: model.GestureExpand, TargetID: "jelly-c"})
	if err != nil {
		t.Fatal(err)
	}
	if event.Field != "pulse" || event.Value != 1 {
		t.Fatalf("event=%+v", event)
	}
	service.Step()
	if service.Scene().SelectedID != "jelly-c" {
		t.Fatal("selection missing")
	}
}
