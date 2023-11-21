package main

import (
	"errors"
	"fmt"
)

type Temperature struct {
	Value        float64
	Unit, Symbol string
}

var description = map[string]Temperature{"Kelvin": {Value: 0, Unit: "Kelvin", Symbol: "K"},
	"Celsius":   {-273.15, "Celsius", "C"},
	"Farenheit": {Symbol: "F"},
	"Rakine":    {Symbol: "Ra"}}

// methods
func (t Temperature) convertToKelvin() (Temperature, error) {
	var k float64
	switch t.Unit {
	case "Celsius":
		k = t.Value + 273.15
	case "Farenheit":
		k = (t.Value-32)*5/9 + 273.5
	case "Rakine", "rankine":
		k = t.Value * 5 / 9
	default:
		k = t.Value
	}
	if k < 0 {
		return Temperature{}, errors.New("temperature can't go below absolute zero")
	}
	return Temperature{k, "Kelvin", "K"}, nil
}
func (t Temperature) convertToCelsius() (float64, error) {
	var c float64
	switch t.Unit {
	case "Kelvin":
		c = t.Value - 273.15
	case "Farenheit":
		c = (t.Value - 32) * 5 / 9
	case "Rakine":
		c = (t.Value - 491.67) * 5 / 9
	default:
		c = t.Value
	}
	if c < -273.15 {
		return 0, errors.New("temperature can't go below absolute zero")
	}
	return c, nil
}
func (t Temperature) convertToFarenheit() (float64, error) {
	var f float64
	switch t.Unit {
	case "Celsius":
		f = t.Value*9/5 + 32
	default:
		f = t.Value
	}
	return f, nil

}
func (t Temperature) convertToRakine() (float64, error) {
	var r float64
	switch t.Unit {
	case "Celsius":
		r = t.Value*9/5 + 491.67
	default:
		r = t.Value
	}
	return r, nil

}

func main() {
	var boiling = Temperature{Value: 100, Unit: "Rakine", Symbol: "C"}
	fmt.Println(boiling)
	fmt.Println(boiling.convertToKelvin())
	fmt.Println(description)
}
