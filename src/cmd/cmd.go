package cmd

import (
	"fmt"
	"os"
)

func Run() error {

	args, err := ParseArgs(os.Args)

	if err != nil {
		return err
	}

	data, err := os.ReadFile(args.ClassPath)
	if err != nil {
		return err
	}

	for _, b := range data {
		fmt.Printf("%02x ", b)
	}

	return nil
}
