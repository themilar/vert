package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/themilar/vert"
)

func convertInputToTemp(v string, u string) (vert.Temperature, error) {
	if v, err := strconv.ParseFloat(v, 64); err == nil {
		switch u {
		case "k":
			return vert.Temperature{v, "Kelvin", "K"}, nil
		case "c":
			return vert.Temperature{v, "Celsius", "C"}, nil
		case "f":
			return vert.Temperature{v, "Farenheit", "F"}, nil
		case "r":
			return vert.Temperature{v, "Rakine", "R"}, nil
		default:
			// working with pointers might have made this more convenient=> return nil instead of an empty struct
			return vert.Temperature{}, errors.New("invalid unit")
		}
	}
	return vert.Temperature{}, errors.New("invalid float value")
}
func main() {
	// unit := flag.String("unit", "kelvin", "set the unit to convert your temperatures into")
	var kelvin, celsius, farenheit, rakine bool
	flag.BoolVar(&kelvin, "kelvin", false, "convert the temperatures provided to kelvin")
	flag.BoolVar(&kelvin, "k", false, "convert the temperatures provided to kelvin")
	flag.BoolVar(&celsius, "celsius", false, "convert the temperatures provided to celsius")
	flag.BoolVar(&celsius, "c", false, "convert the temperatures provided to celsius")
	flag.BoolVar(&farenheit, "farenheit", false, "convert the temperatures provided to farenheit")
	flag.BoolVar(&farenheit, "f", false, "convert the temperatures provided to farenheit")
	flag.BoolVar(&rakine, "rakine", false, "convert the temperatures provided to rakine")
	flag.BoolVar(&rakine, "r", false, "convert the temperatures provided to rakine")
	flag.Parse()
	if len(os.Args) == 1 {
		fmt.Println("insufficient arguments")
	} else {
		switch {
		case kelvin:
			for _, x := range os.Args[2:] {
				// x[len(x)-1] this returns the unicode
				value, unit := x[:len(x)-1], x[len(x)-1:]
				// vert.Description[unit].Value=value
				t, err := convertInputToTemp(value, unit)
				if err != nil {
					fmt.Fprint(os.Stderr, err)
				}
				t, _ = t.ToKelvin()
				fmt.Println(value, unit, t.Value)

			}
		case celsius:
			for _, x := range os.Args[2:] {
				value, unit := x[:len(x)-1], x[len(x)-1:]
				t, err := convertInputToTemp(value, unit)
				if err != nil {
					fmt.Fprint(os.Stderr, err)
				}
				t, _ = t.ToCelsius()
				fmt.Println(value, unit, t.Value)

			}
		case farenheit:
			for _, x := range os.Args[2:] {
				value, unit := x[:len(x)-1], x[len(x)-1:]
				t, err := convertInputToTemp(value, unit)
				if err != nil {
					fmt.Fprint(os.Stderr, err)
				}
				t, _ = t.ToFarenheit()
				fmt.Println(value, unit, t.Value)

			}
		case rakine:
			for _, x := range os.Args[2:] {
				value, unit := x[:len(x)-1], x[len(x)-1:]
				t, err := convertInputToTemp(value, unit)
				if err != nil {
					fmt.Fprint(os.Stderr, err)
				}
				t, _ = t.ToRakine()
				fmt.Println(value, unit, t.Value)

			}

		}
	}
}
