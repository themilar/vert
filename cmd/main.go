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
					fmt.Fprint(os.Stderr, err)
				}
				t, _ = t.ToKelvin()
				fmt.Println(value, unit, t.Value)

			}
		}
	}
}
