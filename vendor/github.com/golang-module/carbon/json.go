package carbon

import (
	"bytes"
	"fmt"
	"strconv"
)

// DateTime defines a DateTime struct.
type DateTime struct {
	Carbon
}

// Date defines a Date struct.
type Date struct {
	Carbon
}

// Time defines a Time struct.
type Time struct {
	Carbon
}

// Timestamp defines a Timestamp struct.
type Timestamp struct {
	Carbon
}

// TimestampWithSecond defines a TimestampWithSecond struct.
type TimestampWithSecond struct {
	Carbon
}

// TimestampWithMillisecond defines a TimestampWithMillisecond struct.
type TimestampWithMillisecond struct {
	Carbon
}

// TimestampWithMicrosecond defines a TimestampWithMicrosecond struct.
type TimestampWithMicrosecond struct {
	Carbon
}

// TimestampWithNanosecond defines a TimestampWithNanosecond struct.
type TimestampWithNanosecond struct {
	Carbon
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *DateTime) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = DateTime{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *Date) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = Date{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *Time) UnmarshalJSON(b []byte) error {
	c := ParseByFormat(string(bytes.Trim(b, `"`)), "H:i:s")
	if c.Error == nil {
		*t = Time{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.Timestamp())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = Timestamp{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampWithSecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampWithSecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithSecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = TimestampWithSecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampWithMillisecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampWithMillisecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithMillisecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = TimestampWithMillisecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampWithMicrosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampWithMicrosecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithMicrosecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = TimestampWithMicrosecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t TimestampWithNanosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampWithNanosecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *TimestampWithNanosecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = TimestampWithNanosecond{c}
	}
	return nil
}

/* The following will be removed from v2.0 and only the above will be retained */

// ToDateTimeString defines a ToDateTimeString struct.
type ToDateTimeString struct {
	Carbon
}

// ToDateString defines a ToDateString struct.
type ToDateString struct {
	Carbon
}

// ToTimeString defines a ToTimeString struct.
type ToTimeString struct {
	Carbon
}

// ToTimestamp defines a ToTimestamp struct.
type ToTimestamp struct {
	Carbon
}

// ToTimestampWithSecond defines a ToTimestampWithSecond struct.
type ToTimestampWithSecond struct {
	Carbon
}

// ToTimestampWithMillisecond defines a ToTimestampWithMillisecond struct.
type ToTimestampWithMillisecond struct {
	Carbon
}

// ToTimestampWithMicrosecond defines a ToTimestampWithMicrosecond struct.
type ToTimestampWithMicrosecond struct {
	Carbon
}

// ToTimestampWithNanosecond defines a ToTimestampWithNanosecond struct.
type ToTimestampWithNanosecond struct {
	Carbon
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToDateTimeString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToDateTimeString) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = ToDateTimeString{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToDateString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToDateString) UnmarshalJSON(b []byte) error {
	c := Parse(string(bytes.Trim(b, `"`)))
	if c.Error == nil {
		*t = ToDateString{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimeString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeString())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimeString) UnmarshalJSON(b []byte) error {
	c := ParseByFormat(string(bytes.Trim(b, `"`)), "H:i:s")
	if c.Error == nil {
		*t = ToTimeString{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestamp())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = ToTimestamp{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimestampWithSecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithSecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimestampWithSecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = ToTimestampWithSecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimestampWithMillisecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithMillisecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimestampWithMillisecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = ToTimestampWithMillisecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimestampWithMicrosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithMicrosecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimestampWithMicrosecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = ToTimestampWithMicrosecond{c}
	}
	return nil
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (t ToTimestampWithNanosecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.ToTimestampWithNanosecond())), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *ToTimestampWithNanosecond) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if ts == 0 || err != nil {
		return err
	}
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = ToTimestampWithNanosecond{c}
	}
	return nil
}
