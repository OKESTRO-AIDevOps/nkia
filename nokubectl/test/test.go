package main

import (
	"fmt"

	apix "github.com/OKESTRO-AIDevOps/nkia/nokubectl/apix"
)

func cmdBuildTest() error {

	flag, args, err := apix.GetNKCTLFlagAndReduceArgs()

	if flag != "" {
		fmt.Println("nkctl: " + flag)
		return nil
	}

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
