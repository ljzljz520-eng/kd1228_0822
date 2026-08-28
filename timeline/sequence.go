package timeline

import "sort"

type Sequence struct{ keys []Keyframe }

func NewSequence() *Sequence { return &Sequence{keys: []Keyframe{}} }
func (s *Sequence) Add(keyframe Keyframe) {
	s.keys = append(s.keys, keyframe)
	sort.Slice(s.keys, func(i, j int) bool { return s.keys[i].Frame < s.keys[j].Frame })
}
func (s *Sequence) Keys() []Keyframe { return append([]Keyframe(nil), s.keys...) }
func (s *Sequence) At(frame uint64) (Keyframe, bool) {
	if len(s.keys) == 0 {
		return Keyframe{}, false
	}
	if frame <= s.keys[0].Frame {
		return s.keys[0], true
	}
	for i := 1; i < len(s.keys); i++ {
		if frame <= s.keys[i].Frame {
			return InterpolateKeyframes(s.keys[i-1], s.keys[i], frame), true
		}
	}
	return s.keys[len(s.keys)-1], true
}
func (s *Sequence) Duration() uint64 {
	if len(s.keys) == 0 {
		return 0
	}
	return s.keys[len(s.keys)-1].Frame
}
func (s *Sequence) Clear()        { s.keys = nil }
func (s *Sequence) IsEmpty() bool { return len(s.keys) == 0 }

func (s *Sequence) Contains(frame uint64) bool {
	for _, key := range s.keys {
		if key.Frame == frame {
			return true
		}
	}
	return false
}

func (s *Sequence) NextAfter(frame uint64) (Keyframe, bool) {
	for _, key := range s.keys {
		if key.Frame > frame {
			return key, true
		}
	}
	return Keyframe{}, false
}

func (s *Sequence) PreviousBefore(frame uint64) (Keyframe, bool) {
	var previous Keyframe
	found := false
	for _, key := range s.keys {
		if key.Frame >= frame {
			break
		}
		previous = key
		found = true
	}
	return previous, found
}
func (s *Sequence) Reverse() *Sequence {
	result := NewSequence()
	for i := len(s.keys) - 1; i >= 0; i-- {
		result.Add(s.keys[i])
	}
	return result
}
