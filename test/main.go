package main

import (
	"fmt"

	cicmd "github.com/OKESTRO-AIDevOps/nkia/test/cmd"
	cigit "github.com/OKESTRO-AIDevOps/nkia/test/git"
)

func main() {

	flag, ci_opts, err := cicmd.CmdParseArgs()

	if err != nil {

		fmt.Println(err.Error())

		return
	}

	if flag == "yaml" {

		// something to fill up ci_opts from yaml file

		fmt.Println("not implemented yet!! :)")

		return

	}

	err = cigit.GetRepo(ci_opts)

	if err != nil {

		fmt.Println(err.Error())

		return
	}

}
