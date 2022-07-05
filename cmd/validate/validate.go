package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"eirclode.voy.technology/data"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	r := csv.NewReader(bytes.NewReader(data.Data))
	num := 0
	for {
		line, err := r.Read()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return lineErr(num, "@", "", err)
		}

		num++

		if err := checkEircode(line[0]); err != nil {
			return lineErr(num, "eircode", line[0], err)
		}
		if err := checkLattitude(line[1]); err != nil {
			return lineErr(num, "lattitude", line[1], err)
		}
		if err := checkLongitude(line[2]); err != nil {
			return lineErr(num, "longitude", line[2], err)
		}
	}
}

func checkEircode(eircode string) error {
	if len(eircode) != 7 {
		return fmt.Errorf("invalid eircode length")
	}
	return nil
}

func checkLattitude(lat string) error {
	return checkCoordinate(lat, 50, 56)
}

func checkLongitude(lon string) error {
	return checkCoordinate(lon, -12, -5)
}

func checkCoordinate(c string, min, max float64) error {
	// Check is parsable
	l, err := strconv.ParseFloat(c, 32)
	if err != nil {
		return fmt.Errorf("unable to parse: %w", err)
	}

	// Check is in ireland
	if !(min < l && l < max) {
		return fmt.Errorf("coordinate is not in Ireland, %.2f < %.2f < %.2f", min, l, max)
	}

	return nil
}

func lineErr(line int, field, value string, err error) error {
	return fmt.Errorf("::error file=data/eircodes.csv,line=%d::%s (%s) = %w", line, field, value, err)
}
