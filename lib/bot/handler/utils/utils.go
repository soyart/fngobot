package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewUUID() string {
	// @TODO: Proper UUID
	return strings.Split(
		uuid.NewString(), "-",
	)[0]
}

func NewQuit() chan struct{} {
	return make(chan struct{}, 1)
}
