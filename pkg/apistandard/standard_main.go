package apistandard

import (
	"fmt"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/kubebase"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/kuberead"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/kubetoolkit"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/kubewrite"
)

func (asgi API_STD) Run(std_cmd API_INPUT) (API_OUTPUT, error) {

	var ret_api_out API_OUTPUT

	if v_failed := asgi.Verify(std_cmd); v_failed != nil {

		return ret_api_out, fmt.Errorf("run failed: %s", v_failed.Error())
	}

	cmd_id := std_cmd["id"]

	switch cmd_id {

	// case "NKADM-INSTENV":

	// case "NKADM-INSTENVRES":

	case "ADMIN-INSTWKOR":

		targetip := std_cmd["targetip"]
		targetid := std_cmd["targetid"]
		targetpw := std_cmd["targetpw"]
		localip := std_cmd["localip"]
		osnm := std_cmd["osnm"]
		cv := std_cmd["cv"]
		token := std_cmd["token"]

		var cmd_err error

		if token == "-" {

			token, cmd_err = kubebase.GetJoinToken()

			if cmd_err != nil {
				return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
			}

		}

		go kubebase.InstallWorkerOnRemote(targetip, targetid, targetpw, localip, osnm, cv, token)

		ret_api_out.BODY = string([]byte("remote worker installation started\n"))

	case "ADMIN-INSTVOLOR":

		targetip := std_cmd["targetip"]
		targetid := std_cmd["targetid"]
		targetpw := std_cmd["targetpw"]
		localip := std_cmd["localip"]

		go kubebase.InstallVolumeOnRemote(targetip, targetid, targetpw, localip)

		ret_api_out.BODY = string([]byte("remote volume installation started\n"))

	case "ADMIN-INSTTKOR":

		targetip := std_cmd["targetip"]
		targetid := std_cmd["targetid"]
		targetpw := std_cmd["targetpw"]

		go kubebase.InstallToolKitOnRemote(targetip, targetid, targetpw)

		ret_api_out.BODY = string([]byte("remote toolkit installation started\n"))

	case "ADMIN-INSTLOGOR":

		targetip := std_cmd["targetip"]
		targetid := std_cmd["targetid"]
		targetpw := std_cmd["targetpw"]

		b_out, cmd_err := kubebase.InstallLogOnRemote(targetip, targetid, targetpw)

		if cmd_err != nil {

			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())

		}

		ret_api_out.BODY = string(b_out)

	case "ADMIN-INIT":

		go kubebase.AdminInitNPIA()

		b_out := []byte("npia init started\n")

		ret_api_out.BODY = string(b_out)

	case "ADMIN-INITLOG":

		b_out, cmd_err := kubebase.AdminGetInitLog()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "SETTING-CRTNS":

		ns := std_cmd["ns"]
		repoaddr := std_cmd["repoaddr"]
		regaddr := std_cmd["regaddr"]

		b_out, cmd_err := kubebase.SettingCreateNamespace(ns, repoaddr, regaddr)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "SETTING-SETREPO":
		ns := std_cmd["ns"]
		repoaddr := std_cmd["repoaddr"]
		repoid := std_cmd["repoid"]
		repopw := std_cmd["repopw"]

		b_out, cmd_err := kubebase.SettingRepoInfo(ns, repoaddr, repoid, repopw)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "SETTING-SETREG":
		ns := std_cmd["ns"]
		regaddr := std_cmd["regaddr"]
		regid := std_cmd["regid"]
		regpw := std_cmd["regpw"]

		b_out, cmd_err := kubebase.SettingRegInfo(ns, regaddr, regid, regpw)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
		//	case "SETTING-CRTNSVOL":

	case "SETTING-CRTVOL":

		main_ns := std_cmd["ns"]
		target_ip := std_cmd["targetip"]

		b_out, cmd_err := kubebase.SettingCreateVolume(main_ns, target_ip)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "SETTING-CRTMON":

		b_out, cmd_err := kubebase.SettingCreateMonitoring()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "SETTING-CRTMONPERS":

		b_out, cmd_err := kubebase.SettingCreateMonitoringPersistent()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

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

		ret_api_out.BODY = string(b_out)

	case "TOOLKIT-BUILDLOG":

		b_out, cmd_err := kubetoolkit.ToolkitBuildImagesGetLog()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "TOOLKIT-PIPE":

		ns := std_cmd["ns"]
		repoaddr := std_cmd["repoaddr"]
		regaddr := std_cmd["regaddr"]

		go kubetoolkit.PipelineBuildStart(ns, repoaddr, regaddr)

		b_out := []byte("build pipeline started\n")

		ret_api_out.BODY = string(b_out)

	case "TOOLKIT-PIPELOG":

		b_out, cmd_err := kubetoolkit.PipelineBuildGetLog()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "TOOLKIT-PIPESETVAR":

		varnm := std_cmd["varnm"]
		varval := std_cmd["varval"]

		b_out, cmd_err := kubetoolkit.PipelineBuildSetVariablesEx(varnm, varval)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "TOOLKIT-PIPEGETVAR":

		b_out, cmd_err := kubetoolkit.PipelineBuildGetVariableMapEx()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NDS":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadNode(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-PDS":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadPod(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-PLOG":

		ns := std_cmd["ns"]

		pod_name := std_cmd["podnm"]

		b_out, cmd_err := kuberead.ReadPodLog(ns, pod_name)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-SVC":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadService(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-DPL":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadDeployment(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	//case "RESOURCE-IMGLI":
	case "RESOURCE-EVNT":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadEvent(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-RSRC":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadResource(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NSPC":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadNamespace(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-INGR":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadIngress(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NDPORT":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadNodePort(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	//case "RESOURCE-PRJPRB":
	case "RESOURCE-PSCH":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadPodScheduled(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-PUNSCH":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadPodUnscheduled(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-CCPU":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerCPUUsage(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-CMEM":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerMemUsage(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-CFSR":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerFSRead(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-CFSW":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerFSWrite(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-CNETR":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerNetworkReceive(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-CNETT":

		ns := std_cmd["ns"]

		b_out, cmd_err := kuberead.ReadContainerNetworkTransmit(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-VOLAVAIL":

		b_out, cmd_err := kuberead.ReadKubeletVolumeAvailable()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-VOLCAP":

		b_out, cmd_err := kuberead.ReadKubeletVolumeCapacity()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-VOLUSD":

		b_out, cmd_err := kuberead.ReadKubeletVolumeUsed()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NTEMP":

		b_out, cmd_err := kuberead.ReadNodeTemperatureCelsius()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NTEMPCH":

		b_out, cmd_err := kuberead.ReadNodeTemperatureCelsiusChange()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NTEMPAV":

		b_out, cmd_err := kuberead.ReadNodeTemperatureCelsiusAverage()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NPROCS":

		b_out, cmd_err := kuberead.ReadNodeProcessRunning()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NCORES":

		b_out, cmd_err := kuberead.ReadNodeCPUCores()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NMEM":

		b_out, cmd_err := kuberead.ReadNodeMemActive()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NMEMTOT":

		b_out, cmd_err := kuberead.ReadNodeMemTotal()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NDISKR":

		b_out, cmd_err := kuberead.ReadNodeDiskRead()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NDISKW":

		b_out, cmd_err := kuberead.ReadNodeDiskWrite()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NNETR":

		b_out, cmd_err := kuberead.ReadNodeNetworkReceive()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NNETT":
		b_out, cmd_err := kuberead.ReadNodeNetworkTransmit()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "RESOURCE-NDISKWT":

		b_out, cmd_err := kuberead.ReadNodeDiskWrittenTotal()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "APPLY-REGSEC":
		ns := std_cmd["ns"]

		b_out, cmd_err := kubewrite.WriteSecret(ns)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "APPLY-DIST":
		ns := std_cmd["ns"]
		repoaddr := std_cmd["repoaddr"]
		regaddr := std_cmd["regaddr"]

		b_out, cmd_err := kubewrite.WriteDeployment(ns, repoaddr, regaddr)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "APPLY-CRTOPSSRC":
		ns := std_cmd["ns"]
		repoaddr := std_cmd["repoaddr"]
		regaddr := std_cmd["regaddr"]

		b_out, cmd_err := kubewrite.WriteOperationSource(ns, repoaddr, regaddr)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-RESTART":
		ns := std_cmd["ns"]
		resource := std_cmd["resource"]
		resourcenm := std_cmd["resourcenm"]
		b_out, cmd_err := kubewrite.WriteUpdateOrRestart(ns, resource, resourcenm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)

	case "APPLY-ROLLBACK":

		ns := std_cmd["ns"]
		resource := std_cmd["resource"]
		resourcenm := std_cmd["resourcenm"]
		b_out, cmd_err := kubewrite.WriteRollback(ns, resource, resourcenm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-KILL":
		ns := std_cmd["ns"]
		resource := std_cmd["resource"]
		resourcenm := std_cmd["resourcenm"]
		b_out, cmd_err := kubewrite.WriteDeletion(ns, resource, resourcenm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-NETRESH":

		b_out, cmd_err := kubewrite.WriteNetworkRefresh()

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-HPA":
		ns := std_cmd["ns"]
		resource := std_cmd["resource"]
		resourcenm := std_cmd["resourcenm"]
		b_out, cmd_err := kubewrite.WriteHPA(ns, resource, resourcenm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-HPAUN":
		ns := std_cmd["ns"]
		resource := std_cmd["resource"]
		resourcenm := std_cmd["resourcenm"]
		b_out, cmd_err := kubewrite.WriteHPAUndo(ns, resource, resourcenm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-QOS":
		ns := std_cmd["ns"]
		resource := std_cmd["resource"]
		resourcenm := std_cmd["resourcenm"]
		b_out, cmd_err := kubewrite.WriteQOS(ns, resource, resourcenm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-QOSUN":
		ns := std_cmd["ns"]
		resource := std_cmd["resource"]
		resourcenm := std_cmd["resourcenm"]
		b_out, cmd_err := kubewrite.WriteQOSUndo(ns, resource, resourcenm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-INGR":
		ns := std_cmd["ns"]
		hostnm := std_cmd["hostnm"]
		svcnm := std_cmd["svcnm"]
		b_out, cmd_err := kubewrite.WriteIngress(ns, hostnm, svcnm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-INGRUN":
		ns := std_cmd["ns"]
		hostnm := std_cmd["hostnm"]
		svcnm := std_cmd["svcnm"]
		b_out, cmd_err := kubewrite.WriteIngressUndo(ns, hostnm, svcnm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-NDPORT":
		ns := std_cmd["ns"]
		svcnm := std_cmd["svcnm"]
		b_out, cmd_err := kubewrite.WriteNodePort(ns, svcnm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	case "APPLY-NDPORTUN":
		ns := std_cmd["ns"]
		svcnm := std_cmd["svcnm"]

		b_out, cmd_err := kubewrite.WriteNodePortUndo(ns, svcnm)

		if cmd_err != nil {
			return ret_api_out, fmt.Errorf("run failed: %s", cmd_err.Error())
		}

		ret_api_out.BODY = string(b_out)
	// case "EXIT":
	default:

		return ret_api_out, fmt.Errorf("failed to run api: %s", "invalid command id")

	}

	return ret_api_out, nil

}
