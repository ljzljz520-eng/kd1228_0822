package protocol

import (
	"fmt"
	"jellyfield/model"
	"strconv"
	"strings"
)

type Command struct {
	Name   string
	Target string
	Value  string
	Color  model.Color
}

func Parse(line string) (Command, error) {
	parts := strings.Fields(strings.TrimSpace(line))
	if len(parts) == 0 {
		return Command{}, fmt.Errorf("empty command")
	}
	command := Command{Name: strings.ToLower(parts[0])}
	if len(parts) > 1 {
		command.Target = parts[1]
	}
	if len(parts) > 2 {
		command.Value = parts[2]
	}
	if command.Name == "color" && len(parts) > 2 {
		color, err := ParseColor(parts[2])
		if err != nil {
			return Command{}, err
		}
		command.Color = color
	}
	if command.Name != "list" && command.Target == "" {
		return Command{}, fmt.Errorf("target required")
	}
	switch command.Name {
	case "list", "step", "select", "speed", "color", "expand", "contract":
	default:
		return Command{}, fmt.Errorf("unknown command %s", command.Name)
	}
	return command, nil
}
func ParseColor(value string) (model.Color, error) {
	value = strings.TrimPrefix(value, "#")
	if len(value) != 6 {
		return model.Color{}, fmt.Errorf("color must be six hex digits")
	}
	r, err := strconv.ParseUint(value[0:2], 16, 8)
	if err != nil {
		return model.Color{}, err
	}
	g, err := strconv.ParseUint(value[2:4], 16, 8)
	if err != nil {
		return model.Color{}, err
	}
	b, err := strconv.ParseUint(value[4:6], 16, 8)
	if err != nil {
		return model.Color{}, err
	}
	return model.NewColor(uint8(r), uint8(g), uint8(b)), nil
}
func (c Command) Speed() (float64, error) {
	if c.Name != "speed" {
		return 0, fmt.Errorf("not a speed command")
	}
	value, err := strconv.ParseFloat(c.Value, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}
func (c Command) IsMutation() bool {
	switch c.Name {
	case "speed", "color", "expand", "contract":
		return true
	}
	return false
}
