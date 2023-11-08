package kubebase

import (
	"fmt"
	"os/exec"

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

func InstallControlPlane(localip string, osnm string, cversion string) ([]byte, error) {

	var ret_byte []byte

	ERR_MSG := "failed to install controle plane: %s"

	fp, err := runfs.OpenFilePointerForNpiaInstallCtrlLog()

	libif, err := libinterface.ConstructLibIface()

	if err != nil {
		return ret_byte, fmt.Errorf(ERR_MSG, err.Error())
	}

	os_release := utils.MakeOSReleaseLinux()

	lib_base_name, err := ConstructBaseName("ADMIN-INSTCTRL", os_release["ID"])

	if err != nil {
		return ret_byte, fmt.Errorf(ERR_MSG, err.Error())
	}

	LIB_BASE_INSTALL_CTRL, err := libif.GetLibComponentAddress("base", lib_base_name)

	if err != nil {
		return ret_byte, fmt.Errorf(ERR_MSG, err.Error())
	}

	cmd := exec.Command(LIB_BASE_INSTALL_CTRL, localip, osnm, cversion)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	return ret_byte, nil
}

//func InstallAnotherControlPlaneCertificate(targetip string, targetid string, targetpw string) ([]byte, error) {
//	var ret_byte []byte

//	return ret_byte, nil
//}

//func InstallAnotherControlPlaneOnLocal(targetip string, targetid string, targetpw string) ([]byte, error) {
//	var ret_byte []byte

//	return ret_byte, nil
//}

//func InstallAnotherControlPlaneOnRemote(targetip string, targetid string, targetpw string, localip string, osnm string, cv string, token string, nrole string, nid string) ([]byte, error) {
//	var ret_byte []byte

//	return ret_byte, nil
//}

func InstallWorkerOnLocal(localip string, osnm string, cv string, token string, nrole string, nid string) ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}

func InstallWorkerOnRemote(targetip string, targetid string, targetpw string, localip string, osnm string, cv string, token string, nrole string, nid string) ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}

func InstallVolumeOnLocal(localip string) ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}

func InstallVolumeOnRemote(targetip string, targetid string, targetpw string) ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}

func InstallToolKitOnLocal() ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}

func InstallToolKitOnRemote(targetip string, targetid string, targetpw string) ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}

func InstallLogOnLocal() ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}

func InstallLogOnRemote(targetip string, targetid string, targetpw string) ([]byte, error) {
	var ret_byte []byte

	return ret_byte, nil
}
