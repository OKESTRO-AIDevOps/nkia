package apix

import "strings"

type API_X map[string]string

var APIX_DEFINITION = "" +
	// front admin option
	"ConnectionCheck                                   : CONNCHK               " + "\n" +
	"Keygen                                            : KEYGEN                " + "\n" +
	"AddCluster                                        : ADDCLUSTER            " + "\n" +
	// query
	"Admin_InstallEnvironment                          : ADMIN-INSTENV         " + "\n" +
	"Admin_InstallControlPlane                         : ADMIN-INSTCTRL        " + "\n" +
	"Admin_InstallAnotherControlPlaneOnLocal           : ADMIN-INSTANCTRLOL    " + "\n" +
	"Admin_InstallAnotherControlPlaneOnRemote          : ADMIN-INSTANCTRLOR    " + "\n" +
	"Admin_InstallWorker                               : ADMIN-INSTWK          " + "\n" +
	"Admin_InstallVolumeOnRemote                       : ADMIN-INSTVOLOR       " + "\n" +
	"Admin_InstallVolumeOnLocal                        : ADMIN-INSTVOLOL       " + "\n" +
	"Admin_InstallToolkit                              : ADMIN-INSTTK          " + "\n" +
	"Admin_InstallLog                                  : ADMIN-INSTLOG         " + "\n" +
	"Admin_InstallLockGet                              : ADMIN-INSTLOCKGET     " + "\n" +
	"Admin_InstallLockSet                              : ADMIN-INSTLOCKSET     " + "\n" +
	"Admin_Init                                        : ADMIN-INIT            " + "\n" +
	"Admin_InitLog                                     : ADMIN-INITLOG         " + "\n" +
	"Setting_CreateNamespace                           : SETTING-CRTNS         " + "\n" +
	"Setting_SetRepository                             : SETTING-SETREPO       " + "\n" +
	"Setting_SetRegistry                               : SETTING-SETREG        " + "\n" +
	"Setting_CreateMonitoring                          : SETTING-CRTMON        " + "\n" +
	"Toolkit_Build                                     : TOOLKIT-BUILD         " + "\n" +
	"Toolkit_BuildLog                                  : TOOLKIT-BUILDLOG      " + "\n" +
	"Resource_Nodes                                    : RESOURCE-NDS          " + "\n" +
	"Resource_Pods                                     : RESOURCE-PDS          " + "\n" +
	"Resource_PodLog                                   : RESOURCE-PLOG         " + "\n" +
	"Resource_Service                                  : RESOURCE-SVC          " + "\n" +
	"Resource_Deployment                               : RESOURCE-DPL          " + "\n" +
	"Resource_Event                                    : RESOURCE-EVNT         " + "\n" +
	"Resource_Resource                                 : RESOURCE-RSRC         " + "\n" +
	"Resource_Namespace                                : RESOURCE-NSPC         " + "\n" +
	"Resource_Ingress                                  : RESOURCE-INGR         " + "\n" +
	"Resource_NodePort                                 : RESOURCE-NDPORT       " + "\n" +
	"Resource_PodScheduled                             : RESOURCE-PSCH         " + "\n" +
	"Resource_PodUnscheduled                           : RESOURCE-PUNSCH       " + "\n" +
	"Resource_ContainerCPU                             : RESOURCE-CCPU         " + "\n" +
	"Resource_ContainerMemory                          : RESOURCE-CMEM         " + "\n" +
	"Resource_ContainerFSRead                          : RESOURCE-CFSR         " + "\n" +
	"Resource_ContainerFSWrite                         : RESOURCE-CFSW         " + "\n" +
	"Resource_ContainerNetReceive                      : RESOURCE-CNETR        " + "\n" +
	"Resource_ContainerNetTransmit                     : RESOURCE-CNETT        " + "\n" +
	"Resource_VolumeAvailable                          : RESOURCE-VOLAVAIL     " + "\n" +
	"Resource_VolumeCapacity                           : RESOURCE-VOLCAP       " + "\n" +
	"Resource_VolumeUsed                               : RESOURCE-VOLUSD       " + "\n" +
	"Resource_NodeTemperature                          : RESOURCE-NTEMP        " + "\n" +
	"Resource_NodeTemperatureChange                    : RESOURCE-NTEMPCH      " + "\n" +
	"Resource_NodeTemperatureAverage                   : RESOURCE-NTEMPAV      " + "\n" +
	"Resource_NodeProcesses                            : RESOURCE-NPROCS       " + "\n" +
	"Resource_NodeCores                                : RESOURCE-NCORES       " + "\n" +
	"Resource_NodeMemory                               : RESOURCE-NMEM         " + "\n" +
	"Resource_NodeMemoryTotal                          : RESOURCE-NMEMTOT      " + "\n" +
	"Resource_NodeDiskRead                             : RESOURCE-NDISKR       " + "\n" +
	"Resource_NodeDiskWrite                            : RESOURCE-NDISKW       " + "\n" +
	"Resource_NodeNetworkReceive                       : RESOURCE-NNETR        " + "\n" +
	"Resource_NodeNetworkTransmit                      : RESOURCE-NNETT        " + "\n" +
	"Resource_NodeDiskWritten                          : RESOURCE-NDISKWT      " + "\n" +
	"Apply_RegisterSecret                              : APPLY-REGSEC          " + "\n" +
	"Apply_Distribute                                  : APPLY-DIST            " + "\n" +
	"Apply_CreateOperationSource                       : APPLY-CRTOPSSRC       " + "\n" +
	"Apply_Restart                                     : APPLY-RESTART         " + "\n" +
	"Apply_Rollback                                    : APPLY-ROLLBACK        " + "\n" +
	"Apply_Kill                                        : APPLY-KILL            " + "\n" +
	"Apply_NetworkRefresh                              : APPLY-NETRESH         " + "\n" +
	"Apply_HorizontalAutoscale                         : APPLY-HPA             " + "\n" +
	"Apply_HorizontalAutoscaleUndo                     : APPLY-HPAUN           " + "\n" +
	"Apply_QoS                                         : APPLY-QOS             " + "\n" +
	"Apply_QoSUndo                                     : APPLY-QOSUN           " + "\n" +
	"Apply_Ingress                                     : APPLY-INGR            " + "\n" +
	"Apply_IngressUndo                                 : APPLY-INGRUN          " + "\n" +
	"Apply_NodePort                                    : APPLY-NDPORT          " + "\n" +
	"Apply_NodePortUndo                                : APPLY-NDPORTUN        " + "\n" +
	""

func _CONSTRUCT_API_X() API_X {

	api_x := make(API_X)

	sanitized_def := strings.ReplaceAll(APIX_DEFINITION, " ", "")

	def_list := strings.Split(sanitized_def, "\n")

	for i := 0; i < len(def_list); i++ {

		if def_list[i] == "" || def_list[i] == " " || def_list[i] == "\n" {
			continue
		}

		raw_record := def_list[i]

		record_list := strings.SplitN(raw_record, ":", 2)

		key := record_list[0]

		api_x[key] = record_list[1]

	}

	return api_x
}

var AXgi = _CONSTRUCT_API_X()
