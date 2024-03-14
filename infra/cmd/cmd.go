package cmd

import (
	"fmt"
	"os"
	"strings"
)

func CmdParseArgs() (string, NKIA_CI_OPTIONS, error) {

	var flag string

	nkiaci_opts := make(NKIA_CI_OPTIONS)

	if len(os.Args) < 2 {

		return flag, nkiaci_opts, fmt.Errorf("cmd parse: %s", "no args specified")

	}

	os_args := os.Args[1:]

	boundary := len(os_args) - 1

	ptr := 0

	for ptr <= boundary {

		arg := os_args[ptr]

		if strings.HasPrefix(arg, "--") {

			if ptr+1 > boundary {

				return flag, nkiaci_opts, fmt.Errorf("failed to build cmd: %s", "missing flag value after flag: "+arg)

			}

			if strings.HasPrefix(os_args[ptr+1], "--") {

				return flag, nkiaci_opts, fmt.Errorf("failed to build cmd: %s", "missing flag value after flag: "+arg)

			}

			stripped_key := strings.ReplaceAll(arg, "--", "")

			val := os_args[ptr+1]

			if _, okay := NKIACIopts[stripped_key]; !okay {

				return flag, nkiaci_opts, fmt.Errorf("failed to build cmd: %s", "impossible key: "+stripped_key)

			}

			nkiaci_opts[stripped_key] = val

			ptr += 2

		}
	}

	if err := CheckNKIACIOptionValidity(nkiaci_opts); err != nil {

		return flag, nkiaci_opts, fmt.Errorf("invalid options: %s", err.Error())
	}

	if _, okay := nkiaci_opts["from-yaml"]; okay {

		flag = "yaml"

	} else {

		flag = "args"
	}

	return flag, nkiaci_opts, nil
}

func CheckNKIACIOptionValidity(opts NKIA_CI_OPTIONS) error {

	from_yaml := 0

	if _, okay := opts["from-yaml"]; okay {

		from_yaml = 1

	}

	if from_yaml == 1 {

		if _, okay := opts["repo"]; okay {

			return fmt.Errorf("from-yaml present, do not specify repo")

		}

		if _, okay := opts["id"]; okay {

			return fmt.Errorf("from-yaml present, do not specify id")

		}

		if _, okay := opts["token"]; okay {

			return fmt.Errorf("from-yaml present, do not specify id")

		}

	}

	return nil

}

func StripLast(in string) string {

	var ret string

	out_rune := make([]rune, 0)
	in_rune := []rune(in)

	last_idx := len(in_rune) - 1

	for i := 0; i < last_idx; i++ {

		out_rune = append(out_rune, in_rune[i])

	}

	ret = string(out_rune)

	return ret

}
