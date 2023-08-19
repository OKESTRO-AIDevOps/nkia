package kubewrite

import (
	"fmt"

	runfs "github.com/OKESTRO-AIDevOps/nkia/pkg/runtimefs"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/libinterface"

	"os/exec"
)

func WriteSecret(main_ns string) ([]byte, error) {

	var ret_byte []byte

	app_origin, err := runfs.LoadAdmOrigin()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ns_found, _, reg_url := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		return ret_byte, fmt.Errorf(": %s", "no such namespace")
	}

	if reg_url == "N" {

		return ret_byte, fmt.Errorf(": %s", "reg url not set")

	}

	addr_found, reg_id, reg_pw := runfs.GetRegInfo(app_origin.REGS, reg_url)

	if !addr_found {

		return ret_byte, fmt.Errorf(": %s", "reg info not complete")

	}

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "secret", "docker-secret", "--no-headers", "-o", "custom-columns=:metadata.name")

	_, err = cmd.Output()

	docker_server := "--docker-server="

	docker_username := "--docker-username="

	docker_password := "--docker-password="

	docker_server += reg_url

	docker_username += reg_id

	docker_password += reg_pw

	if err != nil {

		cmd = exec.Command("kubectl", "-n", main_ns, "create", "secret", "docker-registry", "docker-secret", docker_server, docker_username, docker_password)

		out, err := cmd.Output()

		if err != nil {
			return ret_byte, fmt.Errorf(": %s", err.Error())
		}

		ret_byte = out

		return ret_byte, nil

	} else {

		cmd = exec.Command("kubectl", "-n", main_ns, "delete", "secret", "docker-secret")

		_, err = cmd.Output()

		if err != nil {
			return ret_byte, fmt.Errorf(": %s", err.Error())
		}

		cmd = exec.Command("kubectl", "-n", main_ns, "create", "secret", "docker-registry", "docker-secret", docker_server, docker_username, docker_password)

		out, err := cmd.Output()

		if err != nil {
			return ret_byte, fmt.Errorf(": %s", err.Error())
		}

		ret_byte = out

		return ret_byte, nil

	}

}

func WriteDeployment(main_ns string, repoaddr string, regaddr string) ([]byte, error) {

	var ret_byte []byte

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	LIBIF_BIN_KOMPOSE, err := libif.GetLibComponentAddress("bin", "kompose")

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	var app_origin runfs.AppOrigin

	app_origin, err = runfs.LoadAdmOrigin()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ns_found, repoaddr, _ := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		return ret_byte, fmt.Errorf(": %s", "no such namespace")
	}

	err = runfs.InitUsrTarget(repoaddr)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	USR_OPS_SRC, err := runfs.CreateUsrTargetOperationSource(LIBIF_BIN_KOMPOSE, regaddr)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	cmd := exec.Command("kubectl", "-n", main_ns, "apply", "-f", USR_OPS_SRC)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func WriteOperationSource(main_ns string, repoaddr string, regaddr string) ([]byte, error) {

	var ret_byte []byte

	libif, err := libinterface.ConstructLibIface()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	LIBIF_BIN_KOMPOSE, err := libif.GetLibComponentAddress("bin", "kompose")

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	app_origin, err := runfs.LoadAdmOrigin()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ns_found, repoaddr, _ := runfs.GetRecordInfo(app_origin.RECORDS, main_ns)

	if !ns_found {
		return ret_byte, fmt.Errorf(": %s", "no such namespace")
	}

	err = runfs.InitUsrTarget(repoaddr)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	USR_OPS_SRC, err := runfs.CreateUsrTargetOperationSource(LIBIF_BIN_KOMPOSE, regaddr)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = append(ret_byte, []byte(USR_OPS_SRC)...)

	ret_byte = append(ret_byte, []byte(" created successfully")...)

	return ret_byte, nil

}

