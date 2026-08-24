package model

import (
	"fmt"
	"math"
)

type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

func NewColor(r, g, b uint8) Color { return Color{R: r, G: g, B: b} }

func (c Color) Hex() string { return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B) }

func (c Color) Luma() float64 {
	return 0.2126*float64(c.R) + 0.7152*float64(c.G) + 0.0722*float64(c.B)
}

func (c Color) Mix(other Color, amount float64) Color {
	if amount < 0 {
		amount = 0
	}
	if amount > 1 {
		amount = 1
	}
	return Color{
		R: uint8(math.Round(float64(c.R) + (float64(other.R)-float64(c.R))*amount)),
		G: uint8(math.Round(float64(c.G) + (float64(other.G)-float64(c.G))*amount)),
		B: uint8(math.Round(float64(c.B) + (float64(other.B)-float64(c.B))*amount)),
	}
}

type Vector struct{ X, Y float64 }

func (v Vector) Add(other Vector) Vector     { return Vector{X: v.X + other.X, Y: v.Y + other.Y} }
func (v Vector) Scale(factor float64) Vector { return Vector{X: v.X * factor, Y: v.Y * factor} }
func (v Vector) Distance(other Vector) float64 {
	dx, dy := v.X-other.X, v.Y-other.Y
	return math.Hypot(dx, dy)
}
func (v Vector) Clamp(min, max Vector) Vector {
	x, y := v.X, v.Y
	if x < min.X {
		x = min.X
	}
	if x > max.X {
		x = max.X
	}
	if y < min.Y {
		y = min.Y
	}
	if y > max.Y {
		y = max.Y
	}
	return Vector{X: x, Y: y}
}

type Particle struct {
	Position Vector  `json:"position"`
	Velocity Vector  `json:"velocity"`
	Energy   float64 `json:"energy"`
	Phase    float64 `json:"phase"`
}

type Jellyfish struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Position  Vector     `json:"position"`
	BaseColor Color      `json:"base_color"`
	Speed     float64    `json:"speed"`
	Radius    float64    `json:"radius"`
	Pulse     float64    `json:"pulse"`
	Particles []Particle `json:"particles"`
}

func NewJellyfish(id, name string, position Vector, color Color, speed float64, count int) Jellyfish {
	if speed <= 0 {
		speed = 1
	}
	if count < 1 {
		count = 1
	}
	particles := make([]Particle, count)
	for i := range particles {
		angle := float64(i) * 0.37
		particles[i] = Particle{Position: position.Add(Vector{X: math.Cos(angle), Y: math.Sin(angle)}), Velocity: Vector{X: math.Sin(angle), Y: math.Cos(angle)}, Energy: 1, Phase: angle}
	}
	return Jellyfish{ID: id, Name: name, Position: position, BaseColor: color, Speed: speed, Radius: 18, Pulse: 0.5, Particles: particles}
}

func (j Jellyfish) Validate() error {
	if j.ID == "" {
		return fmt.Errorf("jellyfish id is required")
	}
	if j.Name == "" {
		return fmt.Errorf("jellyfish name is required")
	}
	if j.Speed < 0.1 || j.Speed > 8 {
		return fmt.Errorf("speed %.2f outside range", j.Speed)
	}
	if j.Radius <= 0 {
		return fmt.Errorf("radius must be positive")
	}
	if len(j.Particles) == 0 {
		return fmt.Errorf("particles are required")
	}
	return nil
}

func (j Jellyfish) ColorAt(intensity float64) Color {
	if intensity < 0 {
		intensity = 0
	}
	if intensity > 1 {
		intensity = 1
	}
	return j.BaseColor.Mix(Color{R: 255, G: 255, B: 255}, intensity*0.35)
}

type SceneState struct {
	ID         string      `json:"id"`
	Width      float64     `json:"width"`
	Height     float64     `json:"height"`
	Frame      uint64      `json:"frame"`
	SelectedID string      `json:"selected_id"`
	Jellyfish  []Jellyfish `json:"jellyfish"`
}

func NewScene(id string, width, height float64) SceneState {
	return SceneState{ID: id, Width: width, Height: height, Jellyfish: []Jellyfish{}}
}
func (s SceneState) Bounds() (Vector, Vector) { return Vector{}, Vector{X: s.Width, Y: s.Height} }
func (s SceneState) Find(id string) (Jellyfish, bool) {
	for _, j := range s.Jellyfish {
		if j.ID == id {
			return j, true
		}
	}
	return Jellyfish{}, false
}
func (s SceneState) Validate() error {
	if s.ID == "" {
		return fmt.Errorf("scene id is required")
	}
	if s.Width <= 0 || s.Height <= 0 {
		return fmt.Errorf("scene dimensions must be positive")
	}
	if len(s.Jellyfish) == 0 {
		return fmt.Errorf("scene needs at least one jellyfish")
	}
	seen := map[string]bool{}
	for _, j := range s.Jellyfish {
		if err := j.Validate(); err != nil {
			return err
		}
		if seen[j.ID] {
			return fmt.Errorf("duplicate jellyfish %s", j.ID)
		}
		seen[j.ID] = true
	}
	return nil
}

type GestureRecord struct {
	ID       string  `json:"id"`
	TargetID string  `json:"target_id"`
	Kind     string  `json:"kind"`
	Amount   float64 `json:"amount"`
	Color    Color   `json:"color"`
	Sequence uint64  `json:"sequence"`
}
type ControlEvent struct {
	ID       string  `json:"id"`
	TargetID string  `json:"target_id"`
	Field    string  `json:"field"`
	Value    float64 `json:"value"`
	Color    Color   `json:"color"`
	Sequence uint64  `json:"sequence"`
}
type RenderFrame struct {
	SceneID    string      `json:"scene_id"`
	Frame      uint64      `json:"frame"`
	SelectedID string      `json:"selected_id"`
	Jellyfish  []Jellyfish `json:"jellyfish"`
}
