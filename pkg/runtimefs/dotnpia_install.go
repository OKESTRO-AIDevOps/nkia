package runtimefs

import (
	"fmt"
	"os"
)

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

func CloseFilePointerForInstallCtrlLogAndMarkDone(fp *os.File) error {

	err := fp.Close()

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	file_byte, err := os.ReadFile(".usr/build_log")

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	err = os.Remove(".usr/build_log")

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	err = os.WriteFile(".usr/build_done", file_byte, 0644)

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	return nil
}
