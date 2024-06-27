package cmd

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	sctrl "github.com/OKESTRO-AIDevOps/nkia/orch.io/osock/controller"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
)

func RequestForwardHandler(email string, query string) (bool, string, error) {

	var result string = ""

	ASgi := apistandard.ASgi

	api_input, err := ASgi.StdCmdInputBuildFromLinearInstruction(query)

	if err != nil {

		return false, "", fmt.Errorf("run failed: %s", err.Error())

	}

	if v_failed := ASgi.Verify(api_input); v_failed != nil {

		return false, "", fmt.Errorf("run failed: %s", v_failed.Error())

	}

	cmd_id := api_input["id"]

	switch cmd_id {

	case "ORCH-CONNCHK":

		var talkback string = "talking back list: "

		for el := range api_input {

			talkback += el + " "

		}

		talkback += "\n"

		result = talkback

		return false, result, nil

	case "ORCH-KEYGEN":

		privkey, pubkey, err := modules.GenerateKeyPair(4096)

		if err != nil {
			return false, result, fmt.Errorf("admin req: %s", err.Error())
		}

		priv_pem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(privkey),
			},
		)

		pub_b, err := x509.MarshalPKIXPublicKey(pubkey)

		if err != nil {
			return false, result, fmt.Errorf("admin req: %s", err.Error())
		}

		pub_pem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "PUBLIC KEY",
				Bytes: pub_b,
			},
		)

		pub_pem_str := string(pub_pem)

		err = sctrl.UpdatePubkeyByEmail(email, pub_pem_str)

		if err != nil {
			return false, result, fmt.Errorf("admin req: %s", err.Error())
		}

		result = string(priv_pem)

		return false, result, nil

	case "ORCH-GETCL":

		return false, result, fmt.Errorf("admin req: %s", "not implemented")

	case "ORCH-ADDCL":

		clusterid := api_input["clusterid"]

		token, err := sctrl.CreateClusterByEmail(email, clusterid)

		if err != nil {
			return false, result, fmt.Errorf("admin req: %s", err.Error())
		}

		result = token

		return false, result, nil

	case "ORCH-INSTCL":

		if len(sctrl.FI_SESSIONS.INST_SESSION) > 100 {
			return false, result, fmt.Errorf("admin req: too many remote install sessions")
		}

		clusterid := api_input["clusterid"]
		targetip := api_input["targetip"]
		targetid := api_input["targetid"]
		targetpw := api_input["targetpw"]
		localip := api_input["localip"]
		osnm := api_input["osnm"]
		cv := api_input["cv"]
		updatetoken := api_input["updatetoken"]

		session_key := email + ":" + clusterid

		_, okay := sctrl.FI_SESSIONS.INST_SESSION[session_key]

		if okay {
			return false, result, fmt.Errorf("admin req: already an ongoing installation")
		}

		sctrl.FI_SESSIONS.INST_SESSION[session_key] = &[]byte{}

		sctrl.FI_SESSIONS.INST_RESULT[session_key] = "-"

		go sctrl.InstallCluster(session_key, clusterid, targetip, targetid, targetpw, localip, osnm, cv, updatetoken)

		result = "remote cluster installation started\n"

		return false, result, nil

	case "ORCH-INSTCLLOG":

		clusterid := api_input["clusterid"]
		targetip := api_input["targetip"]
		targetid := api_input["targetid"]
		targetpw := api_input["targetpw"]

		session_key := email + ":" + clusterid

		log_b, err := sctrl.InstallClusterLog(session_key, clusterid, targetip, targetid, targetpw)

		if err != nil {
			return false, result, fmt.Errorf("admin req: %s", err.Error())
		}

		result = string(log_b)

		return false, result, nil

	default:

		fmt.Println("forward")

	}

	return true, "", nil

}
