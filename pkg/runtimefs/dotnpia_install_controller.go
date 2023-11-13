package runtimefs

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// control-plane
//
// another-control-plane
//
// worker
//
// (volume)
//
// (toolkit)

func InstallOpenForward(head_value string) (string, error) {

	var install_path string

	ERR_MSG := "failed install open: %s"

	head_dir := ".npia/install/"

	new_head_dir := ".npia/install/"

	head_file_value := ""

	equal_heads := 0

	cmd := exec.Command("mkdir", "-p", ".npia/install")

	_, err := cmd.Output()

	if err != nil {
		return install_path, fmt.Errorf(ERR_MSG, err.Error())
	}

	if _, err := os.Stat(".npia/install/HEAD"); err == nil && head_value != "volume" && head_value != "toolkit" {

		head_file_b, err := os.ReadFile(".npia/install/HEAD")

		if err != nil {
			return install_path, fmt.Errorf(ERR_MSG, err.Error())
		}

		head_file_value = string(head_file_b)

		if head_file_value == "control-plane" {

			if head_value != head_file_value {

				return install_path, fmt.Errorf(ERR_MSG, "already a control plane")

			} else {
				equal_heads = 1
			}

		} else if head_file_value == "another-control-plane" {

			if head_value != head_file_value {

				return install_path, fmt.Errorf(ERR_MSG, "already an another control plane")

			} else {
				equal_heads = 1
			}

		} else if head_file_value == "worker" {

			if head_value != head_file_value {

				return install_path, fmt.Errorf(ERR_MSG, "already a worker")

			} else {
				equal_heads = 1
			}

		}

	}

	if equal_heads == 1 {

		head_dir += head_value + "/"

		head_dir_open := head_dir + "open"

		head_dir_close := head_dir + "close"

		if _, err := os.Stat(head_dir_open); err == nil {

			if _, err := os.Stat(head_dir_close); err != nil {

				return install_path, fmt.Errorf(ERR_MSG, "another installation in process for: "+head_value)

			} else {

				return install_path, fmt.Errorf(ERR_MSG, "installation already completed for: "+head_value)
			}

		} else {

			return install_path, fmt.Errorf(ERR_MSG, "possibly a broken installation: no open file")

		}

	}

	cmd = exec.Command("rm", ".npia/install/TAIL")

	_ = cmd.Run()

	if _, err := os.Stat(".npia/install/HEAD"); err != nil {

		_ = os.WriteFile(".npia/install/TAIL", []byte(""), 0644)

	} else {

		cmd = exec.Command("cp", ".npia/install/HEAD", ".npia/install/TAIL")

		_ = cmd.Run()

	}

	_ = os.WriteFile(".npia/install/HEAD", []byte(head_value), 0644)

	t_now := time.Now()

	t_str := t_now.Format("2006-01-02-15-04-05")

	new_head_dir += head_value

	if _, err := os.Stat(new_head_dir); err != nil {

		cmd := exec.Command("mv", new_head_dir, new_head_dir+"."+t_str)

		_ = cmd.Run()

	}

	cmd = exec.Command("mkdir", "-p", new_head_dir)

	err = cmd.Run()

	if err != nil {
		return install_path, fmt.Errorf(ERR_MSG, err.Error())
	}

	new_head_dir += "/"

	new_head_dir_open := new_head_dir + "open"

	_ = os.WriteFile(new_head_dir_open, []byte(t_str), 0644)

	install_path = new_head_dir + "log"

	return install_path, nil
}

func InstallClose() error {

	ERR_MSG := "failed build close: %s"

	head_value := "0"

	head_dir := ".npia/install/"

	if _, err := os.Stat(".npia/install"); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error())

	}

	if _, err := os.Stat(".npia/install/HEAD"); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error())

	}

	head, err := os.ReadFile(".npia/install/HEAD")

	if err != nil {
		return fmt.Errorf(ERR_MSG, err.Error())
	}

	head_value = string(head)

	head_dir += head_value + "/"

	head_dir_open := head_dir + "open"

	head_dir_close := head_dir + "close"

	head_dir_result := head_dir + "result"

	if _, err := os.Stat(head_dir_open); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error()+": possibly a broken installation")

	}

	if _, err := os.Stat(head_dir_close); err == nil {

		return fmt.Errorf(ERR_MSG, "installation already closed")

	}

	t_now := time.Now()

	t_str := t_now.Format("2006-01-02-15-04-05")

	_ = os.WriteFile(head_dir_close, []byte(t_str), 0644)

	_ = os.WriteFile(head_dir_result, []byte("SUCCESS"), 0644)

	if head_value == "control-plane" || head_value == "another-control-plane" || head_value == "worker" {

		_ = os.WriteFile(".npia/install/BODY", []byte(head_value), 0644)

	}

	return nil
}

