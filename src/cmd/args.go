package cmd

import "fmt"

type Args struct {
	ClassPath string
}

func ParseArgs(args []string) (Args, error) {
	if len(args) != 2 {
		return Args{}, fmt.Errorf("argument error: expected 1 argument, got %d", len(args))
	}

	return Args{ClassPath: args[1]}, nil
}
