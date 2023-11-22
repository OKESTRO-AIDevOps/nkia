package runtimefs

import (
	"fmt"
	"os"
)

// control-plane
//  >
//   toolkit
//
// another-control-plane
//
// worker
//  >
//   volume
//    >
//     toolkit

func OpenFilePointerForNpiaInstallCtrlLog() (*os.File, error) {

	var outfile *os.File
	var err error

	install_log_path, err := InstallOpenForward("control-plane")

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	if _, err := os.Stat(install_log_path); err == nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", "another installation in process")
	}

	outfile, err = os.Create(install_log_path)

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	return outfile, nil
}

func OpenFilePointerForNpiaInstallAnotherCtrlLog() (*os.File, error) {

	var outfile *os.File
	var err error

	install_log_path, err := InstallOpenForward("another-control-plane")

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	if _, err := os.Stat(install_log_path); err == nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", "another installation in process")
	}

	outfile, err = os.Create(install_log_path)

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	return outfile, nil
}

func OpenFilePointerForNpiaInstallWorkerLog() (*os.File, error) {

	var outfile *os.File
	var err error

	install_log_path, err := InstallOpenForward("worker")

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	if _, err := os.Stat(install_log_path); err == nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", "another installation in process")
	}

	outfile, err = os.Create(install_log_path)

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	return outfile, nil
}

func OpenFilePointerForNpiaInstallVolumeLog() (*os.File, error) {

	var outfile *os.File
	var err error

	install_log_path, err := InstallOpenForward("volume")

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	if _, err := os.Stat(install_log_path); err == nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", "another installation in process")
	}

	outfile, err = os.Create(install_log_path)

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	return outfile, nil
}

func OpenFilePointerForNpiaInstallToolkitLog() (*os.File, error) {

	var outfile *os.File
	var err error

	install_log_path, err := InstallOpenForward("toolkit")

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	if _, err := os.Stat(install_log_path); err == nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", "another installation in process")
	}

	outfile, err = os.Create(install_log_path)

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	return outfile, nil
}

func CloseFilePointerForInstallLogAndMarkDone(fp *os.File) error {

	err := fp.Close()

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	err = InstallClose()

	if err != nil {
		return fmt.Errorf("failed to close installation: %s", err.Error())
	}

	return nil
}

func CloseFilePointerForInstallLogAndMarkFail(fp *os.File, close_msg string) error {

	err := fp.Close()

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	err = InstallCloseBackward(close_msg)

	if err != nil {
		return fmt.Errorf("failed to close installation: %s", err.Error())
	}

	return nil
}

func GetOngoingInstallLog() ([]byte, error) {

	var ret_byte []byte

	already_closed := 0

	head_dir := ".npia/install/"

	ERR_MSG := "failed to get ongoing install log: %s"

	if _, err := os.Stat(".npia/install/HEAD"); err != nil {

		return ret_byte, fmt.Errorf(ERR_MSG, err.Error())

	}

	head_value_b, err := os.ReadFile(".npia/install/HEAD")

	if err != nil {
		return ret_byte, fmt.Errorf(ERR_MSG, err.Error())
	}

	head_value := string(head_value_b)

	if head_value == "" {
		return ret_byte, fmt.Errorf(ERR_MSG, "empty HEAD, re-run previous command to retry initiation")
	}

	head_dir += head_value

	if _, err := os.Stat(head_dir); err != nil {

		return ret_byte, fmt.Errorf(ERR_MSG, err.Error())

	}

	head_dir += "/"

	head_dir_open := head_dir + "open"

	head_dir_close := head_dir + "close"

	head_dir_log := head_dir + "log"

	if _, err := os.Stat(head_dir_open); err != nil {

		return ret_byte, fmt.Errorf(ERR_MSG, err.Error())

	}

	if _, err := os.Stat(head_dir_close); err == nil {

		already_closed = 1

	}

	file_b, err := os.ReadFile(head_dir_log)

	if err != nil {
		return ret_byte, fmt.Errorf(ERR_MSG, err.Error())
	}

	if already_closed == 1 {
		file_b = append(file_b, []byte("\n-----------INSTALLATION CLOSED-----------\n")...)
	} else {
		file_b = append(file_b, []byte("\n-----------INSTALLATION ONGOING-----------\n")...)
	}

	ret_byte = file_b

	return ret_byte, nil

}

func OpenFilePointerForNpiaInstallRemoteLog() (*os.File, error) {

	var outfile *os.File
	var err error

	install_remote_log_path, err := InstallOpenForwardRemote()

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	if _, err := os.Stat(install_remote_log_path); err == nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", "another build in process")
	}

	outfile, err = os.Create(install_remote_log_path)

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	return outfile, nil
}

func CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp *os.File, close_msg string) error {

	err := fp.Close()

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	err = InstallCloseRemote(close_msg)

	if err != nil {
		return fmt.Errorf("failed to close install remote: %s", err.Error())
	}

	return nil
}
