package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestData(t *testing.T) {
	cadena := "NN100987654321"
	_, err := convertirCadena(cadena)
	assert.Equal(t, err != nil, "fallo")
}
