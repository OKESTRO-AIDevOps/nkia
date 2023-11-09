package kubebase

import (
	"os/exec"
	"strings"

	runfs "github.com/OKESTRO-AIDevOps/nkia/pkg/runtimefs"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/libinterface"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/utils"
)

//func InstallEnvironment(localip string) ([]byte, error) {
//	var ret_byte []byte

//	return ret_byte, nil
//}

//func InstallEnvironmentRestart(localip string) ([]byte, error) {
//	var ret_byte []byte

//	return ret_byte, nil
//}

func InstallControlPlane(localip string, osnm string, cv string) {

	fp, err := runfs.OpenFilePointerForNpiaInstallCtrlLog()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	os_release := utils.MakeOSReleaseLinux()

	lib_base_name, err := ConstructBaseName("ADMIN-INSTCTRL", os_release["ID"])

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	LIB_BASE_INSTALL_CTRL, err := libif.GetLibComponentAddress("base", lib_base_name)

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	cmd := exec.Command(LIB_BASE_INSTALL_CTRL, localip, osnm, cv)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)

	}

	cmd = exec.Command("kubectl", "get", "nodes")

	_, err = cmd.Output()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
	}

	err = runfs.CloseFilePointerForInstallLogAndMarkDone(fp)

	if err != nil {
		return
	}

	return
}

//func InstallAnotherControlPlaneCertificate(targetip string, targetid string, targetpw string) ([]byte, error) {
//	var ret_byte []byte

//	return ret_byte, nil
//}

//func InstallAnotherControlPlaneOnLocal(localip string, osnm string, cv string, token string) ([]byte, error) {
//	var ret_byte []byte

//	return ret_byte, nil
//}

//func InstallAnotherControlPlaneOnRemote(targetip string, targetid string, targetpw string, localip string, osnm string, cv string, token string) ([]byte, error) {
//	var ret_byte []byte

//	return ret_byte, nil
//}

func InstallWorkerOnLocal(localip string, osnm string, cv string, token string) {

	ip_no_dot := strings.ReplaceAll(localip, ".", "-")

	hex_random, _ := utils.RandomHex(4)

	nid := "wk-" + osnm + "-" + ip_no_dot + "-" + hex_random

	fp, err := runfs.OpenFilePointerForNpiaInstallWorkerLog()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	os_release := utils.MakeOSReleaseLinux()

	lib_base_name, err := ConstructBaseName("ADMIN-INSTWKOL", os_release["ID"])

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	LIB_BASE_INSTALL_WK_OL, err := libif.GetLibComponentAddress("base", lib_base_name)

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	cmd := exec.Command(LIB_BASE_INSTALL_WK_OL, localip, osnm, cv)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return

	}

	token += " " + "--node-name" + " " + nid

	cmd_args := strings.Fields(token)

	cmd = exec.Command(cmd_args[0], cmd_args[1:]...)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return

	}

	cmd = exec.Command("cat", "/var/lib/kubelet/kubeadm-flags.env")

	stdout, _ := cmd.Output()

	str_kubeenv := string(stdout)

	if !strings.Contains(str_kubeenv, hex_random) {
		close_msg := "failed installing worker node: kubelet flags checking failed\n"

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	err = runfs.CloseFilePointerForInstallLogAndMarkDone(fp)

	if err != nil {
		return
	}

	return
}

func InstallWorkerOnRemote(targetip string, targetid string, targetpw string, localip string, osnm string, cv string, token string) ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}

func InstallVolumeOnLocal(localip string) {

	fp, err := runfs.OpenFilePointerForNpiaInstallVolumeLog()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	os_release := utils.MakeOSReleaseLinux()

	lib_base_name, err := ConstructBaseName("ADMIN-INSTVOLOL", os_release["ID"])

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	LIB_BASE_INSTALL_VOL_OL, err := libif.GetLibComponentAddress("base", lib_base_name)

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	cmd := exec.Command(LIB_BASE_INSTALL_VOL_OL, localip)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return

	}

	cmd = exec.Command("exportfs", "-v")

	stdout, err := cmd.Output()

	printout := string(stdout)

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return

	}

	if printout == "" {
		close_msg := "failed to install volume: nothing exported\n"

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return

	}

	err = runfs.CloseFilePointerForInstallLogAndMarkDone(fp)

	if err != nil {
		return
	}

	return
}

//func InstallVolumeOnRemote(targetip string, targetid string, targetpw string) ([]byte, error) {
//	var ret_byte []byte

//	return ret_byte, nil
//}

func InstallToolKitOnLocal() {

	fp, err := runfs.OpenFilePointerForNpiaInstallToolkitLog()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	os_release := utils.MakeOSReleaseLinux()

	lib_base_name, err := ConstructBaseName("ADMIN-INSTTKOL", os_release["ID"])

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	LIB_BASE_INSTALL_TK_OL, err := libif.GetLibComponentAddress("base", lib_base_name)

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return
	}

	cmd := exec.Command(LIB_BASE_INSTALL_TK_OL)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return

	}

	err = runfs.CloseFilePointerForInstallLogAndMarkDone(fp)

	if err != nil {
		return
	}

	return
}

//func InstallToolKitOnRemote(targetip string, targetid string, targetpw string) ([]byte, error) {
//	var ret_byte []byte

//	return ret_byte, nil
//}

func InstallLogOnLocal() ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}

func InstallLogOnRemote(targetip string, targetid string, targetpw string) ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}
