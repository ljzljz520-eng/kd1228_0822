package jellyfield

import (
	"jellyfield/app"
	"jellyfield/model"
	"jellyfield/persistence"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	path := t.TempDir() + "/reopen.db"
	store, err := persistence.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	service, err := app.New(app.DemoScene(), store)
	if err != nil {
		t.Fatal(err)
	}
	if err := service.Save(); err != nil {
		t.Fatal(err)
	}
	if err := store.AppendGesture(structuredGesture()); err != nil {
		t.Fatal(err)
	}
	if err := store.AppendControl(structuredControl()); err != nil {
		t.Fatal(err)
	}
	if err := store.Close(); err != nil {
		t.Fatal(err)
	}
	reopened, err := persistence.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer reopened.Close()
	scene, err := reopened.LoadScene("reef-demo")
	if err != nil {
		t.Fatal(err)
	}
	if len(scene.Jellyfish) != 3 {
		t.Fatalf("jellyfish=%d", len(scene.Jellyfish))
	}
	if _, err := reopened.LoadGesture("gesture-1"); err != nil {
		t.Fatal(err)
	}
	if _, err := reopened.LoadControl("control-1"); err != nil {
		t.Fatal(err)
	}
}

func structuredGesture() model.GestureRecord {
	return model.GestureRecord{ID: "gesture-1", TargetID: "jelly-a", Kind: "select", Sequence: 1}
}
func structuredControl() model.ControlEvent {
	return model.ControlEvent{ID: "control-1", TargetID: "jelly-a", Field: "speed", Value: 1.2, Sequence: 1}
}

var _ = persistence.EntityJellyfish
