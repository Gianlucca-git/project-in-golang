package service

import (
	"IMPORTS/model/dto"
	"IMPORTS/repository"
	"errors"
)

// ServiceManager implement methods
type ServiceManager interface {
	OrderList(unclassified *dto.ClassifiedList) (invalid error)
}

// NewServiceManager Constructs a new ServiceManager
func NewServiceManager(t repository.Type) ServiceManager {
	return &serviceStruct{
		ReplaceManager: repository.NewReplaceManager(t),
		Utilities:      NewUtil(),
	}
}

type serviceStruct struct {
	repository.ReplaceManager
	Utilities
}

func (sm *serviceStruct) OrderList(list *dto.ClassifiedList) (invalid error) {
	const MaxLengthAllow = 100

	if len(list.Unclassified) > MaxLengthAllow {
		return errors.New(InvalidLengthList)
	}

	list.Classified = list.Unclassified
	if len(list.Unclassified) == 0 {
		return nil
	}

	var unclassifiedCopy []int
	unclassifiedCopy = append(unclassifiedCopy, list.Classified...)

	var duplicated []int
	oneTimes := true
	lenList := len(list.Classified) - 1
	i := 0
	indexCurrentMinor := 0
	var noDuplicated []int
	for i <= lenList {

		// sort the values without repeating them
		j := i + 1
		for j <= lenList {
			if list.Classified[j] < list.Classified[indexCurrentMinor] {
				indexCurrentMinor = j
			}
			j++
		}

		if oneTimes {
			noDuplicated = append(noDuplicated, list.Classified[indexCurrentMinor])
			oneTimes = false
		} else {
			if noDuplicated[len(noDuplicated)-1] != list.Classified[indexCurrentMinor] {
				noDuplicated = append(noDuplicated, list.Classified[indexCurrentMinor])
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

		list.Classified[indexCurrentMinor] = list.Classified[i]
		i++
		indexCurrentMinor = i
	}

	//log.Println("LISTA INGRESADA: ", list.Classified)
	//log.Println("LISTA ORDENADA SIN REPETIDOS: ", noDuplicated)
	//log.Println("LISTA DE DUPLICADOS: ", duplicated)
	//log.Println("LISTA RESULTADO: ", append(noDuplicated, duplicated...))

	list.Classified = append(noDuplicated, duplicated...)
	return nil
}
