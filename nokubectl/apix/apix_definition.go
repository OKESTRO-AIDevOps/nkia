package apix

import "strings"

type API_X map[string]string

type API_X_ID []string

type API_X_OPTIONS map[string]string

var APIX_QUERY_MAP = "" +
	// front admin option
	"conncheck                                         : CONNCHK               " + "\n" +
	"keygen                                            : KEYGEN                " + "\n" +
	"addcluster                                        : ADDCLUSTER            " + "\n" +
	// query
	"admin-install-env                                 : ADMIN-INSTENV         " + "\n" +
	"admin-install-mainctrl                            : ADMIN-INSTCTRL        " + "\n" +
	"admin-install-subctrl-prep                        : ADMIN-INSTANCTRLOL    " + "\n" +
	"admin-install-subctrl-add                         : ADMIN-INSTANCTRLOR    " + "\n" +
	"admin-install-worker                              : ADMIN-INSTWK          " + "\n" +
	"admin-install-volume-prep                         : ADMIN-INSTVOLOR       " + "\n" +
	"admin-install-volume-add                          : ADMIN-INSTVOLOL       " + "\n" +
	"admin-install-toolkit                             : ADMIN-INSTTK          " + "\n" +
	"admin-install-log                                 : ADMIN-INSTLOG         " + "\n" +
	"admin-install-lock-get                            : ADMIN-INSTLOCKGET     " + "\n" +
	"admin-install-lock-set                            : ADMIN-INSTLOCKSET     " + "\n" +
	"admin-init                                        : ADMIN-INIT            " + "\n" +
	"admin-init-log                                    : ADMIN-INITLOG         " + "\n" +
	"setting-create-namespace                          : SETTING-CRTNS         " + "\n" +
	"setting-set-repo                                  : SETTING-SETREPO       " + "\n" +
	"setting-set-reg                                   : SETTING-SETREG        " + "\n" +
	"setting-create-monitoring                         : SETTING-CRTMON        " + "\n" +
	"toolkit-build                                     : TOOLKIT-BUILD         " + "\n" +
	"toolkit-build-log                                 : TOOLKIT-BUILDLOG      " + "\n" +
	"resource-nodes                                    : RESOURCE-NDS          " + "\n" +
	"resource-pods                                     : RESOURCE-PDS          " + "\n" +
	"resource-pods-log                                 : RESOURCE-PLOG         " + "\n" +
	"resource-service                                  : RESOURCE-SVC          " + "\n" +
	"resource-deployment                               : RESOURCE-DPL          " + "\n" +
	"resource-event                                    : RESOURCE-EVNT         " + "\n" +
	"resource-resource                                 : RESOURCE-RSRC         " + "\n" +
	"resource-namespace                                : RESOURCE-NSPC         " + "\n" +
	"resource-ingress                                  : RESOURCE-INGR         " + "\n" +
	"resource-nodeport                                 : RESOURCE-NDPORT       " + "\n" +
	"resource-pod-scheduled                            : RESOURCE-PSCH         " + "\n" +
	"resource-pod-unscheduled                          : RESOURCE-PUNSCH       " + "\n" +
	"resource-container-cpu                            : RESOURCE-CCPU         " + "\n" +
	"resource-container-mem                            : RESOURCE-CMEM         " + "\n" +
	"resource-container-fs-read                        : RESOURCE-CFSR         " + "\n" +
	"resource-container-fs-write                       : RESOURCE-CFSW         " + "\n" +
	"resource-container-net-receive                    : RESOURCE-CNETR        " + "\n" +
	"resource-container-net-transmit                   : RESOURCE-CNETT        " + "\n" +
	"resource-volume-available                         : RESOURCE-VOLAVAIL     " + "\n" +
	"resource-volume-capacity                          : RESOURCE-VOLCAP       " + "\n" +
	"resource-volume-used                              : RESOURCE-VOLUSD       " + "\n" +
	"resource-node-temperature                         : RESOURCE-NTEMP        " + "\n" +
	"resource-node-temperature-change                  : RESOURCE-NTEMPCH      " + "\n" +
	"resource-node-temperature-average                 : RESOURCE-NTEMPAV      " + "\n" +
	"resource-node-process                             : RESOURCE-NPROCS       " + "\n" +
	"resource-node-cores                               : RESOURCE-NCORES       " + "\n" +
	"resource-node-mem                                 : RESOURCE-NMEM         " + "\n" +
	"resource-node-mem-total                           : RESOURCE-NMEMTOT      " + "\n" +
	"resource-node-disk-read                           : RESOURCE-NDISKR       " + "\n" +
	"resource-node-disk-write                          : RESOURCE-NDISKW       " + "\n" +
	"resource-node-net-receive                         : RESOURCE-NNETR        " + "\n" +
	"resource-node-net-transmit                        : RESOURCE-NNETT        " + "\n" +
	"resource-node-disk-written                        : RESOURCE-NDISKWT      " + "\n" +
	"apply-reg-secret                                  : APPLY-REGSEC          " + "\n" +
	"apply-distro                                      : APPLY-DIST            " + "\n" +
	"apply-create-operation-source                     : APPLY-CRTOPSSRC       " + "\n" +
	"apply-restart                                     : APPLY-RESTART         " + "\n" +
	"apply-rollback                                    : APPLY-ROLLBACK        " + "\n" +
	"apply-kill                                        : APPLY-KILL            " + "\n" +
	"apply-net-refresh                                 : APPLY-NETRESH         " + "\n" +
	"apply-hpa                                         : APPLY-HPA             " + "\n" +
	"apply-hpa-undo                                    : APPLY-HPAUN           " + "\n" +
	"apply-qos                                         : APPLY-QOS             " + "\n" +
	"apply-qos-undo                                    : APPLY-QOSUN           " + "\n" +
	"apply-ingress                                     : APPLY-INGR            " + "\n" +
	"apply-ingress-undo                                : APPLY-INGRUN          " + "\n" +
	"apply-nodeport                                    : APPLY-NDPORT          " + "\n" +
	"apply-nodeport-undo                               : APPLY-NDPORTUN        " + "\n" +
	""