func WriteUpdateOrRestart(main_ns string, resource string, resourcenm string) ([]byte, error) {

	var ret_byte []byte

	rsc_rscnm := resource + "/" + resourcenm

	cmd := exec.Command("kubectl", "-n", main_ns, "rollout", "restart", rsc_rscnm)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func WriteRollback(main_ns string, resource string, resourcenm string) ([]byte, error) {

	var ret_byte []byte

	rsc_rscnm := resource + "/" + resourcenm

	cmd := exec.Command("kubectl", "-n", main_ns, "rollout", "undo", rsc_rscnm)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil
}

func WriteDeletion(main_ns string, resource string, resourcenm string) ([]byte, error) {

	var ret_byte []byte

	USR_DEL_OPS_SRC, err := runfs.CreateUsrDelOperationSource(resourcenm)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	cmd := exec.Command("kubectl", "-n", main_ns, "delete", "-f", USR_DEL_OPS_SRC)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func WriteNetworkRefresh() ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "-n", "kube-system", "rollout", "restart", "deployment/coredns")

	_, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = []byte("network refreshed")

	return ret_byte, nil

}

func WriteHPA(main_ns string, resource string, resourcenm string) ([]byte, error) {

	var ret_byte []byte

	var resource_key string

	if resource == "deployment" {
		resource_key = "Deployment"

	} else {
		return ret_byte, fmt.Errorf(": %s", "not a deployment")
	}

	USR_HPA_SRC, err := runfs.CreateHPASource(resourcenm, resource_key)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	cmd := exec.Command("kubectl", "-n", main_ns, "apply", "-f", USR_HPA_SRC)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func WriteHPAUndo(main_ns string, resource string, resourcenm string) ([]byte, error) {

	var ret_byte []byte

	var resource_key string

	if resource == "deployment" {
		resource_key = "Deployment"

	} else {
		return ret_byte, fmt.Errorf(": %s", "not a deployment")
	}

	USR_HPA_SRC, err := runfs.CreateHPASource(resourcenm, resource_key)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	cmd := exec.Command("kubectl", "-n", main_ns, "delete", "-f", USR_HPA_SRC)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func WriteQOS(main_ns string, resource string, resourcenm string) ([]byte, error) {

	var ret_byte []byte

	var resource_key string

	if resource == "deployment" {
		resource_key = "Deployment"

	} else {
		return ret_byte, fmt.Errorf(": %s", "not a deployment")
	}

	USR_QOS_SRC, err := runfs.CreateQOSSource(resourcenm, resource_key)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	cmd := exec.Command("kubectl", "-n", main_ns, "apply", "-f", USR_QOS_SRC)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func WriteQOSUndo(main_ns string, resource string, resourcenm string) ([]byte, error) {

	var ret_byte []byte

	var resource_key string

	if resource == "deployment" {
		resource_key = "Deployment"

	} else {
		return ret_byte, fmt.Errorf(": %s", "not a deployment")
	}

	USR_DEL_QOS_SRC, err := runfs.CreateDelQOSSource(resourcenm, resource_key)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	cmd := exec.Command("kubectl", "-n", main_ns, "apply", "-f", USR_DEL_QOS_SRC)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil
}

func WriteIngress(main_ns string, hostnm string, svcnm string) ([]byte, error) {

	var ret_byte []byte

	USR_INGR_SRC, err := runfs.CreateIngressSource(main_ns, hostnm, svcnm)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	cmd := exec.Command("kubectl", "-n", main_ns, "apply", "-f", USR_INGR_SRC)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func WriteIngressUndo(main_ns string, hostnm string, svcnm string) ([]byte, error) {

	var ret_byte []byte

	USR_DEL_INGR_SRC, err := runfs.CreateIngressSource(main_ns, hostnm, svcnm)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	cmd := exec.Command("kubectl", "-n", main_ns, "delete", "-f", USR_DEL_INGR_SRC)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func WriteNodePort(main_ns string, svcnm string) ([]byte, error) {

	var ret_byte []byte

	USR_NDPT_SRC, err := runfs.CreateNodePortSource(main_ns, svcnm)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	cmd := exec.Command("kubectl", "-n", main_ns, "apply", "-f", USR_NDPT_SRC)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func WriteNodePortUndo(main_ns string, svcnm string) ([]byte, error) {

	var ret_byte []byte

	USR_DEL_NDPT_SRC, err := runfs.CreateNodePortSource(main_ns, svcnm)

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	cmd := exec.Command("kubectl", "-n", main_ns, "delete", "-f", USR_DEL_NDPT_SRC)

	out, err := cmd.Output()

	if err != nil {
		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}
