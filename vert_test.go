package vert_test

import (
	"fmt"
	"testing"

	"github.com/themilar/vert"
)

var boilingArray = [4]vert.Temperature{
	{373.15, "Kelvin", "K"},
	{100, "Celsius", "C"},
	{212, "Farenheit", "F"},
	{671.67, "Rankine", "Ra"},
}

func TestToKelvin(t *testing.T) {

	for i := 0; i < len(boilingArray); i++ {
		exp := boilingArray[0]
		res, _ := boilingArray[i].ToKelvin()

		if res != exp {
			t.Error("wrong")
		}
	}
	// v := vert.Temperature{100, "Celsius", "C"}
	// exp := 373.15
	// res, _ := v.ToKelvin()
	// if res.Value != exp {
	// 	// t.Errorf("Expected %d, got %d instead.\n", exp, res)
	// 	t.Error("wrong")
	// }
}

func TestToCelsius(t *testing.T) {

	for i := 0; i < len(boilingArray); i++ {
		exp := boilingArray[1]
		res, _ := boilingArray[i].ToCelsius()
		fmt.Print(res)
		if res != exp {
			t.Error("wrong")
		}
	}
}
func TestToFarenheit(t *testing.T) {

	for i := 0; i < len(boilingArray); i++ {
		exp := boilingArray[2]
		res, _ := boilingArray[i].ToFarenheit()
		fmt.Print(res)
		if res != exp {
			t.Error("wrong")
		}
	}
}
func TestToRankine(t *testing.T) {

	for i := 0; i < len(boilingArray); i++ {
		exp := boilingArray[3]
		res, _ := boilingArray[i].ToRankine()
		fmt.Print(res)
		if res != exp {
			t.Error("wrong")
		}
	}
}
