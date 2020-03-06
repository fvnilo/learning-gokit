package main

import (
	"errors"
	"strings"
)

var ErrEmpty = errors.New("Empty string")

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

type ServiceMiddleware func(StringService) StringService

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
