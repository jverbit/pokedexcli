package main

import (
	"strings"
)

func cleanInput(text string) []string {
	slice := strings.Fields(strings.ToLower(text))
	if slice != nil {
		return slice
	}
	return []string{}
}
