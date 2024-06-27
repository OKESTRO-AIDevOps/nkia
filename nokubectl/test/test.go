package main

import (
	"fmt"
	"os"

	apix "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard/apix"
)

func cmdBuildTest() error {

	args := os.Args[1:]

	oreq, err := apix.AXgi.BuildOrchRequestFromCommandLine(args)

	if err != nil {

		return fmt.Errorf("failed: %s", err.Error())
	}

	fmt.Println(oreq)

	return nil

}

func main() {

	if err := cmdBuildTest(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("success")
	}

}
