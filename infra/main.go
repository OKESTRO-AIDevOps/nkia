package main

import (
	"fmt"

	ci "github.com/OKESTRO-AIDevOps/nkia/infra/ci"
	cicmd "github.com/OKESTRO-AIDevOps/nkia/infra/cmd"
	cigit "github.com/OKESTRO-AIDevOps/nkia/infra/git"
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

	err = cigit.GetSourceRepo(ci_opts)

	if err != nil {

		fmt.Println(err.Error())

		return
	}

	targets_ctl := ci.CITargetsFactory()

	cred_store := ci.CICredFactory()

	err = ci.StoreCICredFromCIFile(cred_store, "stdin")

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	if err := ci.StartTargetsFromCIFile(targets_ctl, cred_store); err != nil {

		fmt.Println(err.Error())

		return
	}

	if err := ci.TargetsControllerStdin(targets_ctl); err != nil {

		fmt.Println(err.Error())

		return
	}

}
