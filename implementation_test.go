package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix1(t *testing.T) {
	res, err := PostfixToPrefix("1 2 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "+ 1 2", res)
	}
}

func TestPostfixToPrefix2(t *testing.T) {
	res, err := PostfixToPrefix("1 2 + 3 ^")
	if assert.Nil(t, err) {
		assert.Equal(t, "^ + 1 2 3", res)
	}
}

func TestPostfixToPrefix3(t *testing.T) {
	res, err := PostfixToPrefix("1 2 3 + * 4 ^")
	if assert.Nil(t, err) {
		assert.Equal(t, "^ * 1 + 2 3 4", res)
	}
}

func TestPostfixToPrefix4(t *testing.T) {
	res, err := PostfixToPrefix("1 2 - 3 4 + /")
	if assert.Nil(t, err) {
		assert.Equal(t, "/ - 1 2 + 3 4", res)
	}
}

func TestPostfixToPrefix5(t *testing.T) {
	res, err := PostfixToPrefix("1 2 3 + * 4 - 5 / 6 7 + *")
	if assert.Nil(t, err) {
		assert.Equal(t, "* / - * 1 + 2 3 4 5 + 6 7", res)
	}
}

func TestPostfixToPrefix6(t *testing.T) {
	res, err := PostfixToPrefix("1 2 * 3 4 - / 5 6 + + 7 8 9 - - 10 - *")
	if assert.Nil(t, err) {
		assert.Equal(t, "* + / * 1 2 - 3 4 + 5 6 - - 7 - 8 9 10", res)
	}
}
