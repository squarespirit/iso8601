package iso8601

import (
	"errors"
	"fmt"
)

var (
	// ErrorZoneCharacters indicates an incorrect amount of characters was passed to ParseISOZone.
	ErrZoneCharacters = errors.New("iso8601: Expected at least 5 characters for zone information")

	// ErrRemainingData indicates that there is extra data after a `Z` character.
	ErrRemainingData = errors.New("iso8601: Unexepected remaining data after `Z`")

	// ErrNotString indicates that a non string type was passed to the UnmarshalJSON method of `Time`.
	ErrNotString = errors.New("iso8601: Invalid json type (expected string)")
)

func newUnexpectedCharacterError(c byte) error {
	return &UnexpectedCharacterError{Character: c}
}

// UnexpectedCharacterError indicates the parser scanned a character that was not expected at that time.
type UnexpectedCharacterError struct {
	Character byte
}

func (e *UnexpectedCharacterError) Error() string {
	return fmt.Sprintf("iso8601: Unexpected character `%c`", e.Character)
}
