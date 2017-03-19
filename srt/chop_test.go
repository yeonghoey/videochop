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
