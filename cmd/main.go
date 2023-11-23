package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/themilar/vert"
)

func convertSToT(v float64, u string) vert.Temperature {
	switch u {
	case "k":
		return vert.Temperature{v, "Kelvin", "K"}
	case "c":
		return vert.Temperature{v, "Celsius", "C"}
	case "f":
		return vert.Temperature{v, "Farenheit", "F"}
	default:
		return vert.Temperature{}
	}
}
func main() {
	// unit := flag.String("unit", "kelvin", "set the unit to convert your temperatures into")
	k := flag.Bool("k", true, "convert the temperatures provided to kelvin")
	flag.Parse()
	if len(os.Args) == 1 {
		fmt.Println("insufficient arguments")
	} else {
		if *k {
			for _, x := range os.Args[2:] {
				// x[len(x)-1] this returns the unicode
				value, unit := x[:len(x)-1], x[len(x)-1:]
				// vert.Description[unit].Value=value
				if v, err := strconv.ParseFloat(value, 64); err == nil {
					t, _ := convertSToT(v, unit).ToKelvin()
					fmt.Println(value, unit, t.Value)
				}
			}
		}
	}
}
