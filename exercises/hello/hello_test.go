package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloEnglish(t *testing.T) {
	actual := Hello("John", "en")
	expected := "Hello, John"

	assert.Equal(t, actual, expected)
}

func TestHelloFrench(t *testing.T) {
	actual := Hello("John", "fr")
	expected := "Bonjour, John"

	assert.Equal(t, actual, expected)
}

func TestHelloSpanish(t *testing.T) {
	actual := Hello("John", "es")
	expected := "Hola, John"

	assert.Equal(t, actual, expected)
}

func TestHelloDefault(t *testing.T) {
	actual := Hello("", "")
	expected := "Hello, world"

	assert.Equal(t, actual, expected)
}
