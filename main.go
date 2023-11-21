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
func (t Temperature) convert() (float64, error) {
	var k float64
	switch t.Unit {
	case "Celsius":
		k = t.Value + 273.15
	case "Farenheit":
		k = (t.Value-32)*5/9 + 273.5
	case "Rakine":
		k = t.Value * 5 / 9
	default:
		k = t.Value
	}
	if k < 0 {
		return 0, errors.New("temperature can't go below absolute zero")
	}
	return k, nil
	// if t.Unit != "Kevin" {
	// 	k = t.Value + 273.15
	// 	if k < 0 {
	// 		return 0, errors.New("temperature can't go below absolute zero")
	// 	}
	// 	return k, nil
	// }
	// return t.Value, nil
}

func main() {
	var boiling = Temperature{Value: 100, Unit: "Celsius", Symbol: "C"}
	fmt.Println(boiling)
	fmt.Println(boiling.convert())
	fmt.Println(description)
}
