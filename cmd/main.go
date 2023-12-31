package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/themilar/vert"
)

const usage = `
Usage of vert command line tool:
  -c, --celsius		convert the temperatures provided to celsius
  -f, --fahrenheit	convert the temperatures provided to fahrenheit
  -k, --kelvin		convert the temperatures provided to kelvin
  -r, --rankine		convert the temperatures provided to rankine
`

func convertInputToTemp(v string, u string) (vert.Temperature, error) {
	if v, err := strconv.ParseFloat(v, 64); err == nil {
		switch u {
		case "k":
			return vert.Temperature{Value: v, Unit: "Kelvin"}, nil
		case "c":
			return vert.Temperature{Value: v, Unit: "Celsius"}, nil
		case "f":
			return vert.Temperature{Value: v, Unit: "Fahrenheit"}, nil
		case "r":
			return vert.Temperature{Value: v, Unit: "Rankine"}, nil
		default:
			// working with pointers might have made this more convenient=> return nil instead of an empty struct
			return vert.Temperature{}, fmt.Errorf("invalid unit: %v", v)
		}
	}
	return vert.Temperature{}, fmt.Errorf("invalid float value: %v", v)
}

func main() {
	// unit := flag.String("unit", "kelvin", "set the unit to convert your temperatures into")
	var kelvin, celsius, fahrenheit, rankine bool

	flag.BoolVar(&kelvin, "kelvin", false, "convert the temperatures provided to kelvin")
	flag.BoolVar(&kelvin, "k", false, "convert the temperatures provided to kelvin")
	flag.BoolVar(&celsius, "celsius", false, "convert the temperatures provided to celsius")
	flag.BoolVar(&celsius, "c", false, "convert the temperatures provided to celsius")
	flag.BoolVar(&fahrenheit, "fahrenheit", false, "convert the temperatures provided to fahrenheit")
	flag.BoolVar(&fahrenheit, "f", false, "convert the temperatures provided to fahrenheit")
	flag.BoolVar(&rankine, "rankine", false, "convert the temperatures provided to rankine")
	flag.BoolVar(&rankine, "r", false, "convert the temperatures provided to rankine")
	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("insufficient arguments")
		fmt.Print(usage)
	} else {
		switch {
		case kelvin:
			fmt.Println("Temperatures in kelvin")
			for _, x := range os.Args[2:] {
				// x[len(x)-1] this returns the unicode
				value, unit := x[:len(x)-1], x[len(x)-1:]
				// vert.Description[unit].Value=value
				t, err := convertInputToTemp(value, unit)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				} else {
					t, err = t.ToKelvin()
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
					} else {
						fmt.Println(value, unit, fmt.Sprintf("%v%v", t.Value, t.Symbol))
					}
				}

			}
			fmt.Println("\xE2\x9C\x94 conversion complete")
		case celsius:
			fmt.Println("Temperatures in celsius")
			for _, x := range os.Args[2:] {
				value, unit := x[:len(x)-1], x[len(x)-1:]
				t, err := convertInputToTemp(value, unit)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				} else {
					t, err = t.ToCelsius()
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
					} else {
						fmt.Println(value, unit, fmt.Sprintf("%v%v", t.Value, t.Symbol))
					}
				}

			}
			fmt.Println("\xE2\x9C\x94 conversion complete")
		case fahrenheit:
			fmt.Println("Temperatures in fahrenheit")
			for _, x := range os.Args[2:] {
				value, unit := x[:len(x)-1], x[len(x)-1:]
				t, err := convertInputToTemp(value, unit)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				t, err = t.ToFahrenheit()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				} else {
					fmt.Println(value, unit, fmt.Sprintf("%v%v", t.Value, t.Symbol))
				}

			}
			fmt.Println("\xE2\x9C\x94 conversion complete")
		case rankine:
			fmt.Println("Temperatures in rankine")
			for _, x := range os.Args[2:] {
				value, unit := x[:len(x)-1], x[len(x)-1:]
				t, err := convertInputToTemp(value, unit)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				} else {
					t, err = t.ToRankine()
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
					} else {
						fmt.Println(value, unit, fmt.Sprintf("%v%v", t.Value, t.Symbol))
					}
				}

			}
			fmt.Println("\xE2\x9C\x94 conversion complete")
		default:
			fmt.Print("No flags set")
		}
	}
}
