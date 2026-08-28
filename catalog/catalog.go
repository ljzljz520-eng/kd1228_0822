package catalog

import (
	"fmt"
	"jellyfield/model"
	"sort"
)

type Entry struct {
	ID            string
	Label         string
	Description   string
	Color         model.Color
	DefaultSpeed  float64
	ParticleCount int
	Tags          []string
}
type Catalog struct{ entries map[string]Entry }

func New() *Catalog { return &Catalog{entries: map[string]Entry{}} }
func (c *Catalog) Add(entry Entry) error {
	if entry.ID == "" {
		return fmt.Errorf("catalog id required")
	}
	if entry.Label == "" {
		return fmt.Errorf("catalog label required")
	}
	if entry.DefaultSpeed < 0.1 || entry.DefaultSpeed > 8 {
		return fmt.Errorf("default speed outside range")
	}
	if entry.ParticleCount < 1 {
		return fmt.Errorf("particle count must be positive")
	}
	c.entries[entry.ID] = entry
	return nil
}
func (c *Catalog) Get(id string) (Entry, bool) { entry, ok := c.entries[id]; return entry, ok }
func (c *Catalog) Remove(id string) bool {
	if _, ok := c.entries[id]; !ok {
		return false
	}
	delete(c.entries, id)
	return true
}
func (c *Catalog) IDs() []string {
	ids := make([]string, 0, len(c.entries))
	for id := range c.entries {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
func (c *Catalog) All() []Entry {
	result := make([]Entry, 0, len(c.entries))
	for _, entry := range c.entries {
		result = append(result, entry)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result
}
func (c *Catalog) FindTag(tag string) []Entry {
	result := []Entry{}
	for _, entry := range c.entries {
		for _, candidate := range entry.Tags {
			if candidate == tag {
				result = append(result, entry)
				break
			}
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Label < result[j].Label })
	return result
}
func (c *Catalog) BuildScene(id string, width, height float64) model.SceneState {
	scene := model.NewScene(id, width, height)
	for index, entry := range c.All() {
		position := model.Vector{X: 80 + float64(index%4)*170, Y: 100 + float64(index/4)*140}
		scene.Jellyfish = append(scene.Jellyfish, model.NewJellyfish(entry.ID, entry.Label, position, entry.Color, entry.DefaultSpeed, entry.ParticleCount))
	}
	return scene
}
