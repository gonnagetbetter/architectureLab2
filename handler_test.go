package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("1 2 +"),
		Output: b,
	}

	if assert.Nil(t, handler.Compute()) {
		assert.Equal(t, "+ 1 2", b.String())
	}
}

func TestComputeHandlerError(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("01c www 99d d"),
		Output: b,
	}

	if assert.Error(t, handler.Compute()) {
		assert.Equal(t, "", b.String())
	}
}
