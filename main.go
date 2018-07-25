package main

import (
	"log"

	"github.com/hellerox/binancecli/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatalf("Problems: %s", err.Error())
	}

}
