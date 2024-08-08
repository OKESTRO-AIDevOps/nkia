package models

import (
	"fmt"
	"log"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/utils"
)

func InstallCluster(sess_key string, cluster_id string, targetip string, targetid string, targetpw string, localip string, osnm string, cv string, update_token string) {

	conn, err := utils.ShellConnect(targetip, targetid, targetpw)

	if err != nil {

		err_msg := fmt.Sprintf("failed to install cluster: %s", err.Error())

		WriteToInstallSessionWithLock(sess_key, err_msg)

		WriteToInstallResultWithLock(sess_key, err_msg)

		return

	}

	output, err := conn.SendCommands("sudo mkdir -p /npia && ls -la /npia")

	if err != nil {

		err_msg := fmt.Sprintf("failed to install cluster: %s", err.Error())

		WriteToInstallSessionWithLock(sess_key, err_msg)

		WriteToInstallResultWithLock(sess_key, err_msg)

		return
	}

	output = append(output, []byte("\n----------ROOT NPIA CREATED----------\n")...)

	WriteToInstallSessionWithLock(sess_key, string(output))

	output, err = conn.SendCommands("sudo curl -L https://github.com/OKESTRO-AIDevOps/nkia/releases/download/latest/bin.tgz -o /npia/bin.tgz")
	if err != nil {

		err_msg := fmt.Sprintf("failed to install cluster: %s", err.Error())

		WriteToInstallSessionWithLock(sess_key, err_msg)

		WriteToInstallResultWithLock(sess_key, err_msg)

		return
	}

	output = append(output, []byte("\n----------NPIA BIN DOWNLOADED----------\n")...)

	WriteToInstallSessionWithLock(sess_key, string(output))

	output, err = conn.SendCommands("sudo tar -xzf /npia/bin.tgz -C /npia")
	if err != nil {

		err_msg := fmt.Sprintf("failed to install cluster: %s", err.Error())

		WriteToInstallSessionWithLock(sess_key, err_msg)

		WriteToInstallResultWithLock(sess_key, err_msg)

		return
	}

	output = append(output, []byte("\n----------NPIA BIN INSTALLED----------\n")...)

	WriteToInstallSessionWithLock(sess_key, string(output))

	output, err = conn.SendCommands("cd /npia/bin/nokubeadm && sudo ./nokubeadm init-npia-default")
	if err != nil {

		err_msg := fmt.Sprintf("failed to install cluster: %s", err.Error())

		WriteToInstallSessionWithLock(sess_key, err_msg)

		WriteToInstallResultWithLock(sess_key, err_msg)

		return
	}

	output = append(output, []byte("\n----------NPIA INITIATED----------\n")...)

	WriteToInstallSessionWithLock(sess_key, string(output))

	options := " " + "--localip " + localip + " " + "--osnm " + osnm + " " + "--cv " + cv

	output, err = conn.SendCommands("cd /npia/bin/nokubeadm && sudo ./nokubeadm install mainctrl" + options)
	if err != nil {

		err_msg := fmt.Sprintf("failed to install cluster: %s", err.Error())

		WriteToInstallSessionWithLock(sess_key, fmt.Sprintf("failed to install cluster: %s", err.Error()))

		WriteToInstallResultWithLock(sess_key, err_msg)

		return
	}

	output = append(output, []byte("\n----------CONTROL PLANE INSTALLED----------\n")...)

	WriteToInstallSessionWithLock(sess_key, string(output))

	options = " " + "--clusterid " + cluster_id + " " + "--updatetoken " + update_token

	output, err = conn.SendCommands("cd /npia/bin/nokubelet && sudo nohup ./nkletd io connect update" + options)
	if err != nil {

		err_msg := fmt.Sprintf("failed to install cluster: %s", err.Error())

		WriteToInstallSessionWithLock(sess_key, err_msg)

		WriteToInstallResultWithLock(sess_key, err_msg)

		return
	}

	output = append(output, []byte("\n----------NOKUBELET CONNECTED----------\n")...)

	WriteToInstallSessionWithLock(sess_key, string(output))

	WriteToInstallResultWithLock(sess_key, "SUCCESS")

	return
}

