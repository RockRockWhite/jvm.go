package main

import "github.com/RockRockWhite/jvm.go/src/cmd"

func main() {
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
