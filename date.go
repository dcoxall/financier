package financier

import "time"

// A wrapper for time.Time which will decode and encode JSON using YYYY-MM-DD
type Date time.Time

// UnmarshalJSON implements the json.Unmarshaler interface.
// The date is expected to be a quoted string in YYYY-MM-DD format.
func (d *Date) UnmarshalJSON(data []byte) (err error) {
	t, err := time.Parse(`"2006-01-02"`, string(data))
	date := Date(t)
	*d = date
	return
}

// MarshalJSON implements the json.Marshaler interface.
// The date is a quoted string in YYYY-MM-DD format, with sub-second precision
// added if present.
func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(t.Format(`"2006-01-02"`)), nil
}