func InstallClusterLog(sess_key string, cluster_id string, targetip string, targetid string, targetpw string) ([]byte, error) {

	var ret_byte []byte

	result, err := ReadFromInstallResult(sess_key)

	if err != nil {
		return ret_byte, fmt.Errorf("install result: %s", err.Error())
	}

	res_str := string(result)

	if res_str == "SUCCESS" {

		ret_byte = []byte("SUCCESS")

		RemoveFromInstallSessionWithLock(sess_key)

		return ret_byte, nil

	} else if res_str == "-" {

		sess_b, err := ReadFromInstallSession(sess_key)

		if err != nil {
			return ret_byte, fmt.Errorf("install log: %s", err.Error())
		}

		log_b, _ := FetchInstallClusterLog(sess_key, targetip, targetid, targetpw)

		//if err != nil {
		//	return ret_byte, fmt.Errorf("install log: %s", err.Error())
		//}

		ret_byte = MergeInstallSessionAndLog(sess_b, log_b)

	} else {

		sess_b, err := ReadFromInstallSession(sess_key)

		if err != nil {
			return ret_byte, fmt.Errorf("install log: %s", err.Error())
		}

		ret_byte = append(ret_byte, sess_b...)

		RemoveFromInstallSessionWithLock(sess_key)

		return ret_byte, nil

	}

	return ret_byte, nil

}

func WriteToInstallSessionWithLock(sess_key string, new_msg string) {

	FI_SESSIONS.Mu.Lock()

	defer FI_SESSIONS.Mu.Unlock()

	new_bytes := []byte(new_msg)

	ret_byte, okay := FI_SESSIONS.INST_SESSION[sess_key]

	if !okay {
		log.Println("in write to install sess could not find session for: " + sess_key)
		return
	}

	*ret_byte = append(*ret_byte, new_bytes...)

	return
}

func WriteToInstallResultWithLock(sess_key string, result string) {

	FI_SESSIONS.Mu.Lock()

	defer FI_SESSIONS.Mu.Unlock()

	_, okay := FI_SESSIONS.INST_SESSION[sess_key]

	if !okay {
		log.Println("in write to install result could not find session for: " + sess_key)
		return
	}

	FI_SESSIONS.INST_RESULT[sess_key] = result

	return

}

func ReadFromInstallSession(sess_key string) ([]byte, error) {

	var ret_byte []byte

	sess_log, okay := FI_SESSIONS.INST_SESSION[sess_key]

	if !okay {
		msg := "in remove session could not find session for: " + sess_key
		log.Println(msg)
		return ret_byte, fmt.Errorf(msg)
	}

	ret_byte = *sess_log

	return ret_byte, nil

}

func ReadFromInstallResult(sess_key string) ([]byte, error) {

	var ret_byte []byte

	result, okay := FI_SESSIONS.INST_RESULT[sess_key]

	if !okay {
		msg := "in remove session could not find result for: " + sess_key
		log.Println(msg)
		return ret_byte, fmt.Errorf(msg)
	}

	ret_byte = []byte(result)

	return ret_byte, nil

}

func RemoveFromInstallSessionWithLock(sess_key string) {

	FI_SESSIONS.Mu.Lock()

	defer FI_SESSIONS.Mu.Unlock()

	_, okay := FI_SESSIONS.INST_SESSION[sess_key]

	if !okay {
		log.Println("in remove session could not find session for: " + sess_key)
		return
	}

	_, okay = FI_SESSIONS.INST_RESULT[sess_key]

	if !okay {
		log.Println("in remove session could not find result for: " + sess_key)
		return
	}

	delete(FI_SESSIONS.INST_SESSION, sess_key)

	delete(FI_SESSIONS.INST_RESULT, sess_key)

	return

}

func FetchInstallClusterLog(sess_key string, targetip string, targetid string, targetpw string) ([]byte, error) {

	var ret_byte []byte

	conn, err := utils.ShellConnect(targetip, targetid, targetpw)

	if err != nil {

		message := fmt.Sprintf("failed to install cluster: %s", err.Error())

		WriteToInstallSessionWithLock(sess_key, message)

		return ret_byte, fmt.Errorf(message)

	}

	output, err := conn.SendCommands("cd /npia/bin/nokubeadm && sudo ./nokubeadm install log")

	if err != nil {

		message := fmt.Sprintf("failed to install cluster: %s", err.Error())
		WriteToInstallSessionWithLock(sess_key, message)

		return ret_byte, fmt.Errorf(message)
	}

	ret_byte = output

	return ret_byte, nil
}

func MergeInstallSessionAndLog(sess_b []byte, log_b []byte) []byte {

	var ret_byte []byte

	ret_byte = append(ret_byte, sess_b...)

	ret_byte = append(ret_byte, []byte("\n\n----------**********----------\n\n---------*   LOG    *---------\n\n----------**********----------\n\n")...)

	ret_byte = append(ret_byte, log_b...)

	return ret_byte
}
