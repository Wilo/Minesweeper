package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumeric(t *testing.T) {
	t.Run("When the result is a number", func(t *testing.T) {
		assert.Equal(t, isNumeric("12"), true)
	})
	t.Run("When the result is not a number", func(t *testing.T) {
		assert.Equal(t, isNumeric("ABC"), false)
	})
	t.Run("When the input is alphanumeric", func(t *testing.T) {
		assert.Equal(t, isNumeric("A12"), false)
	})
}

func TestGenerateRows(t *testing.T) {
	t.Run("When the data is correct", func(t *testing.T) {
		rows, err := GenerateRows(3)
		if assert.NoError(t, err) {
			assert.Equal(t, rows, []string{"A", "B", "C"})
		}
	})
	t.Run("When the data is incorrect", func(t *testing.T) {
		rows, err := GenerateRows(30)
		if assert.NotNil(t, err) {
			assert.Nil(t, rows)
		}
	})
}
