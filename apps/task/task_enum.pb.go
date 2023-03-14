// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package task

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseTypeFromString Parse Type from string
func ParseTypeFromString(str string) (Type, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := Type_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Type: %s", str)
	}

	return Type(v), nil
}

// Equal type compare
func (t Type) Equal(target Type) bool {
	return t == target
}

// IsIn todo
func (t Type) IsIn(targets ...Type) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t Type) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Type) UnmarshalJSON(b []byte) error {
	ins, err := ParseTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseStageFromString Parse Stage from string
func ParseStageFromString(str string) (Stage, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := Stage_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Stage: %s", str)
	}

	return Stage(v), nil
}

// Equal type compare
func (t Stage) Equal(target Stage) bool {
	return t == target
}

// IsIn todo
func (t Stage) IsIn(targets ...Stage) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t Stage) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Stage) UnmarshalJSON(b []byte) error {
	ins, err := ParseStageFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}