package kubebase

import (
	"fmt"
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

func InstallControlPlane(localip string, osnm string, cv string) error {

	ERR_MSG := "failed installing control plane: %s\n"

	fp, err := runfs.OpenFilePointerForNpiaInstallCtrlLog()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	os_release := utils.MakeOSReleaseLinux()

	lib_base_name, err := ConstructBaseName("ADMIN-INSTCTRL", os_release["ID"])

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	LIB_BASE_INSTALL_CTRL, err := libif.GetLibComponentAddress("base", lib_base_name)

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	cmd := exec.Command(LIB_BASE_INSTALL_CTRL, localip, osnm, cv)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	cmd = exec.Command("kubectl", "get", "nodes")

	_, err = cmd.Output()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	err = runfs.CloseFilePointerForInstallLogAndMarkDone(fp)

	if err != nil {
		return fmt.Errorf(ERR_MSG, err.Error())
	}

	return nil
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

func InstallWorkerOnLocal(localip string, osnm string, cv string, token string) error {

	ERR_MSG := "failed installing worker: %s\n"

	ip_no_dot := strings.ReplaceAll(localip, ".", "-")

	hex_random, _ := utils.RandomHex(4)

	nid := "wk-" + osnm + "-" + ip_no_dot + "-" + hex_random

	fp, err := runfs.OpenFilePointerForNpiaInstallWorkerLog()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	os_release := utils.MakeOSReleaseLinux()

	lib_base_name, err := ConstructBaseName("ADMIN-INSTWKOL", os_release["ID"])

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	LIB_BASE_INSTALL_WK_OL, err := libif.GetLibComponentAddress("base", lib_base_name)

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	cmd := exec.Command(LIB_BASE_INSTALL_WK_OL, localip, osnm, cv)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)

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
		return fmt.Errorf(ERR_MSG, close_msg)

	}

	cmd = exec.Command("cat", "/var/lib/kubelet/kubeadm-flags.env")

	stdout, _ := cmd.Output()

	str_kubeenv := string(stdout)

	if !strings.Contains(str_kubeenv, hex_random) {
		close_msg := "failed installing worker node: kubelet flags checking failed\n"

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	err = runfs.CloseFilePointerForInstallLogAndMarkDone(fp)

	if err != nil {
		return fmt.Errorf(ERR_MSG, err.Error())
	}

	return nil
}

func InstallWorkerOnRemote(targetip string, targetid string, targetpw string, localip string, osnm string, cv string, token string) {

	fp, err := runfs.OpenFilePointerForNpiaInstallRemoteLog()

	conn, err := utils.ShellConnect(targetip, targetid, targetpw)

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))

		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		return
	}

	output, err := conn.SendCommands("sudo mkdir -p /npia && ls -la /npia")
	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))

		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

	}

	fp.Write(output)

	fp.Write([]byte("\n----------ROOT NPIA CREATED----------\n"))

	output, err = conn.SendCommands("sudo curl -L https://github.com/OKESTRO-AIDevOps/nkia/releases/download/latest/bin.tgz -o /npia/bin.tgz")
	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))

		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

	}

	fp.Write(output)

	fp.Write([]byte("\n----------NPIA BIN DOWNLOADED----------\n"))

	output, err = conn.SendCommands("sudo tar -xzf /npia/bin.tgz -C /npia")
	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))

		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

	}

	fp.Write(output)

	fp.Write([]byte("\n----------NPIA BIN INSTALLED----------\n"))

	output, err = conn.SendCommands("cd /npia/bin/nokubeadm && sudo ./nokubeadm init-npia-default")
	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))

		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

	}

	fp.Write(output)

	fp.Write([]byte("\n----------NPIA INITIATED----------\n"))

	options := " " + "--localip " + localip + " " + "--osnm " + osnm + " " + "--cv " + cv + " " + "--token " + "'" + token + "'"

	output, err = conn.SendCommands("cd /npia/bin/nokubeadm && sudo ./nokubeadm install worker" + options)
	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))

		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

	}

	fp.Write(output)

	fp.Write([]byte("\n----------WORKER INSTALLED----------\n"))

	runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, "SUCCESS")

	return
}

