package aoc2024

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed data/example-day-2.txt
	dayTwoExample []byte

	//go:embed data/input-day-2.txt
	dayTwoInput []byte
)

func TestDayTwo(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got, err := dayTwo(dayTwoExample)
		assert.NoError(t, err)
		assert.Equal(t, 2, got)
	})

	t.Run("input", func(t *testing.T) {
		got, err := dayTwo(dayTwoInput)
		assert.NoError(t, err)
		t.Log(got)
	})
}
