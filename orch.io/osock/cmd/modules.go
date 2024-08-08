package cmd

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"

	models "github.com/OKESTRO-AIDevOps/nkia/orch.io/osock/models"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/kubebase"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/kubetoolkit"
)

func RequestForwardHandler(email string, query string) (bool, []byte, error) {

	var result []byte

	ASgi := apistandard.ASgi

	api_input, err := ASgi.StdCmdInputBuildFromLinearInstruction(query)

	if err != nil {

		return false, nil, fmt.Errorf("run failed: %s", err.Error())

	}

	if v_failed := ASgi.Verify(api_input); v_failed != nil {

		return false, nil, fmt.Errorf("run failed: %s", v_failed.Error())

	}

	cmd_id := api_input["id"]

	switch cmd_id {

	case "ORCH-CONNCHK":

		var talkback string = "alive\n"

		out := apistandard.API_OUTPUT{
			BODY: talkback,
		}

		result, err := json.Marshal(out)

		if err != nil {

			return false, nil, fmt.Errorf("admin req: %s", err.Error())

		}

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

		err = models.UpdatePubkeyByEmail2(email, pub_pem_str)

		if err != nil {
			return false, result, fmt.Errorf("admin req: %s", err.Error())
		}

		out := apistandard.API_OUTPUT{
			BODY: string(priv_pem),
		}

		result, err := json.Marshal(out)

		if err != nil {

			return false, nil, fmt.Errorf("admin req: %s", err.Error())

		}
		return false, result, nil

	case "ORCH-GETCL":

		out := apistandard.API_OUTPUT{
			BODY: "not implemented",
		}

		result, err := json.Marshal(out)

		if err != nil {

			return false, nil, fmt.Errorf("admin req: %s", err.Error())

		}
		return false, result, nil

	case "ORCH-ADDCL":

		clusterid := api_input["clusterid"]

		token, err := models.CreateClusterByEmail2(email, clusterid)

		if err != nil {
			return false, result, fmt.Errorf("admin req: %s", err.Error())
		}

		out := apistandard.API_OUTPUT{
			BODY: token,
		}

		result, err := json.Marshal(out)

		if err != nil {

			return false, nil, fmt.Errorf("admin req: %s", err.Error())

		}

		return false, result, nil

	case "ORCH-INSTCL":

		if len(models.FI_SESSIONS.INST_SESSION) > 100 {
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

		_, okay := models.FI_SESSIONS.INST_SESSION[session_key]

		if okay {
			return false, result, fmt.Errorf("admin req: already an ongoing installation")
		}

		models.FI_SESSIONS.INST_SESSION[session_key] = &[]byte{}

		models.FI_SESSIONS.INST_RESULT[session_key] = "-"

		go models.InstallCluster(session_key, clusterid, targetip, targetid, targetpw, localip, osnm, cv, updatetoken)

		out := apistandard.API_OUTPUT{
			BODY: "remote cluster installation started\n",
		}

		result, err := json.Marshal(out)

		if err != nil {

			return false, nil, fmt.Errorf("admin req: %s", err.Error())

		}

		return false, result, nil

	case "ORCH-INSTCLLOG":

		/*
			clusterid := api_input["clusterid"]
			targetip := api_input["targetip"]
			targetid := api_input["targetid"]
			targetpw := api_input["targetpw"]

			session_key := email + ":" + clusterid

			log_b, err := sctrl.InstallClusterLog(session_key, clusterid, targetip, targetid, targetpw)

			if err != nil {
				return false, result, fmt.Errorf("admin req: %s", err.Error())
			}

			out := apistandard.API_OUTPUT{
				BODY: string(log_b),
			}

			result, err := json.Marshal(out)

			if err != nil {

				return false, nil, fmt.Errorf("admin req: %s", err.Error())

			}

		*/

		out := apistandard.API_OUTPUT{
			BODY: "do not use it for now",
		}

		result, err := json.Marshal(out)

		if err != nil {

			return false, nil, fmt.Errorf("admin req: %s", err.Error())

		}
		return false, result, nil

	default:

		fmt.Println("forward")

	}

	return true, nil, nil

}

func OrchestrationRequestHandler(email string, std_cmd apistandard.API_INPUT) ([]byte, error) {

	cmd_id := std_cmd["id"]

	switch cmd_id {
	case "ORCH-SETTING-CRTNS":

		ns := std_cmd["ns"]
		repoaddr := std_cmd["repoaddr"]
		regaddr := std_cmd["regaddr"]

		b_out, cmd_err := kubebase.SettingCreateNamespace(ns, repoaddr, regaddr)

		if cmd_err != nil {
			return nil, fmt.Errorf("orchio failed: %s", cmd_err.Error())
		}

		out := apistandard.API_OUTPUT{
			BODY: string(b_out),
		}

		result, err := json.Marshal(out)

		if err != nil {

			return nil, fmt.Errorf("orchio req: %s", err.Error())

		}

		return result, nil

	case "SETTING-SETREPO":
		ns := std_cmd["ns"]
		repoaddr := std_cmd["repoaddr"]
		repoid := std_cmd["repoid"]
		repopw := std_cmd["repopw"]

		b_out, cmd_err := kubebase.SettingRepoInfo(ns, repoaddr, repoid, repopw)

		if cmd_err != nil {
			return nil, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		out := apistandard.API_OUTPUT{
			BODY: string(b_out),
		}

		result, err := json.Marshal(out)

		if err != nil {

			return nil, fmt.Errorf("orchio req: %s", err.Error())

		}

		return result, nil

	case "SETTING-SETREG":
		ns := std_cmd["ns"]
		regaddr := std_cmd["regaddr"]
		regid := std_cmd["regid"]
		regpw := std_cmd["regpw"]

		b_out, cmd_err := kubebase.SettingRegInfo(ns, regaddr, regid, regpw)

		if cmd_err != nil {
			return nil, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		out := apistandard.API_OUTPUT{
			BODY: string(b_out),
		}

		result, err := json.Marshal(out)

		if err != nil {

			return nil, fmt.Errorf("orchio req: %s", err.Error())

		}

		return result, nil

	case "SETTING-CRTVOL":

		/*
			main_ns := std_cmd["ns"]
			target_ip := std_cmd["targetip"]

			b_out, cmd_err := kubebase.SettingCreateVolume(main_ns, target_ip)

			if cmd_err != nil {
				return nil, fmt.Errorf("run failed: %s", cmd_err.Error())
			}

			out := apistandard.API_OUTPUT{
				BODY: string(b_out),
			}

			result, err := json.Marshal(out)

			if err != nil {

				return nil, fmt.Errorf("orchio req: %s", err.Error())

			}

		*/

		result := []byte("do not use: create vol\n")

		return result, nil

	case "SETTING-CRTMON":

		/*
			b_out, cmd_err := kubebase.SettingCreateMonitoring()

			if cmd_err != nil {
				return nil, fmt.Errorf("run failed: %s", cmd_err.Error())
			}
			out := apistandard.API_OUTPUT{
				BODY: string(b_out),
			}

			result, err := json.Marshal(out)

			if err != nil {

				return nil, fmt.Errorf("orchio req: %s", err.Error())

			}

		*/
		result := []byte("do not use: create mon\n")

		return result, nil

	case "SETTING-CRTMONPERS":

		/*
			b_out, cmd_err := kubebase.SettingCreateMonitoringPersistent()

			if cmd_err != nil {
				return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
			}

			ret_api_out.BODY = string(b_out)

		*/
		result := []byte("do not use: create mon persistent\n")

		return result, nil

		//	case "SETTING-DELNS":
		//	case "SUBMIT":
		//	case "CALLME":
		//	case "GITLOG":
		//	case "PIPEHIST":
		//	case "PIPE":
		//	case "PIPELOG":
	case "TOOLKIT-BUILD":

		ns := std_cmd["ns"]
		repoaddr := std_cmd["repoaddr"]
		regaddr := std_cmd["regaddr"]

		go kubetoolkit.ToolkitBuildImagesStart(ns, repoaddr, regaddr)

		b_out := []byte("build images started\n")

		out := apistandard.API_OUTPUT{
			BODY: string(b_out),
		}

		result, err := json.Marshal(out)

		if err != nil {

			return nil, fmt.Errorf("orchio req: %s", err.Error())

		}

		return result, nil

	case "TOOLKIT-BUILDLOG":

		b_out, cmd_err := kubetoolkit.ToolkitBuildImagesGetLog()

		if cmd_err != nil {
			return nil, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		out := apistandard.API_OUTPUT{
			BODY: string(b_out),
		}

		result, err := json.Marshal(out)

		if err != nil {

			return nil, fmt.Errorf("orchio req: %s", err.Error())

		}

		return result, nil

	case "TOOLKIT-PIPE":

		ns := std_cmd["ns"]
		repoaddr := std_cmd["repoaddr"]
		regaddr := std_cmd["regaddr"]

		go kubetoolkit.PipelineBuildStart(ns, repoaddr, regaddr)

		b_out := []byte("build pipeline started\n")

		out := apistandard.API_OUTPUT{
			BODY: string(b_out),
		}

		result, err := json.Marshal(out)

		if err != nil {

			return nil, fmt.Errorf("orchio req: %s", err.Error())

		}

		return result, nil

	case "TOOLKIT-PIPELOG":

		b_out, cmd_err := kubetoolkit.PipelineBuildGetLog()

		if cmd_err != nil {
			return nil, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		out := apistandard.API_OUTPUT{
			BODY: string(b_out),
		}

		result, err := json.Marshal(out)

		if err != nil {

			return nil, fmt.Errorf("orchio req: %s", err.Error())

		}

		return result, nil

	case "TOOLKIT-PIPESETVAR":

		varnm := std_cmd["varnm"]
		varval := std_cmd["varval"]

		b_out, cmd_err := kubetoolkit.PipelineBuildSetVariablesEx(varnm, varval)

		if cmd_err != nil {
			return nil, fmt.Errorf("run failed: %s", cmd_err.Error())
		}
		out := apistandard.API_OUTPUT{
			BODY: string(b_out),
		}

		result, err := json.Marshal(out)

		if err != nil {

			return nil, fmt.Errorf("orchio req: %s", err.Error())

		}

		return result, nil

	case "TOOLKIT-PIPEGETVAR":

		b_out, cmd_err := kubetoolkit.PipelineBuildGetVariableMapEx()

		if cmd_err != nil {
			return nil, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		out := apistandard.API_OUTPUT{
			BODY: string(b_out),
		}

		result, err := json.Marshal(out)

		if err != nil {

			return nil, fmt.Errorf("orchio req: %s", err.Error())

		}

		return result, nil

	default:

		return nil, fmt.Errorf("failed to run: no such cmd id: %s", cmd_id)

	}

}
