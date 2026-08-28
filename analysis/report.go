package analysis

import (
	"fmt"
	"jellyfield/model"
	"jellyfield/simulation"
	"sort"
)

type JellyfishStats struct {
	ID            string
	Name          string
	Speed         float64
	Energy        float64
	ParticleCount int
	ColorHex      string
	Rank          int
}
type SceneReport struct {
	SceneID        string
	Frame          uint64
	SelectedID     string
	Total          int
	TotalParticles int
	MeanSpeed      float64
	MeanEnergy     float64
	Brightest      string
	Entries        []JellyfishStats
}

func BuildReport(scene model.SceneState, frame model.RenderFrame, metrics simulation.Metrics) SceneReport {
	report := SceneReport{SceneID: scene.ID, Frame: frame.Frame, SelectedID: frame.SelectedID, Total: len(frame.Jellyfish), TotalParticles: metrics.ParticleCount, MeanSpeed: metrics.AverageSpeed, MeanEnergy: metrics.Brightness, Entries: []JellyfishStats{}}
	for _, j := range frame.Jellyfish {
		energy := meanEnergy(j)
		report.Entries = append(report.Entries, JellyfishStats{ID: j.ID, Name: j.Name, Speed: j.Speed, Energy: energy, ParticleCount: len(j.Particles), ColorHex: j.BaseColor.Hex()})
	}
	sort.Slice(report.Entries, func(i, j int) bool { return report.Entries[i].Energy > report.Entries[j].Energy })
	for i := range report.Entries {
		report.Entries[i].Rank = i + 1
		if i == 0 {
			report.Brightest = report.Entries[i].ID
		}
	}
	return report
}

func meanEnergy(j model.Jellyfish) float64 {
	if len(j.Particles) == 0 {
		return 0
	}
	sum := 0.0
	for _, p := range j.Particles {
		sum += p.Energy
	}
	return sum / float64(len(j.Particles))
}
func (r SceneReport) Selected() (JellyfishStats, bool) {
	for _, entry := range r.Entries {
		if entry.ID == r.SelectedID {
			return entry, true
		}
	}
	return JellyfishStats{}, false
}
func (r SceneReport) Find(id string) (JellyfishStats, bool) {
	for _, entry := range r.Entries {
		if entry.ID == id {
			return entry, true
		}
	}
	return JellyfishStats{}, false
}
func (r SceneReport) SpeedRange() (float64, float64) {
	if len(r.Entries) == 0 {
		return 0, 0
	}
	low, high := r.Entries[0].Speed, r.Entries[0].Speed
	for _, entry := range r.Entries[1:] {
		if entry.Speed < low {
			low = entry.Speed
		}
		if entry.Speed > high {
			high = entry.Speed
		}
	}
	return low, high
}
func (r SceneReport) String() string {
	return fmt.Sprintf("%s frame=%d jellyfish=%d particles=%d brightest=%s", r.SceneID, r.Frame, r.Total, r.TotalParticles, r.Brightest)
}
func (r SceneReport) FilterBySpeed(max float64) []JellyfishStats {
	result := []JellyfishStats{}
	for _, entry := range r.Entries {
		if max <= 0 || entry.Speed <= max {
			result = append(result, entry)
		}
	}
	return result
}
func (r SceneReport) ColorSet() []string {
	seen := map[string]bool{}
	result := []string{}
	for _, entry := range r.Entries {
		if !seen[entry.ColorHex] {
			seen[entry.ColorHex] = true
			result = append(result, entry.ColorHex)
		}
	}
	sort.Strings(result)
	return result
}
