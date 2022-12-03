package service

import (
	"Replace/repository"
)

// ServiceManager implement methods
type ServiceManager interface {
}

// NewServiceManager Constructs a new ServiceManager
func NewServiceManager(t repository.Type) ServiceManager {
	return &serviceStruct{
		ReplaceManager: repository.NewReplaceManager(t),
		Utilities:      NewUtil(),
	}
}

type serviceStruct struct {
	Utilities
	repository.ReplaceManager
}
