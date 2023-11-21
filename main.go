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
func (t Temperature) convertToKelvin() (float64, error) {
	var k Temperature
	switch t.Unit {
	case "Celsius":
		k.Value = t.Value + 273.15
	case "Farenheit":
		k.Value = (t.Value-32)*5/9 + 273.5
	case "Rakine", "rankine":
		k.Value = t.Value * 5 / 9
	default:
		k.Value = t.Value
	}
	if k.Value < 0 {
		return 0, errors.New("temperature can't go below absolute zero")
	}
	return k.Value, nil
}
func (t Temperature) convertToCelsius() (float64, error) {
	var c float64
	switch t.Unit {
	case "Kelvin":
		c = t.Value - 273.15
	default:
		c = t.Value
	}
	if c < -273.15 {
		return 0, errors.New("temperature can't go below absolute zero")
	}
	return c, nil
}
func (t Temperature) convertToFarenheit() (float64, error) {}
func (t Temperature) convertToRakine() (float64, error) {

}

func main() {
	var boiling = Temperature{Value: 100, Unit: "Rakine", Symbol: "C"}
	fmt.Println(boiling)
	fmt.Println(boiling.convertToKelvin())
	fmt.Println(description)
}
