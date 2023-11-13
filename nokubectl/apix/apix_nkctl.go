package apix

import (
	"fmt"
	"os"
)

var NKCTL_FLAGS = "" +
	"init              : initiate  (or re-init) a nokubectl config and runtime directory" + "\n" +
	"read              : read out the previous query result if present" + "\n" +
	//	"help              : given command line argument, prints out relevant information" + "\n" +
	//	"interactive       : (recommended) run command in interactive mode" + "\n" +
	"apix-md           : exports all apix information to a markdown file" + "\n" +
	"apix-js           : exports all apix information to an importable js file" + "\n" +
	"apix-py           : exports all apix information to an importable py file" + "\n" +
	""

func GetNKCTLFlagAndReduceArgs() (string, []string, error) {

	var flag string = ""

	var args []string

	nkctl_flag_detected := 0

	if len(os.Args) < 2 {
		return flag, args, fmt.Errorf("error cmd args: %s", "no args")
	}

	osargs := os.Args[1:]

	for i := 0; i < len(osargs); i++ {

		oa := osargs[i]

		_, okay := NKCTLflag[oa]

		if okay {
			if nkctl_flag_detected == 1 {
				return flag, args, fmt.Errorf("error cmd args: %s", "nokubectl special args cannot be joined like: "+flag+" "+oa)
			}

			flag = oa
		} else {

			args = append(args, oa)

		}

	}

	return flag, args, nil
}
