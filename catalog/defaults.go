package catalog

import "jellyfield/model"

func Default() *Catalog {
	catalog := New()
	entries := []Entry{{ID: "aurora", Label: "Aurora", Description: "cyan bell with a long drifting skirt", Color: model.NewColor(40, 220, 255), DefaultSpeed: 1.2, ParticleCount: 14, Tags: []string{"cool", "calm"}}, {ID: "lumen", Label: "Lumen", Description: "violet bell with a bright pulse", Color: model.NewColor(160, 90, 255), DefaultSpeed: 2.4, ParticleCount: 18, Tags: []string{"bright", "fast"}}, {ID: "pulse", Label: "Pulse", Description: "mint bell tuned for gentle movement", Color: model.NewColor(90, 255, 180), DefaultSpeed: .8, ParticleCount: 12, Tags: []string{"cool", "gentle"}}, {ID: "ember", Label: "Ember", Description: "warm coral bell for contrast", Color: model.NewColor(255, 100, 80), DefaultSpeed: 1.6, ParticleCount: 16, Tags: []string{"warm", "bright"}}}
	for _, entry := range entries {
		_ = catalog.Add(entry)
	}
	return catalog
}
func (c *Catalog) SpeedPreset(id string, level string) (float64, bool) {
	entry, ok := c.Get(id)
	if !ok {
		return 0, false
	}
	switch level {
	case "slow":
		return entry.DefaultSpeed * .5, true
	case "normal":
		return entry.DefaultSpeed, true
	case "fast":
		return min(entry.DefaultSpeed*1.5, 8), true
	default:
		return 0, false
	}
}
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func ColorNames() map[string]model.Color {
	return map[string]model.Color{"cyan": model.NewColor(40, 220, 255), "violet": model.NewColor(160, 90, 255), "mint": model.NewColor(90, 255, 180), "coral": model.NewColor(255, 100, 80)}
}
