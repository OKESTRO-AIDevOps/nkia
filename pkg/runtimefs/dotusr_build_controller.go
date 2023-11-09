package runtimefs

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func BuildOpenForward() (string, error) {

	var build_dir string

	ERR_MSG := "failed build open forward: %s"

	head_value := "0"

	head_value_int := 0

	head_dir := ".usr/build/"

	new_head_dir := ".usr/build/"

	cmd := exec.Command("mkdir", "-p", ".usr/build")

	_, err := cmd.Output()

	if err != nil {
		return build_dir, fmt.Errorf(ERR_MSG, err.Error())
	}

	if _, err := os.Stat(".usr/build/HEAD"); err != nil {

		cmd := exec.Command("mkdir", "-p", ".usr/build/0")

		_, _ = cmd.Output()

		t_now := time.Now()

		t_str := t_now.Format("2006-01-02-15-04-05")

		_ = os.WriteFile(".usr/build/0/open", []byte(t_str), 0644)

		_ = os.WriteFile(".usr/build/0/close", []byte(t_str), 0644)

		_ = os.WriteFile(".usr/build/HEAD", []byte("0"), 0644)
	}

	head, err := os.ReadFile(".usr/build/HEAD")

	if err != nil {
		return build_dir, fmt.Errorf(ERR_MSG, err.Error())
	}

	head_value = string(head)

	head_value_int, err = strconv.Atoi(head_value)

	if err != nil {
		return build_dir, fmt.Errorf(ERR_MSG, err.Error())
	}

	head_dir += head_value + "/"

	head_dir_open := head_dir + "open"

	head_dir_close := head_dir + "close"

	if _, err := os.Stat(head_dir_open); err != nil {

		head_dir_ignore := head_dir + "ignore"

		t_now := time.Now()

		t_str := t_now.Format("2006-01-02-15-04-05")

		_ = os.WriteFile(head_dir_ignore, []byte("ignore"), 0644)

		_ = os.WriteFile(head_dir_open, []byte(t_str), 0644)

		_ = os.WriteFile(head_dir_close, []byte(t_str), 0644)

		return build_dir, fmt.Errorf(ERR_MSG, err.Error()+": forced termination")

	}

	if _, err := os.Stat(head_dir_close); err != nil {

		return build_dir, fmt.Errorf(ERR_MSG, err.Error()+": build in progress")

	}

	head_value_int += 1

	head_value = strconv.Itoa(head_value_int)

	_ = os.WriteFile(".usr/build/HEAD", []byte(head_value), 0644)

	new_head_dir += head_value

	cmd = exec.Command("mkdir", "-p", new_head_dir)

	err = cmd.Run()

	if err != nil {
		return build_dir, fmt.Errorf(ERR_MSG, err.Error())
	}

	new_head_dir += "/"

	t_now := time.Now()

	t_str := t_now.Format("2006-01-02-15-04-05")

	new_head_dir_open := new_head_dir + "open"

	_ = os.WriteFile(new_head_dir_open, []byte(t_str), 0644)

	build_dir = new_head_dir + "log"

	return build_dir, nil
}

func BuildClose(close_msg string) error {

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

	head_dir += head_value + "/"

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
