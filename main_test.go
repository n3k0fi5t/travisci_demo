package main_test

import (
	"testing"

	main "github.com/n3k0fi5t/travisci_demo"
	"github.com/stretchr/testify/assert"
)

func TestWork(t *testing.T) {
	res := main.SomeWork(5)
	assert.Equal(t, res, true)
}
