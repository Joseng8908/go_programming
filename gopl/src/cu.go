package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"gopl/lengthconv"
	"gopl/tempconv"
	"gopl/weightconv"
)

func main() {
	tempFlag := flag.String("t", "", "temperature value")
	lengthFlag := flag.String("l", "", "length value")
	weightFlag := flag.String("w", "", "weight value")
	flag.Parse()

	if *tempFlag != "" {
		num, unit := parseValue(*tempFlag)

		switch unit {
		case "C":
			fmt.Printf("%gC = %gF", num, tempconv.CToF(num))
		case "F":
			fmt.Printf("%gF = %gC", num, tempconv.FToC(num))
		default:
			log.Fatal("unknown temperature unit: %s.", unit)
		}
	}

	if *lengthFlag != "" {
		num, unit := parseValue(*lengthFlag)

		switch unit {
		case "m":
			fmt.Printf("%gm = %gft", num, lengthconv.MToFt(num))
		case "ft":
			fmt.Printf("%gft = %gm", num, lengthconv.FtToM(num))
		default:
			log.Fatal("unknown lenghth unit: %s.", unit)
		}
	}

	if *weightFlag != "" {
		num, unit := parseValue(*weightFlag)

		switch unit {
		case "lb":
			fmt.Printf("%glb = %gkg", num, weightconv.LbToKg(num))
		case "kg":
			fmt.Printf("%gkg = %glb", num, weightconv.KgToLb(num))
		default:
			log.Fatal("unknown weight unit: %s.", unit)
		}
	}
}

func parseValue(s string) (float64, string) {
	if len(s) < 2 {
		log.Fatal("invalid value: %s", s)
	}

	numPart := s[:len(s)-1]
	unitPart := s[len(s)-1:]
	num, err := strconv.ParseFloat(numPart, 64)
	if err != nil {
		log.Fatalf("invalid number: %v", err)
	}
	return num, strings.ToUpper(unitPart)
}
