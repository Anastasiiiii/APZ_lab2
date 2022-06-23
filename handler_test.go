package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix(test *testing.T) {
	expression := "+ 2 * 3 6"
	expected := "2 3 6 + *"
	input := strings.NewReader(expression)
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if assert.Nil(test, err) {
		assert.Equal(test, expected, output.String())
	}
}

func TestEmptyExpression(test *testing.T) {
	expression := ""
	input := strings.NewReader(expression)
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	assert.NotNil(test, err)
}

func TestWrongExpression(test *testing.T) {
	expression := ""
	input := strings.NewReader(expression)
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	assert.NotNil(test, err)
}
