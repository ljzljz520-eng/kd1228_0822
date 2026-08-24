package persistence

import (
	"jellyfield/model"
	"testing"
)

func TestStoreListsEntities(t *testing.T) {
	store, err := Open(t.TempDir() + "/store.db")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	scene := model.NewScene("s", 120, 80)
	scene.Jellyfish = []model.Jellyfish{model.NewJellyfish("j", "J", model.Vector{X: 20, Y: 20}, model.NewColor(1, 2, 3), 1, 2)}
	if err := store.SaveScene(scene); err != nil {
		t.Fatal(err)
	}
	if ids, err := store.SceneIDs(); err != nil || len(ids) != 1 {
		t.Fatalf("ids=%v err=%v", ids, err)
	}
	if err := store.AppendGesture(model.GestureRecord{ID: "g", TargetID: "j", Kind: "select"}); err != nil {
		t.Fatal(err)
	}
	if ids, err := store.GestureIDs(); err != nil || len(ids) != 1 {
		t.Fatalf("gesture ids=%v err=%v", ids, err)
	}
}
