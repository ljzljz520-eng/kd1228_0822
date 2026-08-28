package gesture_test

import (
	"jellyfield/app"
	"jellyfield/gesture"
	"jellyfield/model"
	"testing"
)

func TestRecognizerTranslatesTargetedGestures(t *testing.T) {
	scene := app.DemoScene()
	recognizer := gesture.NewRecognizer(scene.Width, scene.Height)
	g, err := recognizer.Speed(scene, model.Vector{X: 470, Y: 270}, 0.25)
	if err != nil {
		t.Fatal(err)
	}
	translator := gesture.NewTranslator()
	command, record, err := translator.Translate(g)
	if err != nil {
		t.Fatal(err)
	}
	if command.TargetID != "jelly-b" || *command.Speed != 2.075 || record.Sequence != 1 {
		t.Fatalf("command=%+v record=%+v", command, record)
	}
}
