package vert_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/themilar/vert"
)

var boilingArray = [4]vert.Temperature{
	{373.15, "Kelvin", "K"},
	{100, "Celsius", "C"},
	{212, "Fahrenheit", "F"},
	{671.67, "Rankine", "Ra"},
}

func roundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
func TestToKelvin(t *testing.T) {

	for i := 0; i < len(boilingArray); i++ {
		exp := boilingArray[0]
		res, _ := boilingArray[i].ToKelvin()

		if res != exp {
			t.Errorf("wrong expected %.18f, got %.18f", exp.Value, roundFloat(res.Value, 2))
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

		if res != exp {
			t.Errorf("wrong expected %.18f, got %.18f", roundFloat(exp.Value, 2), roundFloat(res.Value, 2))
		}
	}
}
func TestToFahrenheit(t *testing.T) {

	for i := 0; i < len(boilingArray); i++ {
		exp := boilingArray[2]
		res, _ := boilingArray[i].ToFahrenheit()
		fmt.Print(res)
		res.Value = roundFloat(res.Value, 2)
		if res != exp {
			t.Errorf("wrong expected %.20f, got %.20f", exp.Value, res.Value)

		}
	}
}
func TestToRankine(t *testing.T) {

	for i := 0; i < len(boilingArray); i++ {
		exp := boilingArray[3]
		res, _ := boilingArray[i].ToRankine()
		fmt.Print(res)
		if res != exp {
			t.Errorf("wrong expected %f, got %f", exp.Value, res.Value)

		}
	}
}
