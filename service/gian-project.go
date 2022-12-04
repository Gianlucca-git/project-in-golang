package service

import (
	"IMPORTS/model/dto"
	"IMPORTS/repository"
	"errors"
	"fmt"
	"log"
	"strings"
)

// ServiceManager implement methods
type ServiceManager interface {
	// Constants Messages
	Constants(key string) string
	// OrderList responsible for the logic to sort the list
	OrderList(request *dto.ClassifiedList) error
	// ValidatedRequestBalance responsible for validating the request from the Balance endpoint
	ValidatedRequestBalance(request *dto.BalanceRequest) error
	// GeneralBalance responsible for making the balance of all months in the application
	GeneralBalance(request *dto.BalanceRequest)
}

// NewServiceManager Constructs a new ServiceManager
func NewServiceManager(t repository.Type) ServiceManager {
	return &ServiceStruct{
		ReplaceManager: repository.NewReplaceManager(t),
	}
}

type ServiceStruct struct {
	repository.ReplaceManager
}

// OrderList responsible for the logic to sort the list
func (sm *ServiceStruct) OrderList(request *dto.ClassifiedList) error {
	const MaxLengthAllow = 100

	if len(request.Unclassified) > MaxLengthAllow {
		return errors.New(sm.Constants("InvalidLengthList"))
	}

	request.Classified = request.Unclassified
	if len(request.Unclassified) == 0 {
		return nil
	}

	var unclassifiedCopy []int
	unclassifiedCopy = append(unclassifiedCopy, request.Classified...)

	var duplicated []int
	oneTimes := true
	lenList := len(request.Classified) - 1
	i := 0
	indexCurrentMinor := 0
	var noDuplicated []int
	for i <= lenList {

		// sort the values without repeating them
		j := i + 1
		for j <= lenList {
			if request.Classified[j] < request.Classified[indexCurrentMinor] {
				indexCurrentMinor = j
			}
			j++
		}

		if oneTimes {
			noDuplicated = append(noDuplicated, request.Classified[indexCurrentMinor])
			oneTimes = false
		} else {
			if noDuplicated[len(noDuplicated)-1] != request.Classified[indexCurrentMinor] {
				noDuplicated = append(noDuplicated, request.Classified[indexCurrentMinor])
			}
		}

		// get the repeated values according to their occurrence
		k := i - 1
		for k >= 0 {
			if unclassifiedCopy[k] == unclassifiedCopy[i] {
				duplicated = append(duplicated, unclassifiedCopy[i])
				break
			}
			k--
		}

		request.Classified[indexCurrentMinor] = request.Classified[i]
		i++
		indexCurrentMinor = i
	}

	request.Classified = append(noDuplicated, duplicated...)
	return nil
}

// ValidatedRequestBalance responsible for validating the request from the Balance endpoint
func (sm *ServiceStruct) ValidatedRequestBalance(request *dto.BalanceRequest) error {
	log.Print("[INFO] init: ValidatedRequestBalance()")

	const MaxLengthAllow = 100
	if len(request.Months) > MaxLengthAllow || len(request.Sales) > MaxLengthAllow || len(request.Bills) > MaxLengthAllow {
		return errors.New(sm.Constants("InvalidLengthList"))
	}

	// check that all lists have the same length
	if !(len(request.Months) == len(request.Sales) && len(request.Sales) == len(request.Bills)) {
		return errors.New(sm.Constants("InvalidLengths"))
	}

	// definition of the months that are enabled for the balance
	var months = map[string]bool{"enero": true, "febrero": true, "marzo": true, "abril": true, "mayo": true, "junio": true, "julio": true, "agosto": true, "septiembre": true, "octubre": true, "noviembre": true, "diciembre": true}
	for i, value := range request.Months {

		if !months[strings.ToLower(value)] {
			return errors.New(fmt.Sprintf("%s (value = %s)", sm.Constants("InvalidMonth"), value))
		}
		if request.Sales[i] < 0 {
			return errors.New(fmt.Sprintf("%s (value = %d)", sm.Constants("InvalidNumber"), request.Sales[i]))
		}
		if request.Bills[i] < 0 {
			return errors.New(fmt.Sprintf("%s (value = %d)", sm.Constants("InvalidNumber"), request.Bills[i]))
		}
	}

	return nil
}

// GeneralBalance responsible for making the balance of all months in the application
func (sm *ServiceStruct) GeneralBalance(request *dto.BalanceRequest) {

}
