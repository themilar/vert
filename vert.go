package vert

import (
	"errors"
	"fmt"
)

type Temperature struct {
	Value        float64
	Unit, Symbol string
}

var DescriptionBaseMap = map[string]Temperature{"k": {Value: 0, Unit: "Kelvin", Symbol: "K"},
	"c": {-273.15, "Celsius", "C"},
	"f": {-459.67, "Farenheit", "F"},
	"r": {0, "Rakine", "Ra"}}

// methods
func (t Temperature) ToKelvin() (Temperature, error) {
	var k float64
	switch t.Unit {
	case "Celsius":
		k = t.Value + 273.15
	case "Farenheit":
		k = (t.Value-32)*5/9 + 273.15
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
func (t Temperature) ToCelsius() (Temperature, error) {
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
		return Temperature{}, errors.New("temperature can't go below absolute zero")
	}
	return Temperature{c, "Celsius", "\u00B0C"}, nil
}
func (t Temperature) ToFarenheit() (Temperature, error) {
	var f float64
	switch t.Unit {
	case "Celsius":
		f = t.Value*9/5 + 32
	case "Kelvin":
		f = (t.Value-273.15)*9/5 + 32
	case "Rakine":
		f = t.Value - 459.67
	default:
		f = t.Value
	}
	if f < -459.67 {
		return Temperature{}, errors.New("temperature can't go below absolute zero")
	}
	return Temperature{f, "Farenheit", "\u00B0F"}, nil

}
func (t Temperature) ToRakine() (Temperature, error) {
	var r float64
	switch t.Unit {
	case "Celsius":
		r = t.Value*9/5 + 491.67
	case "Kelvin":
		r = t.Value * 1.8
	case "Farenheit":
		r = t.Value + 459.67
	default:
		r = t.Value
	}
	if r < 0 {
		return Temperature{}, errors.New("temperature can't go below absolute zero")
	}
	return Temperature{r, "Rakine", "\u00B0Ra"}, nil

}

func main() {
	var boiling = Temperature{Value: 373.15, Unit: "Kelvin", Symbol: "Ra"}
	var boiling2 = Temperature{Value: 100, Unit: "Celsius", Symbol: "C"}
	var boiling3 = Temperature{671.67, "Rakine", "Ra"}
	fmt.Println(boiling.ToRakine())
	fmt.Println(boiling2.ToRakine())
	fmt.Println(boiling3.ToCelsius())
	// fmt.Println(description)
}
