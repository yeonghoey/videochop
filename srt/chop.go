package srt

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

func Chop(r io.Reader, w io.Writer, start time.Duration, end time.Duration) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	fmt.Println(buf.String())
}

func parseTime(s string) (time.Duration, error) {
	return time.Duration(0), nil
}
