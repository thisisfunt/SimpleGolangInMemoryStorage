package services

import (
	"errors"
	"regexp"
)

var (
	ErrKeyIsTooLong          = errors.New("Key is too long (max 256)")
	ErrKeyContainInvalidChar = errors.New("Key contain invalid char")
	ErrKeyIsEmpty            = errors.New("Key is empty")
)

func IsKeyValid(key string) error {
	if key == "" {
		return ErrKeyIsEmpty
	}

	if !regexp.MustCompile("^[a-zA-Z0-9_]+$").MatchString(key) {
		return ErrKeyContainInvalidChar
	}

	if len(key) > 256 {
		return ErrKeyIsTooLong
	}

	return nil
}