func InstallCloseBackward(err_msg string) error {

	ERR_MSG := "failed build close backward: %s"

	head_value := "0"

	tail_value := "0"

	head_dir := ".npia/install/"

	new_head_dir := ".npia/install/"

	if _, err := os.Stat(".npia/install"); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error())

	}

	if _, err := os.Stat(".npia/install/HEAD"); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error())

	}

	if _, err := os.Stat(".npia/install/TAIL"); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error())

	}

	tail, err := os.ReadFile(".npia/install/TAIL")

	if err != nil {
		return fmt.Errorf(ERR_MSG, err.Error())
	}

	tail_value = string(tail)

	head, err := os.ReadFile(".npia/install/HEAD")

	if err != nil {
		return fmt.Errorf(ERR_MSG, err.Error())
	}

	head_value = string(head)

	head_dir += head_value + "/"

	head_dir_open := head_dir + "open"

	head_dir_close := head_dir + "close"

	head_dir_result := head_dir + "result"

	if _, err := os.Stat(head_dir_open); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error()+": possibly a broken installation")

	}

	if _, err := os.Stat(head_dir_close); err == nil {

		return fmt.Errorf(ERR_MSG, "installation already closed")

	}

	t_now := time.Now()

	t_str := t_now.Format("2006-01-02-15-04-05")

	new_head_dir += head_value

	_ = os.WriteFile(head_dir_close, []byte(t_str), 0644)

	_ = os.WriteFile(head_dir_result, []byte(err_msg), 0644)

	cmd := exec.Command("mv", new_head_dir, new_head_dir+"."+t_str)

	_ = cmd.Run()

	_ = os.WriteFile(".npia/install/HEAD", []byte(tail_value), 0644)

	return nil
}

func InstallOpenForwardRemote() (string, error) {

	var build_dir string

	ERR_MSG := "failed build open forward: %s"

	head_value := "0"

	head_value_int := 0

	head_dir := ".npia/install/remote/"

	new_head_dir := ".npia/install/remote/"

	cmd := exec.Command("mkdir", "-p", ".npia/install/remote")

	_, err := cmd.Output()

	if err != nil {
		return build_dir, fmt.Errorf(ERR_MSG, err.Error())
	}

	if _, err := os.Stat(".npia/install/remote/HEAD"); err != nil {

		cmd := exec.Command("mkdir", "-p", ".npia/install/remote/0")

		_, _ = cmd.Output()

		t_now := time.Now()

		t_str := t_now.Format("2006-01-02-15-04-05")

		_ = os.WriteFile(".npia/install/remote/0/open", []byte(t_str), 0644)

		_ = os.WriteFile(".npia/install/remote/0/close", []byte(t_str), 0644)

		_ = os.WriteFile(".npia/install/remote/HEAD", []byte("0"), 0644)
	}

	head, err := os.ReadFile(".npia/install/remote/HEAD")

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

		return build_dir, fmt.Errorf(ERR_MSG, err.Error()+": installation in progress")

	}

	head_value_int += 1

	head_value = strconv.Itoa(head_value_int)

	_ = os.WriteFile(".npia/install/remote/HEAD", []byte(head_value), 0644)

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

func InstallCloseRemote(close_msg string) error {

	ERR_MSG := "failed build close: %s"

	head_value := "0"

	head_dir := ".npia/install/remote/"

	if _, err := os.Stat(".npia/install/remote"); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error())

	}

	if _, err := os.Stat(".npia/install/remote/HEAD"); err != nil {

		return fmt.Errorf(ERR_MSG, err.Error())

	}

	head, err := os.ReadFile(".npia/install/remote/HEAD")

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

		return fmt.Errorf(ERR_MSG, "installation already closed")

	}

	t_now := time.Now()

	t_str := t_now.Format("2006-01-02-15-04-05")

	_ = os.WriteFile(head_dir_close, []byte(t_str), 0644)

	_ = os.WriteFile(head_dir_result, []byte(close_msg), 0644)

	return nil
}
