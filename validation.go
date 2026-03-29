package main

import (
	"strconv"
	"strings"
	"time"
)

func isValidName(s string) bool {
	trimmed := strings.TrimSpace(s)
	return len(trimmed) > 0
}

func isValidAge(s string) bool {
	age, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return false
	}
	return age > 0 && age < 150
}

func isValidAddictYears(s string) bool {
	years, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return false
	}
	return years >= 0 && years < 100
}

func isValidYear(s string) bool {
	yr, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return false
	}
	now := time.Now()
	return yr >= 1900 && yr <= now.Year()
}

func isValidMonth(s string) bool {
	mo, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return false
	}
	return mo >= 1 && mo <= 12
}

func isValidDay(s string) bool {
	dy, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return false
	}
	return dy >= 1 && dy <= 31
}

func isValidHour(s string) bool {
	hr, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return false
	}
	return hr >= 0 && hr <= 23
}

func isValidMinute(s string) bool {
	mn, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return false
	}
	return mn >= 0 && mn <= 59
}
