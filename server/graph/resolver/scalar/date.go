package scalar

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// Date is a custom GraphQL scalar type to represent time.Time data.
// It wraps a string to provide time.Time data and reverse functionality.
type Date struct {
	*time.Time
}

// MarshalGQL implements the graphql.Marshaler interface found in gqlgen,
// allowing the type to be marshaled by gqlgen and sent over the wire.
// This will convert the string as a time.Time string.
func (d Date) MarshalGQL(w io.Writer) {
	if d.Time == nil || d.Time.IsZero() {
		io.WriteString(w, "null")
		return
	}
	io.WriteString(w, strconv.Quote(d.Format(time.RFC3339)))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface found in gqlgen,
// allowing the type to be received by a graphql client and unmarshaled.
// The input is expected to be a base64-encoded string, which will be decoded
// into the byte slice.
func (d *Date) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("Date must be a RFC3339 formatted string")
	}

	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return fmt.Errorf("Date could not be parsed: %w", err)
	}

	d.Time = &t
	return nil
}
