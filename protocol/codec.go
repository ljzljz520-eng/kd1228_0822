package protocol

import (
	"encoding/json"
	"fmt"
	"jellyfield/model"
)

type Envelope struct {
	Type    string             `json:"type"`
	Scene   string             `json:"scene"`
	Command *Command           `json:"command,omitempty"`
	Frame   *model.RenderFrame `json:"frame,omitempty"`
	Error   string             `json:"error,omitempty"`
}

func Encode(value Envelope) (string, error) {
	data, err := json.Marshal(value)
	return string(data), err
}
func Decode(value string) (Envelope, error) {
	var envelope Envelope
	if err := json.Unmarshal([]byte(value), &envelope); err != nil {
		return envelope, err
	}
	if envelope.Type == "" {
		return envelope, fmt.Errorf("message type required")
	}
	return envelope, nil
}
func Ack(scene string, frame model.RenderFrame) Envelope {
	return Envelope{Type: "frame", Scene: scene, Frame: &frame}
}
func Failure(scene string, err error) Envelope {
	message := "unknown error"
	if err != nil {
		message = err.Error()
	}
	return Envelope{Type: "error", Scene: scene, Error: message}
}
func CommandEnvelope(scene string, command Command) Envelope {
	return Envelope{Type: "command", Scene: scene, Command: &command}
}
