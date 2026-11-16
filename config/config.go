package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	WorkTime   uint32
	RelaxTime  uint32
	Iterations uint32
}

func ParseArgs() (Config, error) {
	var cfg Config

	args := os.Args[1:]

	if len(args) != 3 {
		return cfg, fmt.Errorf("usage: %s <work_min> <relax_time> <iterations>", os.Args[0])
	}

	vals := make([]uint32, 3)

	for i, a := range args {
		v, err := strconv.Atoi(a)

		if err != nil {
			return cfg, fmt.Errorf("argument %d (%q) is not a number: %w", i+1, a, err)
		}

		if v <= 0 {
			return cfg, fmt.Errorf("argument %d must be positive, got %d", i+1, v)
		}

		vals[i] = uint32(v)

	}

	return Config{
		WorkTime:   vals[0],
		RelaxTime:  vals[1],
		Iterations: vals[2],
	}, nil
}
