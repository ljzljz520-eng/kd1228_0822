package analysis

import (
	"jellyfield/app"
	"jellyfield/simulation"
	"testing"
)

func TestBuildReportRanksBrightness(t *testing.T) {
	scene := app.DemoScene()
	engine, err := simulation.NewEngine(scene)
	if err != nil {
		t.Fatal(err)
	}
	frame := engine.Step()
	report := BuildReport(scene, frame, engine.Measure())
	if report.Total != 3 || len(report.Entries) != 3 {
		t.Fatal(report)
	}
	if report.Brightest == "" || !report.AcceptableQuality() {
		t.Fatalf("report=%+v", report)
	}
}

func (r SceneReport) AcceptableQuality() bool { return r.Total > 0 && r.MeanEnergy > 0 }
