package veriff

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	formats := []string{
		time.DateOnly,
	}

	for _, f := range formats {
		parsed, err := time.Parse(f, s)
		if err == nil {
			t.Time = parsed
			return nil
		}
	}

	return fmt.Errorf("cannot parse time string %s", s)
}
