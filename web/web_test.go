package web

import (
	"strings"
)

func contains(value string, values []string) bool {
	for i := range values {
		if values[i] == value {
			return true
		}
	}

	return false
}

func match(value string, values []string) bool {
	for i := range values {
		if strings.Contains(values[i], value) {
			return true
		}
	}

	return false
}

func CreateService() *HttpService {
	return newStubService("localhost", 9999, "../static")
}
