package kubetoolkit

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"

	runfs "github.com/OKESTRO-AIDevOps/nkia/pkg/runtimefs"
)

func PipelineBuildStart(main_ns string, repoaddr string, regaddr string) {

	bpctl, err := runfs.OpenFilePointersForUsrBuildPipelineController()

	if err != nil {
		return
	}

	app_origin, _ := runfs.LoadAdmOrigin()

	ns_found, _, _ := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {

		close_msg := "namespace not found in ADMorigin\n"

		AbortAll(bpctl, close_msg)

		return
	}

	err = runfs.InitUsrTarget(repoaddr)

	if err != nil {

		AbortAll(bpctl, err.Error())
		return
	}

	RunPipe(bpctl)

	return
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

		cmd_args := strings.Fields(command_list[i])

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

	runfs.CloseFilePointerForUsrBuildPipelineLog(bpctl, job_name, 1, "success\n")

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
