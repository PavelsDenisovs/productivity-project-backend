package utils

import (
	"math/rand"
	"fmt"
	"time"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateVerificationCode() string {
	return fmt.Sprintf("%06d", randomGenerator.Intn(1000000))
}