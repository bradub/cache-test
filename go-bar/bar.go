// Package bar provides small helpers used to exercise CI caching.
package bar

import (
	"os"

	"github.com/aarondl/opt/null"
	"github.com/rs/zerolog"
)

// Sum returns the sum of the provided integers.
func Sum(xs []int) int {
	total := 0

	for _, x := range xs {
		total += x
	}

	return total
}

// Describe renders a human-readable description for the optional name.
func Describe(name null.Val[string]) string {
	value, ok := name.Get()
	if !ok {
		return "anonymous"
	}

	return "name=" + value
}

// Logger returns a zerolog logger writing to stderr.
func Logger() zerolog.Logger {
	return zerolog.New(os.Stderr).With().Timestamp().Logger()
}
