package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RockRockWhite/jvm.go/src/class_file"
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

	cf := class_file.NewClassFile(data)
	class_info, err := cf.IntoClassInfo()
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(class_info, "", "  ")
	fmt.Println(string(b))

	return nil
}
