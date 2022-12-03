package service

// Utilities is utilities for development
type Utilities interface {
}

// NewUtil constructs a new Util
func NewUtil() Utilities {
	return &utilities{}
}

type utilities struct{}
