package catalog

import "testing"

func TestDefaultsBuildScene(t *testing.T) {
	catalog := Default()
	if len(catalog.IDs()) != 4 {
		t.Fatal(catalog.IDs())
	}
	scene := catalog.BuildScene("catalog-scene", 700, 400)
	if len(scene.Jellyfish) != 4 {
		t.Fatal(len(scene.Jellyfish))
	}
	if speed, ok := catalog.SpeedPreset("lumen", "fast"); !ok || speed <= 2.4 {
		t.Fatal(speed, ok)
	}
	if len(catalog.FindTag("cool")) != 2 {
		t.Fatal("tag")
	}
}
