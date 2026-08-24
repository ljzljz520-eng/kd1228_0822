package protocol

import "testing"

func TestCommandRoundTrip(t *testing.T) {
	command, err := Parse("speed jelly-a 2.5")
	if err != nil {
		t.Fatal(err)
	}
	payload, err := Encode(CommandEnvelope("reef-demo", command))
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := Decode(payload)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Type != "command" || decoded.Command.Target != "jelly-a" {
		t.Fatal(decoded)
	}
}
