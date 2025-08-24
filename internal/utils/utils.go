package utils

import (
	"github.com/google/uuid"
	"strings"
)

// GenerateRandomID generates a random 10-digit ID
func GenerateSimplePayID() string {
	u := uuid.New().String()
	u = strings.ReplaceAll(u, "-", "")
	return u[:10]
}
