package vert

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
	"Farenheit": {-459.67, "Farenheit", "F"},
	"Rankine":   {0, "Rankine", "Ra"}}

// methods
func (t *Temperature) ToKelvin() error {

	switch t.Unit {
	case "Celsius":
		t.Value = t.Value + 273.15
	case "Farenheit":
		t.Value = (t.Value-32)*5/9 + 273.15
	case "Rankine", "rankine":
		t.Value = t.Value * 5 / 9
	}
	if t.Value < 0 {
		return errors.New("temperature can't go below absolute zero")
	}
	t.Unit = "Kelvin"
	t.Symbol = "K"
	return nil
}
func (t Temperature) ToCelsius() (Temperature, error) {
	var c float64
	switch t.Unit {
	case "Kelvin":
		c = t.Value - 273.15
	case "Farenheit":
		c = (t.Value - 32) * 5 / 9
	case "Rankine":
		c = (t.Value - 491.67) * 5 / 9
	default:
		c = t.Value
	}
	if c < -273.15 {
		return Temperature{}, errors.New("temperature can't go below absolute zero")
	}
	return Temperature{c, "Celsius", "C"}, nil
}
func (t Temperature) ToFarenheit() (Temperature, error) {
	var f float64
	switch t.Unit {
	case "Celsius":
		f = t.Value*9/5 + 32
	case "Kelvin":
		f = (t.Value-273.15)*9/5 + 32
	case "Rankine":
		f = t.Value - 459.67
	default:
		f = t.Value
	}
	if f < -459.67 {
		return Temperature{}, errors.New("temperature can't go below absolute zero")
	}
	return Temperature{f, "Farenheit", "F"}, nil

}
func (t Temperature) ToRankine() (Temperature, error) {
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
	return Temperature{r, "Rankine", "Ra"}, nil

}

func main() {
	var boiling = Temperature{Value: 373.15, Unit: "Rankine", Symbol: "Ra"}
	var boiling2 = Temperature{Value: 100, Unit: "Celsius", Symbol: "C"}
	var boiling3 = Temperature{671.67, "Rankine", "Ra"}
	fmt.Println(boiling.ToKelvin())
	fmt.Println(boiling)
	fmt.Println(boiling2.ToRankine())
	fmt.Println(boiling3.ToCelsius())
	fmt.Println(description)
}
