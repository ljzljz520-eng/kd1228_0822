package simulation_test

import (
	"jellyfield/app"
	"jellyfield/simulation"
	"testing"
)

func TestEngineStepAndMetrics(t *testing.T) {
	engine, err := simulation.NewEngine(app.DemoScene())
	if err != nil {
		t.Fatal(err)
	}
	first := engine.Step()
	if first.Frame != 1 {
		t.Fatal(first.Frame)
	}
	metrics := engine.Measure()
	if metrics.ParticleCount != 44 {
		t.Fatal(metrics.ParticleCount)
	}
	if metrics.Brightness <= 0 {
		t.Fatal(metrics.Brightness)
	}
}
func TestEngineRejectsInvalidSpeed(t *testing.T) {
	engine, err := simulation.NewEngine(app.DemoScene())
	if err != nil {
		t.Fatal(err)
	}
	if err := engine.SetSpeed("jelly-a", 9); err == nil {
		t.Fatal("expected validation")
	}
}
