package analysis

import (
	"jellyfield/model"
	"math"
)

type QualityScore struct {
	Balance   float64
	Diversity float64
	Stability float64
	Total     float64
}

func Score(scene model.SceneState, frame model.RenderFrame) QualityScore {
	balance := spacingScore(scene)
	diversity := colorDiversity(frame)
	stability := speedStability(frame)
	return QualityScore{Balance: balance, Diversity: diversity, Stability: stability, Total: (balance + diversity + stability) / 3}
}
func spacingScore(scene model.SceneState) float64 {
	if len(scene.Jellyfish) < 2 {
		return 1
	}
	sum := 0.0
	count := 0
	for i, a := range scene.Jellyfish {
		for j, b := range scene.Jellyfish {
			if j <= i {
				continue
			}
			distance := a.Position.Distance(b.Position)
			sum += math.Min(distance/180, 1)
			count++
		}
	}
	if count == 0 {
		return 1
	}
	return sum / float64(count)
}
func colorDiversity(frame model.RenderFrame) float64 {
	if len(frame.Jellyfish) < 2 {
		return 1
	}
	sum := 0.0
	count := 0
	for i, a := range frame.Jellyfish {
		for j, b := range frame.Jellyfish {
			if j <= i {
				continue
			}
			sum += math.Abs(a.BaseColor.Luma()-b.BaseColor.Luma()) / 255
			count++
		}
	}
	if count == 0 {
		return 1
	}
	return math.Min(sum/float64(count)*2, 1)
}
func speedStability(frame model.RenderFrame) float64 {
	if len(frame.Jellyfish) == 0 {
		return 0
	}
	mean := 0.0
	for _, j := range frame.Jellyfish {
		mean += j.Speed
	}
	mean /= float64(len(frame.Jellyfish))
	variance := 0.0
	for _, j := range frame.Jellyfish {
		delta := j.Speed - mean
		variance += delta * delta
	}
	variance /= float64(len(frame.Jellyfish))
	return 1 / (1 + variance)
}
func (q QualityScore) Acceptable() bool { return q.Total >= 0.5 && q.Balance >= 0.25 }
func (q QualityScore) Label() string {
	if q.Total >= 0.8 {
		return "luminous"
	}
	if q.Total >= 0.5 {
		return "balanced"
	}
	return "needs-attention"
}
func (q QualityScore) Clamp() QualityScore {
	values := []*float64{&q.Balance, &q.Diversity, &q.Stability, &q.Total}
	for _, value := range values {
		if *value < 0 {
			*value = 0
		}
		if *value > 1 {
			*value = 1
		}
	}
	return q
}
