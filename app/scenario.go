package app

import (
	"jellyfield/model"
)

func DemoScene() model.SceneState {
	scene := model.NewScene("reef-demo", 800, 500)
	scene.Jellyfish = []model.Jellyfish{model.NewJellyfish("jelly-a", "Aurora", model.Vector{X: 220, Y: 210}, model.NewColor(40, 220, 255), 1.2, 14), model.NewJellyfish("jelly-b", "Lumen", model.Vector{X: 470, Y: 270}, model.NewColor(160, 90, 255), 2.4, 18), model.NewJellyfish("jelly-c", "Pulse", model.Vector{X: 640, Y: 160}, model.NewColor(90, 255, 180), 0.8, 12)}
	scene.SelectedID = "jelly-a"
	return scene
}
func SeedStore(service *Service) error {
	if err := service.Save(); err != nil {
		return err
	}
	_, err := service.HandleGesture(model.Gesture{Kind: model.GestureSelect, TargetID: "jelly-a"})
	return err
}
