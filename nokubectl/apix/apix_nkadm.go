package apix

import (
	"fmt"
	"os"
)

var NKADM_FLAGS = "" +
	"init              : initiate  (or re-init) a nokubeadm config and runtime directory" + "\n" +
	"init-npia         : fully initiate  (or re-init) a nokubeadm config and runtime directory" + "\n" +
	"init-npia-default : fully and automatically initiate  (or re-init) a nokubeadm config and runtime directory" + "\n" +
	//	"interactive       : enter nokubeadm interactive mode" + "\n" +
	"debug             : debug whatever" + "\n" +
	""

func GetNKADMFlagAndReduceArgs() (string, []string, error) {

	var flag string = ""

	var args []string

	nkadm_flag_detected := 0

	if len(os.Args) < 2 {
		return flag, args, fmt.Errorf("error cmd args: %s", "no args")
	}

	osargs := os.Args[1:]

	for i := 0; i < len(osargs); i++ {

		oa := osargs[i]

		_, okay := NKADMflag[oa]

		if okay {
			if nkadm_flag_detected == 1 {
				return flag, args, fmt.Errorf("error cmd args: %s", "nokubeadm special args cannot be joined like: "+flag+" "+oa)
			}
			nkadm_flag_detected = 1
			flag = oa
		} else {

			args = append(args, oa)

		}

	}

	return flag, args, nil
}
