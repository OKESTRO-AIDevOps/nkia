package runtimefs

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type BuildPipeline struct {
	Jobs []struct {
		Name  string   `yaml:"name"`
		Steps []string `yaml:"steps"`
		Needs string   `yaml:"needs,omitempty"`
	} `yaml:"jobs"`
}
type BuildPipelineController struct {
	Jobs []struct {
		Name     string
		Status   int
		Pending  int
		Error    error
		LogFile  *os.File
		SIGTERM  chan int
		SIGABORT chan int
	}
	Pipeline BuildPipeline
}

func BuildOpenForwardWithNames(job_names []string) ([]string, error) {

	var build_dirs []string

	ERR_MSG := "failed build open forward: %s"

	head_value := "0"

	head_value_int := 0

	head_dir := ".usr/build/"

	new_head_dir := ".usr/build/"

	cmd := exec.Command("mkdir", "-p", ".usr/build")

	_, err := cmd.Output()

	if err != nil {
		return build_dirs, fmt.Errorf(ERR_MSG, err.Error())
	}

	if _, err := os.Stat(".usr/build/HEAD"); err != nil {

		for i := 0; i < len(job_names); i++ {

			cmd := exec.Command("mkdir", "-p", ".usr/build/0/"+job_names[i])

			_, _ = cmd.Output()

			t_now := time.Now()

			t_str := t_now.Format("2006-01-02-15-04-05")

			_ = os.WriteFile(".usr/build/0/"+job_names[i]+"/open", []byte(t_str), 0644)

			build_dirs = append(build_dirs, ".usr/build/0/"+job_names[i]+"/log")
		}

		_ = os.WriteFile(".usr/build/HEAD", []byte("0"), 0644)

		return build_dirs, nil
	}

	head, err := os.ReadFile(".usr/build/HEAD")

	if err != nil {
		return build_dirs, fmt.Errorf(ERR_MSG, err.Error())
	}

	head_value = string(head)

	head_value_int, err = strconv.Atoi(head_value)

	if err != nil {
		return build_dirs, fmt.Errorf(ERR_MSG, err.Error())
	}

	for i := 0; i < len(job_names); i++ {

		head_dir = ".usr/build/"

		head_dir += head_value + "/" + job_names[i] + "/"

		head_dir_open := head_dir + "open"

		head_dir_close := head_dir + "close"

		if _, err := os.Stat(head_dir_open); err != nil {

			head_dir_ignore := head_dir + "ignore"

			t_now := time.Now()

			t_str := t_now.Format("2006-01-02-15-04-05")

			_ = os.WriteFile(head_dir_ignore, []byte("ignore"), 0644)

			_ = os.WriteFile(head_dir_open, []byte(t_str), 0644)

			_ = os.WriteFile(head_dir_close, []byte(t_str), 0644)

			return build_dirs, fmt.Errorf(ERR_MSG, err.Error()+": forced termination")

		}

		if _, err := os.Stat(head_dir_close); err != nil {

			return build_dirs, fmt.Errorf(ERR_MSG, err.Error()+": build in progress")

		}

	}

	head_value_int += 1

	head_value = strconv.Itoa(head_value_int)

	_ = os.WriteFile(".usr/build/HEAD", []byte(head_value), 0644)

	for i := 0; i < len(job_names); i++ {

		new_head_dir = ".usr/build/"

		new_head_dir += head_value + "/" + job_names[i]

		cmd = exec.Command("mkdir", "-p", new_head_dir)

		err = cmd.Run()

		if err != nil {
			return build_dirs, fmt.Errorf(ERR_MSG, err.Error())
		}

		new_head_dir += "/"

		t_now := time.Now()

		t_str := t_now.Format("2006-01-02-15-04-05")

		new_head_dir_open := new_head_dir + "open"

		_ = os.WriteFile(new_head_dir_open, []byte(t_str), 0644)

		build_dirs = append(build_dirs, new_head_dir+"log")

	}

	return build_dirs, nil
}

func BuildCloseWithName(job_name string, close_msg string) error {

	ERR_MSG := "failed build close: %s"

	head_value := "0"

	head_dir := ".usr/build/"

	if _, err := os.Stat(".usr/build"); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error())

	}

	if _, err := os.Stat(".usr/build/HEAD"); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error())

	}

	head, err := os.ReadFile(".usr/build/HEAD")

	if err != nil {
		return fmt.Errorf(ERR_MSG, err.Error())
	}

	head_value = string(head)

	head_dir += head_value + "/" + job_name + "/"

	head_dir_open := head_dir + "open"

	head_dir_close := head_dir + "close"

	head_dir_result := head_dir + "result"

	if _, err := os.Stat(head_dir_open); err != nil {

		head_dir_ignore := head_dir + "ignore"

		t_now := time.Now()

		t_str := t_now.Format("2006-01-02-15-04-05")

		_ = os.WriteFile(head_dir_ignore, []byte("ignore"), 0644)

		_ = os.WriteFile(head_dir_open, []byte(t_str), 0644)

		_ = os.WriteFile(head_dir_close, []byte(t_str), 0644)

		return fmt.Errorf(ERR_MSG, err.Error()+": forced termination")

	}

	if _, err := os.Stat(head_dir_close); err == nil {

		return fmt.Errorf(ERR_MSG, "build already closed")

	}

	t_now := time.Now()

	t_str := t_now.Format("2006-01-02-15-04-05")

	_ = os.WriteFile(head_dir_close, []byte(t_str), 0644)

	_ = os.WriteFile(head_dir_result, []byte(close_msg), 0644)

	return nil
}

func CreateUsrPipelineVariables(stream_b []byte) error {

	err := os.WriteFile(".usr/pipeline_variables.yaml", stream_b, 0644)

	if err != nil {

		return fmt.Errorf("failed to create pipeline var: %s", err.Error())

	}

	return nil

}

func LoadUsrPipelineVariables() ([]byte, error) {

	var ret_b []byte

	ret_b, err := os.ReadFile(".usr/pipeline_variables.yaml")

	if err != nil {

		return ret_b, fmt.Errorf("failed to load pipeline var: %s", err.Error())

	}

	return ret_b, nil

}
