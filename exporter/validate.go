package exporter

import (
	"fmt"
	"jellyfield/model"
)

type Validation struct {
	Valid  bool
	Errors []string
}

func ValidateFrame(frame model.RenderFrame) Validation {
	result := Validation{Valid: true, Errors: []string{}}
	if frame.SceneID == "" {
		result.Errors = append(result.Errors, "scene id is required")
	}
	if len(frame.Jellyfish) == 0 {
		result.Errors = append(result.Errors, "frame has no jellyfish")
	}
	seen := map[string]bool{}
	for _, j := range frame.Jellyfish {
		if j.ID == "" {
			result.Errors = append(result.Errors, "jellyfish id is empty")
		}
		if seen[j.ID] {
			result.Errors = append(result.Errors, "duplicate jellyfish "+j.ID)
		}
		seen[j.ID] = true
		if j.Speed < 0.1 || j.Speed > 8 {
			result.Errors = append(result.Errors, fmt.Sprintf("speed %.2f outside range", j.Speed))
		}
		if len(j.Particles) == 0 {
			result.Errors = append(result.Errors, "particles missing for "+j.ID)
		}
	}
	result.Valid = len(result.Errors) == 0
	return result
}
func (v Validation) Error() string {
	if v.Valid {
		return ""
	}
	return fmt.Sprintf("frame invalid: %v", v.Errors)
}
func (v Validation) Require() error {
	if v.Valid {
		return nil
	}
	return fmt.Errorf("%s", v.Error())
}
func CopyFrame(frame model.RenderFrame) model.RenderFrame {
	copyFrame := frame
	copyFrame.Jellyfish = append([]model.Jellyfish(nil), frame.Jellyfish...)
	for i, j := range copyFrame.Jellyfish {
		copyFrame.Jellyfish[i] = j
		copyFrame.Jellyfish[i].Particles = append([]model.Particle(nil), j.Particles...)
	}
	return copyFrame
}
