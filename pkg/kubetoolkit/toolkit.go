package kubetoolkit

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	bsource "github.com/OKESTRO-AIDevOps/nkia/pkg/builtinresource"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/libinterface"
	runfs "github.com/OKESTRO-AIDevOps/nkia/pkg/runtimefs"
	goya "github.com/goccy/go-yaml"
)

func ToolkitBuildImagesStart2(main_ns string, repoaddr string, regaddr string) {

	bid, err := runfs.SetBuildId()

	if err != nil {
		return
	}

	fp, err := runfs.OpenFilePointerForUsrBuildLog()

	if err != nil {
		return
	}

	app_origin, err := runfs.LoadAdmOrigin()

	ns_found, _, _ := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		close_msg := "namespace not found in ADMorigin\n"
		fp.Write([]byte(close_msg))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, close_msg)
		return
	}

	repo_found, repo_id, repo_pw := runfs.GetRepoInfo(app_origin.REPOS, repoaddr)

	if !repo_found {

		close_msg := "repoaddr not found in ADMorigin\n"
		fp.Write([]byte(close_msg))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, close_msg)
		return
	}

	reg_found, reg_id, reg_pw := runfs.GetRegInfo(app_origin.REGS, regaddr)

	if !reg_found {

		close_msg := "regaddr not found in ADMorigin\n"
		fp.Write([]byte(close_msg))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, close_msg)
		return
	}

	v_repoaddr := strings.ReplaceAll(repoaddr, "https://", "")
	v_repoaddr = strings.ReplaceAll(v_repoaddr, "http://", "")
	v_repoaddr = "git://" + v_repoaddr

	v_regaddr := strings.ReplaceAll(regaddr, "https://", "")
	v_regaddr = strings.ReplaceAll(v_regaddr, "http://", "")

	ns_bid := "npia-build-ns-" + bid
	sec_bid := "npia-build-secret-" + bid
	pod_bid := "npia-build-pod-" + bid

	cmd := exec.Command("kubectl", "create", "namespace", ns_bid)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {
		fp.Write([]byte(err.Error()))
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	docker_server := "--docker-server=" + v_regaddr
	docker_uname := "--docker-username=" + reg_id
	docker_pword := "--docker-password=" + reg_pw

	cmd = exec.Command("kubectl", "-n", ns_bid, "create", "secret", "docker-registry", sec_bid, docker_server, docker_uname, docker_pword)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {
		fp.Write([]byte(err.Error()))
		ToolkitBuildCancelForce(ns_bid)
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	kb := bsource.KanikoBuilder{}

	kb_c := bsource.KanikoBuilder_Container{}

	kb_c_vm := bsource.KanikoBuilder_Container_VolumeMount{}

	kb_c_e1 := bsource.KanikoBuilder_Container_Env{}
	kb_c_e2 := bsource.KanikoBuilder_Container_Env{}

	kb_v := bsource.KanikoBuilder_Volume{}

	kb_v_i := bsource.KanikoBuilder_Volume_Item{}

	kb.APIVersion = "v1"
	kb.Kind = "Pod"
	kb.Metadata.Name = pod_bid
	kb.Spec.RestartPolicy = "Never"

	kb_c.Name = pod_bid
	kb_c.Image = "gcr.io/kaniko-project/executor:latest"
	kb_c.Args = append(kb_c.Args, "--dockerfile=Dockerfile")
	kb_c.Args = append(kb_c.Args, "--context="+v_repoaddr)
	kb_c.Args = append(kb_c.Args, "--destination="+v_regaddr)

	kb_c_vm.MountPath = "/kaniko/.docker"
	kb_c_vm.Name = "kaniko-secret"

	kb_c_e1.Name = "GIT_USERNAME"
	kb_c_e1.Value = repo_id

	kb_c_e2.Name = "GIT_PASSWORD"
	kb_c_e2.Value = repo_pw

	kb_v.Name = "kaniko-secret"
	kb_v.Secret.SecretName = sec_bid

	kb_v_i.Key = ".dockerconfigjson"
	kb_v_i.Path = "config.json"

	kb_v.Secret.Items = append(kb_v.Secret.Items, kb_v_i)

	kb_c.Env = append(kb_c.Env, kb_c_e1)
	kb_c.Env = append(kb_c.Env, kb_c_e2)

	kb_c.VolumeMounts = append(kb_c.VolumeMounts, kb_c_vm)

	kb.Spec.Containers = append(kb.Spec.Containers, kb_c)
	kb.Spec.Volumes = append(kb.Spec.Volumes, kb_v)

	yb, err := goya.Marshal(kb)

	if err != nil {

		close_msg := err.Error() + "\n"
		fp.Write([]byte(close_msg))
		ToolkitBuildCancelForce(ns_bid)
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, close_msg)
		return
	}

	build_yaml, err := runfs.SetBuildManifestPath(yb)

	if err != nil {

		close_msg := err.Error() + "\n"
		fp.Write([]byte(close_msg))
		ToolkitBuildCancelForce(ns_bid)
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, close_msg)
		return
	}

	cmd = exec.Command("kubectl", "-n", ns_bid, "apply", "-f", build_yaml)

	cmd.Stdout = fp

	cmd.Stderr = fp

	err = cmd.Run()

	if err != nil {
		fp.Write([]byte(err.Error()))
		ToolkitBuildCancelForce(ns_bid)
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	var outb, errb bytes.Buffer

	for {

		outb = bytes.Buffer{}
		errb = bytes.Buffer{}

		cmd = exec.Command("kubectl", "-n", ns_bid, "get", "pod", pod_bid, "--no-headers")

		cmd.Stdout = &outb

		cmd.Stderr = &errb

		err = cmd.Run()

		if err != nil {
			fp.Write([]byte(err.Error()))
			ToolkitBuildCancelForce(ns_bid)
			_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
			return
		}

		stdout_str := outb.String()
		_ = errb.String()

		if strings.Contains(stdout_str, "Completed") {

			break
		}

		if strings.Contains(stdout_str, "Error") {

			break
		}

		outb = bytes.Buffer{}

		errb = bytes.Buffer{}

		cmd = exec.Command("kubectl", "-n", ns_bid, "logs", pod_bid)

		cmd.Stdout = &outb

		cmd.Stderr = &errb

		err = cmd.Run()

		if err != nil {
			fp.Write([]byte(err.Error()))
			ToolkitBuildCancelForce(ns_bid)
			_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
			return
		}

		stdout_b := outb.Bytes()
		_ = errb.String()

		err = runfs.UpdateBuildLogExt(stdout_b)

		if err != nil {
			fp.Write([]byte(err.Error()))
			ToolkitBuildCancelForce(ns_bid)
			_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
			return
		}

		time.Sleep(time.Millisecond * 100)

	}

	outb = bytes.Buffer{}

	errb = bytes.Buffer{}

	cmd = exec.Command("kubectl", "-n", ns_bid, "logs", pod_bid)

	cmd.Stdout = &outb

	cmd.Stderr = &errb

	err = cmd.Run()

	if err != nil {
		fp.Write([]byte(err.Error()))
		ToolkitBuildCancelForce(ns_bid)
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	stdout_b := outb.Bytes()
	_ = errb.String()

	err = runfs.UpdateBuildLogExt(stdout_b)

	if err != nil {
		fp.Write([]byte(err.Error()))
		ToolkitBuildCancelForce(ns_bid)
		_ = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, err.Error())
		return
	}

	err = runfs.CloseFilePointerForUsrBuildLogAndMarkDone(fp, "SUCCESS")

	if err != nil {
		ToolkitBuildCancelForce(ns_bid)
		return
	}
	ToolkitBuildCancelForce(ns_bid)

	return

}

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

func ToolkitBuildImagesGetLogExt() ([]byte, error) {

	var ret_byte []byte

	log_b, err := runfs.GetUsrBuildLogExt()

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get log ext: %s", err.Error())
	}

	ret_byte = log_b

	return ret_byte, nil

}

func ToolkitBuildCancelForce(ns string) {

	cmd := exec.Command("kubectl", "delete", "namespace", ns)

	_ = cmd.Run()
}
