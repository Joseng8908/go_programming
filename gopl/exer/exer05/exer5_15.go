package main

import (
	"errors"
	"sort"
)

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("min: at least one argument")
	}
	sort.Ints(vals)
	return vals[0], nil
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("max: at least one argument")
	}
	sort.Ints(vals)
	return vals[len(vals)-1], nil
}
