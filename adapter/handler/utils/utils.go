package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewUUID(short bool) string {
	s := uuid.NewString()
	if short {
		return strings.Split(s, "-")[0]
	}
	return s
}

func NewQuit() chan struct{} {
	return make(chan struct{}, 1)
}
