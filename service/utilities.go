package service

import (
	"regexp"
)

const (
	yyyy_mm_dd = "2006-01-02"
)

// Constants retorna mensajes compartidos
func (sm *ServiceStruct) Constants(key string) string {
	constants := map[string]string{
		"InvalidLengthList": "list length exceeds allowed length",
		"InvalidMonth":      "a value entered in months is invalid",
		"InvalidNumber":     "an integer entered in the request is less than zero",
		"InvalidLengths":    "the lengths of the lists are not equal",
	}
	return constants[key]
}

func (sm *ServiceStruct) RegularExpression(str string, typeExpression string) bool {

	switch typeExpression {
	case "alphabetic":
		expression, _ := regexp.Compile(`^[a-z|A-Z]+$`)
		return expression.MatchString(str)
	}

	return false
}
