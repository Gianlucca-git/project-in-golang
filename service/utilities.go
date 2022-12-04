package service

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
