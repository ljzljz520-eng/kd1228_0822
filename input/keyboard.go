package input

import "strings"

type KeyMap struct{ bindings map[string]string }

func NewKeyMap() *KeyMap {
	return &KeyMap{bindings: map[string]string{"space": "pause", "enter": "apply", "r": "reset", "1": "select-a", "2": "select-b", "3": "select-c"}}
}
func (m *KeyMap) Bind(key, action string) { m.bindings[strings.ToLower(key)] = action }
func (m *KeyMap) Action(key string) (string, bool) {
	action, ok := m.bindings[strings.ToLower(key)]
	return action, ok
}
func (m *KeyMap) Keys() []string {
	keys := make([]string, 0, len(m.bindings))
	for key := range m.bindings {
		keys = append(keys, key)
	}
	return keys
}
func (m *KeyMap) Resolve(sequence []string) []string {
	actions := []string{}
	for _, key := range sequence {
		if action, ok := m.Action(key); ok {
			actions = append(actions, action)
		}
	}
	return actions
}

type Shortcut struct {
	Keys   []string
	Action string
}

func DefaultShortcuts() []Shortcut {
	return []Shortcut{{Keys: []string{"space"}, Action: "pause"}, {Keys: []string{"shift", "s"}, Action: "save"}, {Keys: []string{"shift", "l"}, Action: "load"}, {Keys: []string{"r"}, Action: "reset"}}
}
func MatchShortcut(keys []string, shortcut Shortcut) bool {
	if len(keys) != len(shortcut.Keys) {
		return false
	}
	for i, key := range keys {
		if strings.ToLower(key) != strings.ToLower(shortcut.Keys[i]) {
			return false
		}
	}
	return true
}

func NormalizeKeys(keys []string) []string {
	result := make([]string, 0, len(keys))
	seen := map[string]bool{}
	for _, key := range keys {
		normalized := strings.ToLower(strings.TrimSpace(key))
		if normalized == "" || seen[normalized] {
			continue
		}
		seen[normalized] = true
		result = append(result, normalized)
	}
	return result
}

func HasModifier(keys []string, modifier string) bool {
	modifier = strings.ToLower(modifier)
	for _, key := range keys {
		if strings.ToLower(key) == modifier {
			return true
		}
	}
	return false
}
