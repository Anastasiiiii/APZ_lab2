package lab2

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix1(t *testing.T) {
	res, err := PrefixToPostfix("+ 1 * - 5 4 2")
	assert.Equal(t, res, "1 5 4 - 2 * +")
	assert.Nil(t, err)
}

func TestPrefixToPostfix2(t *testing.T) {
	res, err := PrefixToPostfix("/ * 6 + 1 2 2")
	assert.Equal(t, res, "6 1 2 + * 2 /")
	assert.Nil(t, err)
}

func TestPrefixToPostfix3(t *testing.T) {
	res, err := PrefixToPostfix("+ 36 * - 2 11 1")
	assert.Equal(t, res, "36 2 11 - 1 * +")
	assert.Nil(t, err)
}

func TestPrefixToPostfix4(t *testing.T) {
	res, err := PrefixToPostfix("- + - - 10 6 * 7 ^ 2 13 55 / 1 3")
	assert.Equal(t, res, "10 6 - 7 2 13 ^ * - 55 + 1 3 / -")
	assert.Nil(t, err)
}

func TestEmptyPrefixError(t *testing.T) {
	_, err := PrefixToPostfix("")
	assert.Error(t, err)
}

func TestNoOperandsError(t *testing.T) {
	_, err := PrefixToPostfix("1 2 3 4 5")
	assert.Nil(t, err)
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 2 10 * 4")
	fmt.Println(res)
}
