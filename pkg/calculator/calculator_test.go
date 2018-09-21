package calculator

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAdd(t *testing.T) {
	service := NewCalculatorSoap("", false, nil)
	resp, err := service.Add(&Add{IntA: 10, IntB: 5})
	if err != nil {
		t.Error("Could not request", err)
	} else {
		assert.Equal(t, resp.AddResult, int32(15))
	}
}

func TestSubtract(t *testing.T) {
	service := NewCalculatorSoap("", false, nil)

	resp, err := service.Subtract(&Subtract{IntA: 10, IntB: 5})
	if err != nil {
		t.Error("Could not request", err)
	} else {
		assert.Equal(t, resp.SubtractResult, int32(5))
	}
}

func TestDevine(t *testing.T) {
	service := NewCalculatorSoap("", false, nil)
	resp, err := service.Divide(&Divide{IntA: 10, IntB: 5})
	if err != nil {
		t.Error("Could not request", err)
	} else {
		assert.Equal(t, resp.DivideResult, int32(2))
	}
}

func TestMultiply(t *testing.T) {
	service := NewCalculatorSoap("", false, nil)
	resp, err := service.Multiply(&Multiply{IntA: 10, IntB: 5})
	if err != nil {
		t.Error("Could not request", err)
	} else {
		assert.Equal(t, resp.MultiplyResult, int32(50))
	}
}