func _CONSTRUCT_API_X() (API_X, API_X_ID) {

	api_x := make(API_X)

	api_x_id := make(API_X_ID, 0)

	sanitized_def := strings.ReplaceAll(APIX_QUERY_MAP, " ", "")

	def_list := strings.Split(sanitized_def, "\n")

	for i := 0; i < len(def_list); i++ {

		if def_list[i] == "" || def_list[i] == " " || def_list[i] == "\n" {
			continue
		}

		raw_record := def_list[i]

		record_list := strings.SplitN(raw_record, ":", 2)

		key := record_list[0]

		api_x[key] = record_list[1]

		api_x_id = append(api_x_id, key)

	}

	return api_x, api_x_id
}

func _CONSTRUCT_API_X_COMMAND() API_X {

	api_x := make(API_X)

	def_list := strings.Split(APIX_COMMAND, "\n")

	for i := 0; i < len(def_list); i++ {

		if def_list[i] == "" || def_list[i] == " " || def_list[i] == "\n" {
			continue
		}

		raw_record := def_list[i]

		record_list := strings.SplitN(raw_record, ":", 2)

		key := record_list[0]

		key = strings.ReplaceAll(key, " ", "")

		api_x[key] = record_list[1]

	}

	return api_x
}

func _CONSTRUCT_API_X_FLAG() API_X {

	api_x := make(API_X)

	def_list := strings.Split(APIX_FLAGS, "\n")

	for i := 0; i < len(def_list); i++ {

		if def_list[i] == "" || def_list[i] == " " || def_list[i] == "\n" {
			continue
		}

		raw_record := def_list[i]

		record_list := strings.SplitN(raw_record, ":", 2)

		key := record_list[0]

		key = strings.ReplaceAll(key, " ", "")

		api_x[key] = record_list[1]

	}

	return api_x
}

func _CONSTRUCT_NKCTL_FLAG() API_X {

	api_x := make(API_X)

	def_list := strings.Split(NKCTL_FLAGS, "\n")

	for i := 0; i < len(def_list); i++ {

		if def_list[i] == "" || def_list[i] == " " || def_list[i] == "\n" {
			continue
		}

		raw_record := def_list[i]

		record_list := strings.SplitN(raw_record, ":", 2)

		key := record_list[0]

		key = strings.ReplaceAll(key, " ", "")

		api_x[key] = record_list[1]

	}

	return api_x
}

var AXgi, AXid = _CONSTRUCT_API_X()

var AXcmd = _CONSTRUCT_API_X_COMMAND()

var AXflag = _CONSTRUCT_API_X_FLAG()

var NKCTLflag = _CONSTRUCT_NKCTL_FLAG()
