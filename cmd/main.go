package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/themilar/vert"
)

func convertInputToTemp(v string, u string) (*vert.Temperature, error) {
	if v, err := strconv.ParseFloat(v, 64); err == nil {
		switch u {
		case "k":
			return &vert.Temperature{Value: v, Unit: "Kelvin", Symbol: "K"}, nil
		case "c":
			return &vert.Temperature{Value: v, Unit: "Celsius", Symbol: "C"}, nil
		case "f":
			return &vert.Temperature{Value: v, Unit: "Farenheit", Symbol: "F"}, nil
		case "r":
			return &vert.Temperature{Value: v, Unit: "Rankine", Symbol: "R"}, nil
		default:
			// working with pointers might have made this more convenient=> return nil instead of an empty struct
			return nil, errors.New("invalid unit " + strconv.FormatFloat(v, 'f', -1, 64))
		}
	}
	return &vert.Temperature{}, errors.New("invalid float value")
}

func main() {
	// unit := flag.String("unit", "kelvin", "set the unit to convert your temperatures into")
	var kelvin bool
	flag.BoolVar(&kelvin, "kelvin", true, "convert the temperatures provided to kelvin")
	flag.BoolVar(&kelvin, "k", true, "convert the temperatures provided to kelvin")
	flag.Parse()
	if len(os.Args) == 1 {
		fmt.Println("insufficient arguments")
	} else {
		if kelvin {
			for _, x := range os.Args[2:] {
				// x[len(x)-1] this returns the unicode
				value, unit := x[:len(x)-1], x[len(x)-1:]
				// vert.Description[unit].Value=value
				t, err := convertInputToTemp(value, unit)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				} else {
					if err := t.ToKelvin(); err == nil {
						fmt.Println(value, unit, t.Value)
					} else {
						fmt.Fprintln(os.Stderr, err)
					}
				}

			}
		}
	}
}
