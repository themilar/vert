package vert_test

import (
	"testing"

	"github.com/themilar/vert"
)

func TestToKelvin(t *testing.T) {
	v := vert.Temperature{100, "Celsius", "C"}
	exp := 373.15
	res, _ := v.ToKelvin()
	if res.Value != exp {
		// t.Errorf("Expected %d, got %d instead.\n", exp, res)
		t.Error("wrong")
	}
}
