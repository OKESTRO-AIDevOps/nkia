package apix

import "strings"

type API_X map[string]string

type API_X_ID []string

type API_X_OPTIONS map[string]string

var APIX_QUERY_MAP = "" +
	/*a00*/ "init                                              : NKCTL-INIT            " + "\n" +
	/*a01*/ "set                                               : NKCTL-SETTO           " + "\n" +
	/*a02*/ "set-opts                                          : NKCTL-SETAS           " + "\n" +
	/*b00*/ "orch-conncheck                                    : ORCH-CONNCHK          " + "\n" +
	/*b01*/ "orch-keygen                                       : ORCH-KEYGEN           " + "\n" +
	/*b02*/ "orch-get-cl                                       : ORCH-GETCL            " + "\n" +
	/*b03*/ "orch-add-cl                                       : ORCH-ADDCL            " + "\n" +
	/*b04*/ "orch-install-cl                                   : ORCH-INSTCL           " + "\n" +
	/*b05*/ "orch-install-cl-log                               : ORCH-INSTCLLOG        " + "\n" +
	/*c00*/ "io-connect-update                                 : NKLET-CONNUP          " + "\n" +
	/*c01*/ "io-connect                                        : NKLET-CONN            " + "\n" +
	/*d00*/ "install-mainctrl                                  : NKADM-INSTCTRL        " + "\n" +
	/*d01*/ "install-worker                                    : NKADM-INSTWKOL        " + "\n" +
	/*f00*/ "admin-install-worker                              : ADMIN-INSTWKOR        " + "\n" +
	/*d02*/ "install-volume                                    : NKADM-INSTVOLOL       " + "\n" +
	/*f01*/ "admin-install-volume                              : ADMIN-INSTVOLOR       " + "\n" +
	/*d03*/ "install-toolkit                                   : NKADM-INSTTKOL        " + "\n" +
	/*f02*/ "admin-install-toolkit                             : ADMIN-INSTTKOR        " + "\n" +
	/*d04*/ "install-log                                       : NKADM-INSTLOGOL       " + "\n" +
	/*f03*/ "admin-install-log                                 : ADMIN-INSTLOGOR       " + "\n" +
	/*f04*/ "admin-init                                        : ADMIN-INIT            " + "\n" +
	/*f05*/ "admin-init-log                                    : ADMIN-INITLOG         " + "\n" +
	/*g00*/ "setting-create-namespace                          : SETTING-CRTNS         " + "\n" +
	/*g01*/ "setting-set-repo                                  : SETTING-SETREPO       " + "\n" +
	/*g02*/ "setting-set-reg                                   : SETTING-SETREG        " + "\n" +
	/*g03*/ "setting-create-volume                             : SETTING-CRTVOL        " + "\n" +
	/*g04*/ "setting-create-monitoring                         : SETTING-CRTMON        " + "\n" +
	/*g05*/ "setting-create-monitoring-persist                 : SETTING-CRTMONPERS    " + "\n" +
	/*h00*/ "toolkit-build                                     : TOOLKIT-BUILD         " + "\n" +
	/*h01*/ "toolkit-build-log                                 : TOOLKIT-BUILDLOG      " + "\n" +
	/*h02*/ "toolkit-pipe                                      : TOOLKIT-PIPE          " + "\n" +
	/*h03*/ "toolkit-pipe-log                                  : TOOLKIT-PIPELOG       " + "\n" +
	/*h04*/ "toolkit-pipe-set-var                              : TOOLKIT-PIPESETVAR    " + "\n" +
	/*h05*/ "toolkit-pipe-get-var                              : TOOLKIT-PIPEGETVAR    " + "\n" +
	/*i00*/ "resource-nodes                                    : RESOURCE-NDS          " + "\n" +
	/*i01*/ "resource-pods                                     : RESOURCE-PDS          " + "\n" +
	/*i02*/ "resource-pods-log                                 : RESOURCE-PLOG         " + "\n" +
	/*i03*/ "resource-service                                  : RESOURCE-SVC          " + "\n" +
	/*i04*/ "resource-deployment                               : RESOURCE-DPL          " + "\n" +
	/*i05*/ "resource-event                                    : RESOURCE-EVNT         " + "\n" +
	/*i06*/ "resource-resource                                 : RESOURCE-RSRC         " + "\n" +
	/*i07*/ "resource-namespace                                : RESOURCE-NSPC         " + "\n" +
	/*i08*/ "resource-ingress                                  : RESOURCE-INGR         " + "\n" +
	/*i09*/ "resource-nodeport                                 : RESOURCE-NDPORT       " + "\n" +
	/*i10*/ "resource-pod-scheduled                            : RESOURCE-PSCH         " + "\n" +
	/*i11*/ "resource-pod-unscheduled                          : RESOURCE-PUNSCH       " + "\n" +
	/*i12*/ "resource-container-cpu                            : RESOURCE-CCPU         " + "\n" +
	/*i13*/ "resource-container-mem                            : RESOURCE-CMEM         " + "\n" +
	/*i14*/ "resource-container-fs-read                        : RESOURCE-CFSR         " + "\n" +
	/*i15*/ "resource-container-fs-write                       : RESOURCE-CFSW         " + "\n" +
	/*i16*/ "resource-container-net-receive                    : RESOURCE-CNETR        " + "\n" +
	/*i17*/ "resource-container-net-transmit                   : RESOURCE-CNETT        " + "\n" +
	/*i18*/ "resource-volume-available                         : RESOURCE-VOLAVAIL     " + "\n" +
	/*i19*/ "resource-volume-capacity                          : RESOURCE-VOLCAP       " + "\n" +
	/*i20*/ "resource-volume-used                              : RESOURCE-VOLUSD       " + "\n" +
	/*i21*/ "resource-node-temperature                         : RESOURCE-NTEMP        " + "\n" +
	/*i22*/ "resource-node-temperature-change                  : RESOURCE-NTEMPCH      " + "\n" +
	/*i23*/ "resource-node-temperature-average                 : RESOURCE-NTEMPAV      " + "\n" +
	/*i24*/ "resource-node-process                             : RESOURCE-NPROCS       " + "\n" +
	/*i25*/ "resource-node-cores                               : RESOURCE-NCORES       " + "\n" +
	/*i26*/ "resource-node-mem                                 : RESOURCE-NMEM         " + "\n" +
	/*i27*/ "resource-node-mem-total                           : RESOURCE-NMEMTOT      " + "\n" +
	/*i28*/ "resource-node-disk-read                           : RESOURCE-NDISKR       " + "\n" +
	/*i29*/ "resource-node-disk-write                          : RESOURCE-NDISKW       " + "\n" +
	/*i30*/ "resource-node-net-receive                         : RESOURCE-NNETR        " + "\n" +
	/*i31*/ "resource-node-net-transmit                        : RESOURCE-NNETT        " + "\n" +
	/*i32*/ "resource-node-disk-written                        : RESOURCE-NDISKWT      " + "\n" +
	/*j00*/ "apply-reg-secret                                  : APPLY-REGSEC          " + "\n" +
	/*j01*/ "apply-distro                                      : APPLY-DIST            " + "\n" +
	/*j02*/ "apply-create-operation-source                     : APPLY-CRTOPSSRC       " + "\n" +
	/*j03*/ "apply-restart                                     : APPLY-RESTART         " + "\n" +
	/*j04*/ "apply-rollback                                    : APPLY-ROLLBACK        " + "\n" +
	/*j05*/ "apply-kill                                        : APPLY-KILL            " + "\n" +
	/*j06*/ "apply-net-refresh                                 : APPLY-NETRESH         " + "\n" +
	/*j07*/ "apply-hpa                                         : APPLY-HPA             " + "\n" +
	/*j08*/ "apply-hpa-undo                                    : APPLY-HPAUN           " + "\n" +
	/*j09*/ "apply-qos                                         : APPLY-QOS             " + "\n" +
	/*j10*/ "apply-qos-undo                                    : APPLY-QOSUN           " + "\n" +
	/*j11*/ "apply-ingress                                     : APPLY-INGR            " + "\n" +
	/*j12*/ "apply-ingress-undo                                : APPLY-INGRUN          " + "\n" +
	/*j13*/ "apply-nodeport                                    : APPLY-NDPORT          " + "\n" +
	/*j14*/ "apply-nodeport-undo                               : APPLY-NDPORTUN        " + "\n" +
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

var AXgi, AXid = _CONSTRUCT_API_X()

var AXcmd = _CONSTRUCT_API_X_COMMAND()

var AXflag = _CONSTRUCT_API_X_FLAG()
