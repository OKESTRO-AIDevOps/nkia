package runtimefs

import (
	"encoding/json"
	"fmt"
	"os"

	pkgutils "github.com/OKESTRO-AIDevOps/nkia/pkg/utils"
	goya "github.com/goccy/go-yaml"
)

func OpenFilePointersForUsrBuildPipelineController() (*BuildPipelineController, error) {

	var pipe_controller BuildPipelineController

	var bp BuildPipeline

	file_byte, err := os.ReadFile(".usr/target/.npia/build.yaml")

	if err != nil {
		return &pipe_controller, fmt.Errorf("failed to get build yaml: %s", err.Error())
	}

	err = goya.Unmarshal(file_byte, &bp)

	if err != nil {
		return &pipe_controller, fmt.Errorf("failed to get build yaml: %s", err.Error())
	}

	jobs_len := len(bp.Jobs)

	keys := make([]string, 0)

	for i := 0; i < jobs_len; i++ {

		if pkgutils.CheckIfSliceContains[string](keys, bp.Jobs[i].Name) {

			return &pipe_controller, fmt.Errorf("failed to parse build yaml: %s", "duplicate name")
		}

		keys = append(keys, bp.Jobs[i].Name)

	}

	build_log_paths, err := BuildOpenForwardWithNames(keys)

	if err != nil {

		return &pipe_controller, fmt.Errorf("failed to build open forward: %s", err.Error())
	}

	pipe_controller.Pipeline = bp

	for i := 0; i < jobs_len; i++ {

		pending := 0

		job_name := bp.Jobs[i].Name

		if bp.Jobs[i].Needs != "" {

			pending = 1

		}

		if _, err := os.Stat(build_log_paths[i]); err == nil {
			return &pipe_controller, fmt.Errorf("failed to get file pointer: %s", "another build in process")
		}

		outfile, err := os.Create(build_log_paths[i])

		if err != nil {
			return &pipe_controller, fmt.Errorf("failed to get file pointer: %s", err.Error())
		}

		pipe_controller.Jobs = append(pipe_controller.Jobs, struct {
			Name     string
			Status   int
			Pending  int
			Error    error
			LogFile  *os.File
			SIGTERM  chan int
			SIGABORT chan int
		}{
			Name:     job_name,
			Status:   0,
			Pending:  pending,
			Error:    nil,
			LogFile:  outfile,
			SIGTERM:  make(chan int),
			SIGABORT: make(chan int),
		})

	}

	return &pipe_controller, nil
}

func CloseFilePointerForUsrBuildPipelineLog(bpctl *BuildPipelineController, job_name string, close_sig int, close_msg string) error {

	job_len := len(bpctl.Jobs)

	job_idx := -1

	for i := 0; i < job_len; i++ {

		if bpctl.Jobs[i].Name == job_name {

			job_idx = i

			break

		}

	}

	if job_idx == -1 {

		return fmt.Errorf("failed to close: no such name: %s", job_name)

	}

	bpctl.Jobs[job_idx].Status = close_sig

	if close_sig < 0 {

		bpctl.Jobs[job_idx].Error = fmt.Errorf(close_msg)

	} else {

		bpctl.Jobs[job_idx].Error = nil

	}

	err := bpctl.Jobs[job_idx].LogFile.Close()

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	err = BuildCloseWithName(job_name, close_msg)

	if err != nil {
		return fmt.Errorf("failed to close build: %s", err.Error())
	}

	return nil
}

func GetUsrPipelineLog() ([]byte, error) {

	var ret_byte []byte

	var err error

	pipe_log_map := make(map[string]string)

	head_dir := ".usr/build/"

	head, err := os.ReadFile(".usr/build/HEAD")

	if err != nil {

		return ret_byte, fmt.Errorf("failed to get usr build log: %s", err.Error())

	}

	head_value := string(head)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get usr build log: %s", err.Error())
	}

	head_dir += head_value + "/"

	var dir_list []string

	dir_entry, err := os.ReadDir(head_dir)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get usr build log: %s", err.Error())
	}

	for i := 0; i < len(dir_entry); i++ {

		if dir_entry[i].IsDir() {

			dir_list = append(dir_list, dir_entry[i].Name())

		}

	}

	for i := 0; i < len(dir_list); i++ {

		target_dir_log := head_dir + dir_list[i] + "/log"

		log_b, err := os.ReadFile(target_dir_log)

		if err != nil {

			return ret_byte, fmt.Errorf("failed to get log: at %s :%s", dir_list[i], err.Error())

		}

		pipe_log_map[dir_list[i]] = string(log_b)

	}

	ret_byte, err = json.Marshal(pipe_log_map)

	if err != nil {

		return ret_byte, fmt.Errorf("failed to get log: marshal: %s", err.Error())

	}

	return ret_byte, nil
}
