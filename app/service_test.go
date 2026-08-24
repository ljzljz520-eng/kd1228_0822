package app

import (
	"jellyfield/model"
	"jellyfield/persistence"
	"testing"
)

func TestServiceRestore(t *testing.T) {
	store, err := persistence.Open(t.TempDir() + "/service.db")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	service, err := New(DemoScene(), store)
	if err != nil {
		t.Fatal(err)
	}
	if err := service.Save(); err != nil {
		t.Fatal(err)
	}
	if err := service.SetColorAt(model.Vector{X: 220, Y: 210}, model.NewColor(1, 2, 3)); err != nil {
		t.Fatal(err)
	}
	if err := service.Restore(); err != nil {
		t.Fatal(err)
	}
	if service.Scene().ID != "reef-demo" {
		t.Fatal(service.Scene().ID)
	}
}
