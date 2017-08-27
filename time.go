package shopgun

import "time"

// Time is a wrapper to parse unmarshal json time
type Time struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05+0000"

// UnmarshalJSON unmarshals a time oject
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	t.Time, err = time.Parse(ctLayout, string(b))
	return
}

// MarshalJSON marshals a time oject
func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(ctLayout)), nil
}
