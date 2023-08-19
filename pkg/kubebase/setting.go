package kubebase

import (
	"fmt"
	"os/exec"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/libinterface"
	runfs "github.com/OKESTRO-AIDevOps/nkia/pkg/runtimefs"
)

func SettingCreateNamespace(main_ns string, repoaddr string, regaddr string) ([]byte, error) {

	var ret_byte []byte

	err := runfs.SetAdminOriginNewNS(main_ns, repoaddr, regaddr)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = []byte("record info registered\n")

	return ret_byte, nil

}

func SettingRepoInfo(main_ns string, repoaddr string, repoid string, repopw string) ([]byte, error) {

	var ret_byte []byte

	var app_origin runfs.AppOrigin

	app_origin, err := runfs.LoadAdmOrigin()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ns_found, _, rec_regaddr := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		return ret_byte, fmt.Errorf(": %s", "no such namespace")
	}

	app_origin.RECORDS = runfs.SetRecordInfo(app_origin.RECORDS, main_ns, repoaddr, rec_regaddr)

	app_origin.REPOS = runfs.SetRepoInfo(app_origin.REPOS, repoaddr, repoid, repopw)

	err = runfs.UnloadAdmOrigin(app_origin)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte = []byte("repo info registered\n")

	return ret_byte, nil
}

func SettingRegInfo(main_ns string, regaddr string, regid string, regpw string) ([]byte, error) {

	var ret_byte []byte

	var app_origin runfs.AppOrigin

	app_origin, err := runfs.LoadAdmOrigin()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ns_found, rec_repoaddr, _ := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		return ret_byte, fmt.Errorf(": %s", "no such namespace")
	}

	app_origin.RECORDS = runfs.SetRecordInfo(app_origin.RECORDS, main_ns, rec_repoaddr, regaddr)

	app_origin.REGS = runfs.SetRegInfo(app_origin.REGS, regaddr, regid, regpw)

	err = runfs.UnloadAdmOrigin(app_origin)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte = []byte("reg info registered\n")

	return ret_byte, nil
}

func SettingCreateMonitoring() ([]byte, error) {

	var ret_byte []byte

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	LIBIF_SCRIPTS_PROM_CREATE, err := libif.GetLibComponentAddress("scripts", "prom_create")

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	cmd := exec.Command(LIBIF_SCRIPTS_PROM_CREATE)

	_ = cmd.Run()

	cmd = exec.Command("kubectl", "get", "deployments")

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil
}
