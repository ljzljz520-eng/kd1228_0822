package presentation

import "jellyfield/model"

type Palette struct {
	Name   string
	Colors []model.Color
}

func DefaultPalettes() []Palette {
	return []Palette{{Name: "reef", Colors: []model.Color{{R: 40, G: 220, B: 255}, {R: 130, G: 80, B: 255}, {R: 90, G: 255, B: 180}}}, {Name: "sunset", Colors: []model.Color{{R: 255, G: 100, B: 80}, {R: 255, G: 190, B: 60}, {R: 230, G: 80, B: 180}}}, {Name: "moon", Colors: []model.Color{{R: 180, G: 210, B: 255}, {R: 120, G: 150, B: 210}, {R: 230, G: 240, B: 255}}}}
}
func FindPalette(name string) (Palette, bool) {
	for _, p := range DefaultPalettes() {
		if p.Name == name {
			return p, true
		}
	}
	return Palette{}, false
}
func (p Palette) Color(index int) model.Color {
	if len(p.Colors) == 0 {
		return model.Color{}
	}
	if index < 0 {
		index = 0
	}
	return p.Colors[index%len(p.Colors)]
}
