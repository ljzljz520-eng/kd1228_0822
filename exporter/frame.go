package exporter

import (
	"encoding/json"
	"fmt"
	"jellyfield/model"
	"sort"
	"strings"
)

type Manifest struct {
	SceneID   string   `json:"scene_id"`
	Frame     uint64   `json:"frame"`
	Selected  string   `json:"selected"`
	Jellyfish int      `json:"jellyfish"`
	Particles int      `json:"particles"`
	Colors    []string `json:"colors"`
}

func BuildManifest(frame model.RenderFrame) Manifest {
	manifest := Manifest{SceneID: frame.SceneID, Frame: frame.Frame, Selected: frame.SelectedID, Jellyfish: len(frame.Jellyfish), Colors: []string{}}
	seen := map[string]bool{}
	for _, j := range frame.Jellyfish {
		manifest.Particles += len(j.Particles)
		color := j.BaseColor.Hex()
		if !seen[color] {
			seen[color] = true
			manifest.Colors = append(manifest.Colors, color)
		}
	}
	sort.Strings(manifest.Colors)
	return manifest
}
func EncodeManifest(manifest Manifest) (string, error) {
	data, err := json.MarshalIndent(manifest, "", "  ")
	return string(data), err
}
func FrameCSV(frame model.RenderFrame) string {
	lines := []string{"id,name,x,y,speed,color,particles"}
	for _, j := range frame.Jellyfish {
		lines = append(lines, fmt.Sprintf("%s,%s,%.2f,%.2f,%.2f,%s,%d", j.ID, strings.ReplaceAll(j.Name, ",", " "), j.Position.X, j.Position.Y, j.Speed, j.BaseColor.Hex(), len(j.Particles)))
	}
	return strings.Join(lines, "\n") + "\n"
}
func Speeds(frame model.RenderFrame) map[string]float64 {
	result := map[string]float64{}
	for _, j := range frame.Jellyfish {
		result[j.ID] = j.Speed
	}
	return result
}
func Colors(frame model.RenderFrame) map[string]string {
	result := map[string]string{}
	for _, j := range frame.Jellyfish {
		result[j.ID] = j.BaseColor.Hex()
	}
	return result
}
