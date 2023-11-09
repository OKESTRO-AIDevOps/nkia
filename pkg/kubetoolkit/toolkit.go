package kubetoolkit

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/libinterface"
	runfs "github.com/OKESTRO-AIDevOps/nkia/pkg/runtimefs"
)

func ToolkitBuildImagesStart(main_ns string, repoaddr string, regaddr string) {

	fp, err := runfs.OpenFilePointerForUsrBuildLog()

	if err != nil {
		return
	}

	libif, err := libinterface.ConstructLibIface()

	LIBIF_BIN_DOCKER_COMPOSE, err := libif.GetLibComponentAddress("bin", "docker-compose")

	app_origin, err := runfs.LoadAdmOrigin()

	ns_found, _, _ := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		close_msg := "namespace not found in ADMorigin\n"
		fp.Write([]byte(close_msg))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, close_msg)
		return
	}

	err = runfs.InitUsrTarget(repoaddr)

	if err != nil {
		fp.Write([]byte(err.Error()))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	USR_TARGET_DOCKER_COMPOSE_YAML_BUILD, err := runfs.GetUsrTargetDockerComposeYamlBuild()

	if err != nil {
		fp.Write([]byte(err.Error()))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	cmd := exec.Command(LIBIF_BIN_DOCKER_COMPOSE, "-f", USR_TARGET_DOCKER_COMPOSE_YAML_BUILD, "up", "-d", "--build")

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {
		fp.Write([]byte(err.Error()))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	cmd = exec.Command(LIBIF_BIN_DOCKER_COMPOSE, "-f", USR_TARGET_DOCKER_COMPOSE_YAML_BUILD, "down")

	err = cmd.Run()

	if err != nil {
		fp.Write([]byte(err.Error()))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	addr_found, regid, regpw := runfs.GetRegInfo(app_origin.REGS, regaddr)

	if !addr_found {
		close_msg := "reg addr not found in ADMorigin\n"
		fp.Write([]byte(close_msg))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, close_msg)
		return
	}

	err = ToolkitBuildImagesStart_Push(fp, regaddr, regid, regpw)

	if err != nil {
		fp.Write([]byte(err.Error()))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	err = runfs.ClearUsrTarget()

	if err != nil {
		fp.Write([]byte(err.Error()))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	err = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, "SUCCESS")

	if err != nil {
		return
	}

	return

}

func ToolkitBuildImagesStart_Push(fp *os.File, regaddr string, regid string, regpw string) error {

	if fp == nil {
		return fmt.Errorf("push failed: %s", "no file pointer to write to")
	}

	USR_TARGET_PUSH_LIST, err := runfs.GetUsrTargetPushList(regaddr)

	if err != nil {
		return fmt.Errorf("push failed: %s", err.Error())
	}

	reg_url_auth := strings.SplitN(regaddr, "/", 2)[0]

	cmd := exec.Command("docker", "logout")

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("push failed: %s", err.Error())
	}

	cmd = exec.Command("docker", "login", reg_url_auth, "-u", regid, "-p", regpw)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("push failed: %s", err.Error())
	}

	for _, SOURCE_DEST := range USR_TARGET_PUSH_LIST {

		source := SOURCE_DEST[0]

		destination := SOURCE_DEST[1]

		cmd := exec.Command("docker", "tag", source, destination)

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("push failed: %s", err.Error())
		}

		cmd = exec.Command("docker", "push", destination)

		cmd.Stdout = fp

		cmd.Stderr = fp

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("push failed: %s", err.Error())
		}

	}

	return nil

}

func ToolkitBuildImagesGetLog() ([]byte, error) {

	var ret_byte []byte

	log_b, err := runfs.GetUsrBuildLog()

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get log: %s", err.Error())
	}

	ret_byte = log_b

	return ret_byte, nil
}
