package service

// Utilities is utilities for development
type Utilities interface {
}

// NewUtil constructs a new Util
func NewUtil() Utilities {
	return &utilities{}
}

type utilities struct{}

const (
	InvalidLengthList = "list length exceeds allowed length"
	InvalidMonth      = "a value entered in months is invalid"
	InvalidNumber     = "an integer entered in the request is less than zero"
	invalidLengths    = "the lengths of the lists are not equal"
)
