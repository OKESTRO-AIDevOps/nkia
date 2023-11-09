package apix

import (
	"fmt"
	"strings"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/nokubelet/controller"
	apistd "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
)

func (axgi API_X) BuildOrchRequest(apix_id string, apix_options API_X_OPTIONS) (ctrl.OrchestratorRequest, error) {

	oreq := ctrl.OrchestratorRequest{}

	apistd_in := make(apistd.API_INPUT)

	apistd_id, okay := axgi[apix_id]

	if !okay {
		return oreq, fmt.Errorf("builder: %s", "no matching api std for: "+apix_id)
	}

	apistd_in_template, okay := apistd.ASgi[apistd_id]

	if !okay {
		return oreq, fmt.Errorf("builder: %s", "no matching api std for: "+apistd_id)
	}

	for k, v := range apistd_in_template {

		if k == 0 {
			apistd_in[v] = apistd_id
			continue
		}

		input_val, okay := apix_options[v]

		if !okay {
			return oreq, fmt.Errorf("builder: %s", "missing required option: "+"--"+v)
		}

		apistd_in[v] = input_val

	}

	linear_inst := LinearInstructionBuildFromAPIInput(apistd_in, apistd_in_template)

	oreq.Query = linear_inst

	to, to_okay := apix_options["to"]

	as, as_okay := apix_options["as"]

	if to_okay {
		oreq.RequestTarget = to
	}

	if as_okay {
		oreq.RequestOption = as
	}

	return oreq, nil
}

func (axgi API_X) BuildAPIInput(apix_id string, apix_options API_X_OPTIONS) (apistd.API_INPUT, error) {

	apistd_in := make(apistd.API_INPUT)

	apistd_id, okay := axgi[apix_id]

	if !okay {
		return apistd_in, fmt.Errorf("builder: %s", "no matching api std for: "+apix_id)
	}

	apistd_in_template, okay := apistd.ASgi[apistd_id]

	if !okay {
		return apistd_in, fmt.Errorf("builder: %s", "no matching api std for: "+apistd_id)
	}

	for k, v := range apistd_in_template {

		if k == 0 {
			apistd_in[v] = apistd_id
			continue
		}

		input_val, okay := apix_options[v]

		if !okay {
			return apistd_in, fmt.Errorf("builder: %s", "missing required option: "+"--"+v)
		}

		apistd_in[v] = input_val

	}

	return apistd_in, nil
}

func (axgi API_X) BuildOrchRequestFromCommandLine(args []string) (ctrl.OrchestratorRequest, error) {

	oreq := ctrl.OrchestratorRequest{}

	apix_id, apix_options, err := BuildCmdIdAndOptions(args)

	if err != nil {
		return oreq, fmt.Errorf("error cmd args: %s", err.Error())
	}

	oreq, err = axgi.BuildOrchRequest(apix_id, apix_options)

	if err != nil {
		return oreq, fmt.Errorf("error cmd args: %s", err.Error())
	}

	return oreq, nil
}

func (axgi API_X) BuildAPIInputFromCommandLine(args []string) (apistd.API_INPUT, error) {

	var api_input apistd.API_INPUT

	apix_id, apix_options, err := BuildCmdIdAndOptions(args)

	if err != nil {
		return api_input, fmt.Errorf("error cmd args: %s", err.Error())
	}

	api_input, err = axgi.BuildAPIInput(apix_id, apix_options)

	if err != nil {
		return api_input, fmt.Errorf("error cmd args: %s", err.Error())
	}

	return api_input, nil
}

func BuildCmdIdAndOptions(args []string) (string, API_X_OPTIONS, error) {

	var apix_id string

	apix_options := make(API_X_OPTIONS)

	boundary := len(args) - 1

	ptr := 0

	for ptr <= boundary {

		arg := args[ptr]

		if strings.HasPrefix(arg, "--") {

			if ptr+1 > boundary {

				return apix_id, apix_options, fmt.Errorf("failed to build cmd: %s", "missing flag value after flag: "+arg)

			}

			if strings.HasPrefix(args[ptr+1], "--") {

				return apix_id, apix_options, fmt.Errorf("failed to build cmd: %s", "missing flag value after flag: "+arg)

			}

			stripped_key := strings.ReplaceAll(arg, "--", "")

			val := args[ptr+1]

			apix_options[stripped_key] = val

			ptr += 2

		} else {

			arg := args[ptr]

			apix_id += arg + "-"

			ptr += 1
		}

	}

	apix_id = StripLast(apix_id)

	return apix_id, apix_options, nil

}

func LinearInstructionBuildFromAPIInput(std_cmd_in apistd.API_INPUT, apistd_template []string) string {

	var lininst string

	i := 0

	lidx := len(apistd_template) - 1

	for i <= lidx {

		key := apistd_template[i]

		if i == 0 {
			lininst = std_cmd_in[key] + ":"
		} else if i < lidx {

			lininst += std_cmd_in[key] + ","

		} else if i == lidx {

			lininst += std_cmd_in[key]

		}

		i++

	}

	return lininst
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
