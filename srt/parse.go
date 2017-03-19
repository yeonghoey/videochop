package srt

import (
	"fmt"
	"regexp"
	"time"
)

var durationRE = regexp.MustCompile(`^(\d\d):(\d\d):(\d\d),(\d\d\d)$`)

// ParseDuration parses a duration string.
// A duration string should be formatted as "hh:mm:ss,SSS",
// which is specified in SubRip specification.
func ParseDuration(s string) (time.Duration, error) {
	ns := durationRE.FindStringSubmatch(s)
	if ns == nil {
		return 0, fmt.Errorf("invalid duration string: %s", s)
	}

	ss := fmt.Sprintf("%sh%sm%ss%sms", ns[1], ns[2], ns[3], ns[4])
	d, err := time.ParseDuration(ss)
	if err != nil {
		return 0, err
	}

	return d, nil
}
