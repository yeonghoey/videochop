package srt

import (
	"testing"
	"time"
)

func TestParseDurtaion(t *testing.T) {
	const (
		h  = time.Hour
		m  = time.Minute
		s  = time.Second
		ms = time.Millisecond
	)

	var tests = []struct {
		input string
		want  time.Duration
		ok    bool
	}{
		{"", 0, false},
		{"0", 0, false},
		{"hello world", 0, false},
		{"00:00:00", 0, false},
		{"00:00:00,000", 0, true},
		{"00:00:00,001", 1 * ms, true},
		{"00:00:01,001", 1*s + 1*ms, true},
		{"00:01:01,001", 1*m + 1*s + 1*ms, true},
		{"01:01:01,001", 1*h + 1*m + 1*s + 1*ms, true},
	}

	for _, test := range tests {
		d, err := ParseDuration(test.input)
		ok := (err == nil)
		if ok != test.ok || d != test.want {
			t.Errorf("parseTime(%q) = (%v, %v)", test.input, d, err)
		}
	}
}