func InstallVolumeOnLocal(localip string) error {

	ERR_MSG := "failed volume: %s\n"

	fp, err := runfs.OpenFilePointerForNpiaInstallVolumeLog()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	os_release := utils.MakeOSReleaseLinux()

	lib_base_name, err := ConstructBaseName("ADMIN-INSTVOLOL", os_release["ID"])

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	LIB_BASE_INSTALL_VOL_OL, err := libif.GetLibComponentAddress("base", lib_base_name)

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	cmd := exec.Command(LIB_BASE_INSTALL_VOL_OL, localip)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)

	}

	cmd = exec.Command("exportfs", "-v")

	stdout, err := cmd.Output()

	printout := string(stdout)

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)

	}

	if printout == "" {
		close_msg := "failed to install volume: nothing exported\n"

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)

	}

	err = runfs.CloseFilePointerForInstallLogAndMarkDone(fp)

	if err != nil {
		return fmt.Errorf(ERR_MSG, err.Error())
	}

	return nil
}

func InstallVolumeOnRemote(targetip string, targetid string, targetpw string, localip string) {

	fp, err := runfs.OpenFilePointerForNpiaInstallRemoteLog()

	plugins, err := runfs.MakePluginJSON()

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)
	}

	kv := map[string]string{
		"targetip": targetip,
		"targetid": targetid,
		"targetpw": targetpw,
	}

	pj := runfs.PluginJSON{
		Name:   "volume",
		Keyval: kv,
	}

	conn, err := utils.ShellConnect(targetip, targetid, targetpw)

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))

		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		return
	}

	_, err = conn.SendCommands("ls -la /npia")

	if err != nil {
		output, err := conn.SendCommands("sudo mkdir -p /npia && ls -la /npia")
		if err != nil {
			close_msg := err.Error()

			fp.Write([]byte(close_msg))

			runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		}

		fp.Write(output)

		fp.Write([]byte("\n----------ROOT NPIA CREATED----------\n"))

		output, err = conn.SendCommands("sudo curl -L https://github.com/OKESTRO-AIDevOps/nkia/releases/download/latest/bin.tgz -o /npia/bin.tgz")
		if err != nil {
			close_msg := err.Error()

			fp.Write([]byte(close_msg))

			runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		}

		fp.Write(output)

		fp.Write([]byte("\n----------NPIA BIN DOWNLOADED----------\n"))

		output, err = conn.SendCommands("sudo tar -xzf /npia/bin.tgz -C /npia")
		if err != nil {
			close_msg := err.Error()

			fp.Write([]byte(close_msg))

			runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		}

		fp.Write(output)

		fp.Write([]byte("\n----------NPIA BIN INSTALLED----------\n"))

		output, err = conn.SendCommands("cd /npia/bin/nokubeadm && sudo ./nokubeadm init-npia-default")
		if err != nil {
			close_msg := err.Error()

			fp.Write([]byte(close_msg))

			runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		}

		fp.Write(output)

		fp.Write([]byte("\n----------NPIA INITIATED----------\n"))

	}

	options := " " + "--localip " + localip

	output, err := conn.SendCommands("cd /npia/bin/nokubeadm && sudo ./nokubeadm install volume" + options)
	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))

		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

	}

	fp.Write(output)

	fp.Write([]byte("\n----------VOLUME INSTALLED----------\n"))

	runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, "SUCCESS")

	plugins, _ = runfs.AddMapToPluginJSON(plugins, pj)

	_ = runfs.SavePluginJSON(plugins)

	return
}

