package control

import "jellyfield/model"

type Filter struct {
	MinLuma      float64
	MaxSpeed     float64
	OnlySelected bool
}

func (f Filter) Apply(scene model.SceneState, selected string) []model.Jellyfish {
	result := []model.Jellyfish{}
	for _, j := range scene.Jellyfish {
		if f.OnlySelected && j.ID != selected {
			continue
		}
		if f.MinLuma > 0 && j.BaseColor.Luma() < f.MinLuma {
			continue
		}
		if f.MaxSpeed > 0 && j.Speed > f.MaxSpeed {
			continue
		}
		result = append(result, j)
	}
	return result
}
func SortBySpeed(items []model.Jellyfish, descending bool) []model.Jellyfish {
	result := append([]model.Jellyfish(nil), items...)
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			swap := result[j].Speed < result[i].Speed
			if descending {
				swap = !swap
			}
			if swap {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}
