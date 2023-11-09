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
