package kubetoolkit

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"

	runfs "github.com/OKESTRO-AIDevOps/nkia/pkg/runtimefs"
	goya "github.com/goccy/go-yaml"
)

func PipelineBuildStart(main_ns string, repoaddr string, regaddr string) {

	app_origin, _ := runfs.LoadAdmOrigin()

	ns_found, _, _ := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {

		return
	}

	err := runfs.InitUsrTargetForPipeline(repoaddr)

	if err != nil {

		return
	}

	bpctl, err := runfs.OpenFilePointersForUsrBuildPipelineController()

	if err != nil {
		return
	}

	RunPipe(bpctl)

	return
}

func PipelineBuildStart_Test() {

	bpctl, err := runfs.OpenFilePointersForUsrBuildPipelineController()

	if err != nil {
		return
	}

	RunPipe(bpctl)

	return
}

func PipelineBuildSetVariablesEx(varnm string, varval string) ([]byte, error) {

	status := "failed"

	var_list := make(map[string]string)

	var_list[varnm] = varval

	err := PipelineBuildSetVariables(var_list)

	if err != nil {

		return []byte(status), fmt.Errorf("failed to set var: %s", err.Error())
	}

	status = "success"

	return []byte(status), nil

}

func PipelineBuildSetVariables(var_list map[string]string) error {

	varmap, _ := PipelineBuildGetVariableMap()

	for k, v := range var_list {

		varmap[k] = v

	}

	new_b, err := goya.Marshal(varmap)

	if err != nil {

		return fmt.Errorf("failed to set variable: %s", err.Error())
	}

	err = runfs.CreateUsrPipelineVariables(new_b)

	if err != nil {

		return fmt.Errorf("failed to write new variables: %s", err.Error())
	}

	return nil
}

func PipelineBuildGetVariableMapEx() ([]byte, error) {

	var ret []byte

	varmap, err := PipelineBuildGetVariableMap()

	if err != nil {

		return ret, fmt.Errorf("failed to get variable map: %s", err.Error())

	}

	ret, err = goya.Marshal(varmap)

	if err != nil {

		return ret, fmt.Errorf("failed to get variable map: %s", err.Error())
	}

	return ret, nil

}

func PipelineBuildGetVariableMap() (map[string]string, error) {

	varmap := make(map[string]string)

	var_b, err := runfs.LoadUsrPipelineVariables()

	if err != nil {

		return varmap, fmt.Errorf("failed to load build var: %s", err.Error())

	}

	err = goya.Unmarshal(var_b, &varmap)

	if err != nil {

		return varmap, fmt.Errorf("failed to marshalr build var: %s", err.Error())

	}

	return varmap, nil

}

func PipelineBuildGetLog() ([]byte, error) {

	var ret_byte []byte

	ret_byte, err := runfs.GetUsrPipelineLog()

	if err != nil {

		return ret_byte, fmt.Errorf("failed to get pipe build log: %s", err.Error())

	}

	return ret_byte, nil
}

func GetEffectiveCommands(raw_cmd string) []string {

	cmd_args := strings.Fields(raw_cmd)

	cmd_len := len(cmd_args)

	for i := 0; i < cmd_len; i++ {

		replace_flag := 0

		if strings.HasPrefix(cmd_args[i], "$(") && strings.HasSuffix(cmd_args[i], ")") {

			replace_flag = 1

		}

		if replace_flag == 1 {

			stripped_arg := strings.ReplaceAll(cmd_args[i], "$(", "")

			stripped_arg = strings.ReplaceAll(stripped_arg, ")", "")

			new_arg, err := GetPipeVariable(stripped_arg)

			if err != nil {

				continue

			}

			cmd_args[i] = new_arg

		}

	}

	return cmd_args

}

func GetPipeVariable(key string) (string, error) {

	var ans string

	varmap := make(map[string]string)

	var_b, err := runfs.LoadUsrPipelineVariables()

	if err != nil {

		return ans, fmt.Errorf("failed to get pipe var: %s", err.Error())

	}

	err = goya.Unmarshal(var_b, &varmap)

	if err != nil {

		return ans, fmt.Errorf("failed to make map: %s", err.Error())

	}

	ans, okay := varmap[key]

	if !okay {

		return ans, fmt.Errorf("failed to retrieve key val for: %s", key)

	}

	return ans, nil

}

func RunPipe(bpctl *runfs.BuildPipelineController) {

	job_len := len(bpctl.Jobs)

	for i := 0; i < job_len; i++ {

		if bpctl.Jobs[i].Pending != 1 {

			go RunJob(bpctl, i)

		}

	}

	success_target := job_len
	success_count := 0

	for {

		for i := 0; i < job_len; i++ {

			if bpctl.Jobs[i].Status != 0 {

				continue
			}

			select {

			case terminate := <-bpctl.Jobs[i].SIGTERM:

				if terminate < 0 {

					common_msg := fmt.Sprintf("SIGTERM %d from job id: %d\n", terminate, i)

					AbortAll(bpctl, common_msg)

					return
				}

				success_count += 1

				next_jobs := NextByJobId(bpctl, i)

				if len(next_jobs) == 0 {
					break
				}

				for j := 0; j < len(next_jobs); j++ {

					bpctl.Jobs[next_jobs[j]].Pending = 0

					go RunJob(bpctl, next_jobs[j])

				}

				break

			}

		}

		if success_count == success_target {

			return

		}

	}

}

func AbortAll(bpctl *runfs.BuildPipelineController, common_msg string) {

	job_len := len(bpctl.Jobs)

	for i := 0; i < job_len; i++ {

		job_name := bpctl.Jobs[i].Name

		bpctl.Jobs[i].SIGABORT <- -1

		runfs.CloseFilePointerForUsrBuildPipelineLog(bpctl, job_name, -1, common_msg)

	}

}

func RunJob(bpctl *runfs.BuildPipelineController, job_id int) {

	job_name := bpctl.Pipeline.Jobs[job_id].Name

	command_list := bpctl.Pipeline.Jobs[job_id].Steps

	com_len := len(command_list)

	for i := 0; i < com_len; i++ {

		cmd_args := GetEffectiveCommands(command_list[i])

		var cmd *exec.Cmd

		if len(cmd_args) == 1 {

			cmd = exec.Command(cmd_args[0])

		} else {

			cmd = exec.Command(cmd_args[0], cmd_args[1:]...)

		}

		cmd_out, _ := cmd.StdoutPipe()
		cmd_err, _ := cmd.StderrPipe()

		go func() {
			merged := io.MultiReader(cmd_out, cmd_err)
			scanner := bufio.NewScanner(merged)
			for scanner.Scan() {
				msg := scanner.Text()
				bpctl.Jobs[job_id].LogFile.Write([]byte(msg))
			}
		}()

		err := cmd.Run()

		if err != nil {

			runfs.CloseFilePointerForUsrBuildPipelineLog(bpctl, job_name, -1, err.Error())

			bpctl.Jobs[job_id].SIGTERM <- -1

			return
		}

	}

	runfs.CloseFilePointerForUsrBuildPipelineLog(bpctl, job_name, 1, "success")

	bpctl.Jobs[job_id].SIGTERM <- 1

	return

}

func NextByJobId(bpctl *runfs.BuildPipelineController, job_id int) []int {

	var next_ids []int

	this_name := bpctl.Pipeline.Jobs[job_id].Name

	job_len := len(bpctl.Pipeline.Jobs)

	for i := 0; i < job_len; i++ {

		if bpctl.Pipeline.Jobs[i].Needs == this_name {

			next_ids = append(next_ids, i)
		}

	}

	return next_ids

}
