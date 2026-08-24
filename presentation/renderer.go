package presentation

import (
	"encoding/json"
	"fmt"
	"jellyfield/model"
	"strings"
)

type Renderer struct{ Compact bool }

func NewRenderer(compact bool) Renderer { return Renderer{Compact: compact} }
func (r Renderer) JSON(frame model.RenderFrame) (string, error) {
	var data []byte
	var err error
	if r.Compact {
		data, err = json.Marshal(frame)
	} else {
		data, err = json.MarshalIndent(frame, "", "  ")
	}
	return string(data), err
}
func (r Renderer) Summary(frame model.RenderFrame) string {
	lines := []string{fmt.Sprintf("scene frame %d selected=%s", frame.Frame, frame.SelectedID)}
	for _, j := range frame.Jellyfish {
		lines = append(lines, fmt.Sprintf("%s speed=%.2f color=%s particles=%d", j.ID, j.Speed, j.BaseColor.Hex(), len(j.Particles)))
	}
	return strings.Join(lines, "\n")
}
func (r Renderer) Table(frame model.RenderFrame) [][]string {
	rows := [][]string{{"id", "speed", "color", "particles"}}
	for _, j := range frame.Jellyfish {
		rows = append(rows, []string{j.ID, fmt.Sprintf("%.2f", j.Speed), j.BaseColor.Hex(), fmt.Sprintf("%d", len(j.Particles))})
	}
	return rows
}