func InstallToolKitOnLocal() error {

	ERR_MSG := "failed toolkit: %s\n"

	fp, err := runfs.OpenFilePointerForNpiaInstallToolkitLog()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	os_release := utils.MakeOSReleaseLinux()

	lib_base_name, err := ConstructBaseName("ADMIN-INSTTKOL", os_release["ID"])

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	LIB_BASE_INSTALL_TK_OL, err := libif.GetLibComponentAddress("base", lib_base_name)

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)
	}

	cmd := exec.Command(LIB_BASE_INSTALL_TK_OL)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForInstallLogAndMarkFail(fp, close_msg)
		return fmt.Errorf(ERR_MSG, close_msg)

	}

	err = runfs.CloseFilePointerForInstallLogAndMarkDone(fp)

	if err != nil {
		return fmt.Errorf(ERR_MSG, err.Error())
	}

	return nil
}

func InstallToolKitOnRemote(targetip string, targetid string, targetpw string) {

	fp, err := runfs.OpenFilePointerForNpiaInstallRemoteLog()

	plugins, err := runfs.MakePluginJSON()

	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))
		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)
	}

	kv := map[string]string{
		"targetip": targetip,
		"targetid": targetid,
		"targetpw": targetpw,
	}

	pj := runfs.PluginJSON{
		Name:   "toolkit",
		Keyval: kv,
	}

	conn, err := utils.ShellConnect(targetip, targetid, targetpw)

	if err != nil {

		close_msg := err.Error()

		fp.Write([]byte(close_msg))

		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		return
	}

	_, err = conn.SendCommands("ls -la /npia")

	if err != nil {
		output, err := conn.SendCommands("sudo mkdir -p /npia && ls -la /npia")
		if err != nil {
			close_msg := err.Error()

			fp.Write([]byte(close_msg))

			runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		}

		fp.Write(output)

		fp.Write([]byte("\n----------ROOT NPIA CREATED----------\n"))

		output, err = conn.SendCommands("sudo curl -L https://github.com/OKESTRO-AIDevOps/nkia/releases/download/latest/bin.tgz -o /npia/bin.tgz")
		if err != nil {
			close_msg := err.Error()

			fp.Write([]byte(close_msg))

			runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		}

		fp.Write(output)

		fp.Write([]byte("\n----------NPIA BIN DOWNLOADED----------\n"))

		output, err = conn.SendCommands("sudo tar -xzf /npia/bin.tgz -C /npia")
		if err != nil {
			close_msg := err.Error()

			fp.Write([]byte(close_msg))

			runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		}

		fp.Write(output)

		fp.Write([]byte("\n----------NPIA BIN INSTALLED----------\n"))

		output, err = conn.SendCommands("cd /npia/bin/nokubeadm && sudo ./nokubeadm init-npia-default")
		if err != nil {
			close_msg := err.Error()

			fp.Write([]byte(close_msg))

			runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

		}

		fp.Write(output)

		fp.Write([]byte("\n----------NPIA INITIATED----------\n"))

	}

	output, err := conn.SendCommands("cd /npia/bin/nokubeadm && sudo ./nokubeadm install toolkit")
	if err != nil {
		close_msg := err.Error()

		fp.Write([]byte(close_msg))

		runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, close_msg)

	}

	fp.Write(output)

	fp.Write([]byte("\n----------TOOLKIT INSTALLED----------\n"))

	runfs.CloseFilePointerForNpiaInstallRemoteLogAndMarkDone(fp, "SUCCESS")

	plugins, _ = runfs.AddMapToPluginJSON(plugins, pj)

	_ = runfs.SavePluginJSON(plugins)

	return
}

func InstallLogOnLocal() ([]byte, error) {

	var ret_byte []byte

	ret_byte, err := runfs.GetOngoingInstallLog()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	return ret_byte, nil
}

func InstallLogOnRemote(targetip string, targetid string, targetpw string) ([]byte, error) {

	var ret_byte []byte

	conn, err := utils.ShellConnect(targetip, targetid, targetpw)

	if err != nil {

		return ret_byte, fmt.Errorf("failed to get remote log: %s", err.Error())
	}

	output, err := conn.SendCommands("cd /npia/bin/nokubeadm && sudo ./nokubeadm install log")

	if err != nil {

		return ret_byte, fmt.Errorf("failed to get remote log: %s", err.Error())

	}

	ret_byte = output

	return ret_byte, nil
}
