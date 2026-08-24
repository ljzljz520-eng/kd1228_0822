package exporter

import (
	"jellyfield/app"
	"jellyfield/simulation"
	"testing"
)

func TestManifestAndCSV(t *testing.T) {
	engine, err := simulation.NewEngine(app.DemoScene())
	if err != nil {
		t.Fatal(err)
	}
	frame := engine.Step()
	frame.SceneID = "reef-demo"
	manifest := BuildManifest(frame)
	if manifest.Jellyfish != 3 || manifest.Particles != 44 {
		t.Fatal(manifest)
	}
	if len(FrameCSV(frame)) < 40 {
		t.Fatal("csv too short")
	}
	if err := ValidateFrame(frame).Require(); err != nil {
		t.Fatal(err)
	}
}
