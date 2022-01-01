package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("40000000000000000", "Jose da Silva", 12, 2024, 123)
	assert.Equal(t, expected:"invalid credit card number", err.Error())

}
