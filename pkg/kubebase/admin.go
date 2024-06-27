package kubebase

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/libinterface"
	runfs "github.com/OKESTRO-AIDevOps/nkia/pkg/runtimefs"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/utils"
)

func AdminInitNPIA() {

	outfile, err := os.Create(".npia/.init_log")

	lib_exists := 0

	if err != nil {
		return
	}

	if _, err := os.Stat("lib"); err == nil {

		lib_exists = 1

		return

	}

	if _, err := os.Stat(".usr"); err == nil {

		outfile.Write([]byte("failed to init: .usr exists\n"))

		return

	}

	if _, err := os.Stat(".etc"); err == nil {

		outfile.Write([]byte("failed to init: .etc exists\n"))

		return
	}

	if lib_exists != 1 {

		cmd := exec.Command("curl", "-L", "https://github.com/OKESTRO-AIDevOps/nkia/releases/download/latest/lib.tgz", "-o", "lib.tgz")

		cmd.Stdout = outfile

		cmd.Stderr = outfile

		err = cmd.Run()

		if err != nil {
			AdminBlindResetNPIA()

			outfile.Write([]byte(err.Error()))

			return
		}

		cmd = exec.Command("tar", "-xzvf", "lib.tgz")

		cmd.Stdout = outfile

		cmd.Stderr = outfile

		err = cmd.Run()

		if err != nil {
			AdminBlindResetNPIA()
			outfile.Write([]byte(err.Error()))
			return
		}

		cmd = exec.Command("rm", "-r", "lib.tgz")

		cmd.Stdout = outfile

		cmd.Stderr = outfile

		err = cmd.Run()

		if err != nil {
			AdminBlindResetNPIA()
			outfile.Write([]byte(err.Error()))
			return
		}

	}

	cmd := exec.Command("mkdir", "-p", ".usr")

	cmd.Stdout = outfile

	cmd.Stderr = outfile

	err = cmd.Run()

	if err != nil {
		AdminBlindResetNPIA()
		outfile.Write([]byte(err.Error()))
		return
	}

	cmd = exec.Command("mkdir", "-p", ".etc")

	cmd.Stdout = outfile

	cmd.Stderr = outfile

	err = cmd.Run()

	if err != nil {
		AdminBlindResetNPIA()
		outfile.Write([]byte(err.Error()))
		return
	}

	err = runfs.CreateAdmOrigin()

	if err != nil {
		AdminBlindResetNPIA()
		outfile.Write([]byte(err.Error()))
		return
	}

	libif, err := libinterface.ConstructLibIface()

	if err != nil {
		AdminBlindResetNPIA()
		outfile.Write([]byte(err.Error()))
		return
	}

	os_release := utils.MakeOSReleaseLinux()

	lib_base_name, err := ConstructBaseName("ADMIN-INIT", os_release["ID"])

	if err != nil {
		outfile.Write([]byte("failed to init: " + err.Error() + "\n"))
		return
	}

	LIBIF_SCRIPTS_ADMIN_INIT_DEPENDENCY, err := libif.GetLibComponentAddress("base", lib_base_name)

	if err != nil {
		AdminBlindResetNPIA()
		outfile.Write([]byte(err.Error()))
		return
	}

	cmd = exec.Command(LIBIF_SCRIPTS_ADMIN_INIT_DEPENDENCY)

	cmd.Stdout = outfile

	cmd.Stderr = outfile

	err = cmd.Run()

	if err != nil {
		AdminBlindResetNPIA()
		outfile.Write([]byte(err.Error()))
		return
	}

	LIBIF_SCRIPTS_ADMIN_INIT_DEPENDENCY2, err := libif.GetLibComponentAddress("bin", lib_base_name)

	if err != nil {
		AdminBlindResetNPIA()
		outfile.Write([]byte(err.Error()))
		return
	}

	cmd = exec.Command(LIBIF_SCRIPTS_ADMIN_INIT_DEPENDENCY2)

	cmd.Stdout = outfile

	cmd.Stderr = outfile

	err = cmd.Run()

	if err != nil {
		AdminBlindResetNPIA()
		outfile.Write([]byte(err.Error()))
		return
	}

	outfile.Write([]byte("npia init success\n"))

	err = outfile.Close()

	if err != nil {
		AdminBlindResetNPIA()
		return

	}

	file_byte, err := os.ReadFile(".npia/.init_log")

	if err != nil {
		AdminBlindResetNPIA()
		return
	}

	err = os.Remove(".npia/.init_log")

	if err != nil {

		AdminBlindResetNPIA()
		return
	}

	err = os.WriteFile(".npia/.init", file_byte, 0644)

	if err != nil {
		AdminBlindResetNPIA()
		return
	}

	return

}

func AdminGetInitLog() ([]byte, error) {

	var ret_byte []byte

	var err error

	if _, err := os.Stat(".npia/.init_log"); err == nil {

		ret_byte, err = os.ReadFile(".npia/.init_log")

	} else if _, err := os.Stat(".npia/.init"); err == nil {

		ret_byte, err = os.ReadFile(".npia/.init")

	} else {

		return ret_byte, fmt.Errorf("failed to get init log: %s", "none found")

	}

	return ret_byte, err

}

func AdminBlindResetNPIA() {

	cmd := exec.Command("rm", "-r", "lib.tgz")

	_ = cmd.Run()

	cmd = exec.Command("rm", "-r", "lib")

	_ = cmd.Run()

	cmd = exec.Command("rm", "-r", ".usr")

	_ = cmd.Run()

	cmd = exec.Command("rm", "-r", ".etc")

	_ = cmd.Run()

}
