package srt

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestChop(t *testing.T) {
	r := strings.NewReader("This is a data")
	w := new(bytes.Buffer)
	s := time.Second
	e := 2 * time.Second
	Chop(r, w, s, e)
}

func TestParseTime(t *testing.T) {
	d, _ := parseTime("00:00:0,000")
	if d != time.Duration(0) {
		t.Errorf("parseTime(00:00:0,000) = %v, want 0s", d)
	}
}
