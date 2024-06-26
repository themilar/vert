package vert

import "fmt"

type Temperature struct {
	Value        float64
	Unit, Symbol string
}
type AbsoluteZeroError float64

func (f AbsoluteZeroError) Error() string {
	return fmt.Sprintf("%g: temperature can't go below absolute zero", f)
}

// methods
func (t Temperature) ToKelvin() (Temperature, error) {
	var k float64
	switch t.Unit {
	case "Celsius":
		k = t.Value + 273.15
	case "Fahrenheit":
		k = (t.Value-32)*5/9 + 273.15
	case "Rankine", "rankine":
		k = t.Value * 5 / 9
	default:
		k = t.Value
	}
	if k < 0 {
		// var ae AbsoluteZeroError = k
		return Temperature{}, AbsoluteZeroError(k)
	}
	return Temperature{k, "Kelvin", "K"}, nil
}
func (t Temperature) ToCelsius() (Temperature, error) {
	var c float64
	switch t.Unit {
	case "Kelvin":
		c = t.Value - 273.15
	case "Fahrenheit":
		c = (t.Value - 32) * 5 / 9
	case "Rankine":
		c = (t.Value - 491.67) * 5 / 9
	default:
		c = t.Value
	}
	if c < -273.15 {
		return Temperature{}, AbsoluteZeroError(c)
	}
	return Temperature{c, "Celsius", "\u00B0C"}, nil
}
func (t Temperature) ToFahrenheit() (Temperature, error) {
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
		return Temperature{}, AbsoluteZeroError(f)
	}
	return Temperature{f, "Fahrenheit", "\u00B0F"}, nil

}
func (t Temperature) ToRankine() (Temperature, error) {
	var r float64
	switch t.Unit {
	case "Celsius":
		r = t.Value*9/5 + 491.67
	case "Kelvin":
		r = t.Value * 1.8
	case "Fahrenheit":
		r = t.Value + 459.67
	default:
		r = t.Value
	}
	if r < 0 {
		return Temperature{}, AbsoluteZeroError(r)
	}
	return Temperature{r, "Rankine", "\u00B0Ra"}, nil

}
