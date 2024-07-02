package apistandard

import "strings"

type API_OUTPUT struct {
	HEAD API_METADATA

	BODY string
}

type API_METADATA struct {
	JSON bool `json:"json"`
}

type API_INPUT map[string]string

type API_STD map[string][]string

var API_DEFINITION string = "" +
	/*a00*/ "NKCTL-INIT               :id                                                                                         " + "\n" +
	/*a01*/ "NKCTL-SETTO              :id, to                                                                                     " + "\n" +
	/*a02*/ "NKCTL-SETAS              :id, as                                                                                     " + "\n" +
	/*b00*/ "ORCH-CONNCHK             :id                                                                                         " + "\n" +
	/*b01*/ "ORCH-KEYGEN              :id                                                                                         " + "\n" +
	/*b02*/ "ORCH-GETCL               :id                                                                                         " + "\n" +
	/*b03*/ "ORCH-ADDCL               :id, clusterid                                                                              " + "\n" +
	/*b04*/ "ORCH-INSTCL              :id, clusterid, targetip, targetid, targetpw, localip, osnm, cv, updatetoken                " + "\n" +
	/*b05*/ "ORCH-INSTCLLOG           :id, clusterid, targetip, targetid, targetpw                                                " + "\n" +
	/*c00*/ "NKLET-CONNUP             :id, clusterid, updatetoken                                                                 " + "\n" +
	/*c01*/ "NKLET-CONN               :id, clusterid                                                                              " + "\n" +
	/*d00*/ "NKADM-INSTCTRL           :id, localip, osnm, cv                                                                      " + "\n" +
	/*d01*/ "NKADM-INSTWKOL           :id, localip, osnm, cv, token                                                               " + "\n" +
	/*f00*/ "ADMIN-INSTWKOR           :id, targetip, targetid, targetpw, localip, osnm, cv, token                                 " + "\n" +
	/*d02*/ "NKADM-INSTVOLOL          :id, localip                                                                                " + "\n" +
	/*f01*/ "ADMIN-INSTVOLOR          :id, targetip, targetid, targetpw, localip                                                  " + "\n" +
	/*d03*/ "NKADM-INSTTKOL           :id                                                                                         " + "\n" +
	/*f02*/ "ADMIN-INSTTKOR           :id, targetip, targetid, targetpw                                                           " + "\n" +
	/*d04*/ "NKADM-INSTLOGOL          :id                                                                                         " + "\n" +
	/*f03*/ "ADMIN-INSTLOGOR          :id, targetip, targetid, targetpw                                                           " + "\n" +
	/*f04*/ "ADMIN-INIT               :id                                                                                         " + "\n" +
	/*f05*/ "ADMIN-INITLOG            :id                                                                                         " + "\n" +
	/*g00*/ "SETTING-CRTNS            :id, ns, repoaddr, regaddr                                                                  " + "\n" +
	/*g01*/ "SETTING-SETREPO          :id, ns, repoaddr, repoid, repopw                                                           " + "\n" +
	/*g02*/ "SETTING-SETREG           :id, ns, regaddr, regid, regpw                                                              " + "\n" +
	/*g03*/ "SETTING-CRTVOL           :id, ns, targetip                                                                           " + "\n" +
	/*g04*/ "SETTING-CRTMON           :id                                                                                         " + "\n" +
	/*g05*/ "SETTING-CRTMONPERS       :id                                                                                         " + "\n" +
	/*h00*/ "TOOLKIT-BUILD            :id, ns, repoaddr, regaddr                                                                  " + "\n" +
	/*h01*/ "TOOLKIT-BUILDLOG         :id                                                                                         " + "\n" +
	/*h02*/ "TOOLKIT-PIPE             :id, ns, repoaddr, regaddr                                                                  " + "\n" +
	/*h03*/ "TOOLKIT-PIPELOG          :id                                                                                         " + "\n" +
	/*h04*/ "TOOLKIT-PIPESETVAR       :id, varnm, varval                                                                             " + "\n" +
	/*h05*/ "TOOLKIT-PIPEGETVAR       :id                                                                                         " + "\n" +
	/*i00*/ "RESOURCE-NDS             :id, ns                                                                                     " + "\n" +
	/*i01*/ "RESOURCE-PDS             :id, ns                                                                                     " + "\n" +
	/*i02*/ "RESOURCE-PLOG            :id, ns, podnm                                                                              " + "\n" +
	/*i03*/ "RESOURCE-SVC             :id, ns                                                                                     " + "\n" +
	/*i04*/ "RESOURCE-DPL             :id, ns                                                                                     " + "\n" +
	/*i05*/ "RESOURCE-EVNT            :id, ns                                                                                     " + "\n" +
	/*i06*/ "RESOURCE-RSRC            :id, ns                                                                                     " + "\n" +
	/*i07*/ "RESOURCE-NSPC            :id, ns                                                                                     " + "\n" +
	/*i08*/ "RESOURCE-INGR            :id, ns                                                                                     " + "\n" +
	/*i09*/ "RESOURCE-NDPORT          :id, ns                                                                                     " + "\n" +
	/*i10*/ "RESOURCE-PSCH            :id, ns                                                                                     " + "\n" +
	/*i11*/ "RESOURCE-PUNSCH          :id, ns                                                                                     " + "\n" +
	/*i12*/ "RESOURCE-CCPU            :id, ns                                                                                     " + "\n" +
	/*i13*/ "RESOURCE-CMEM            :id, ns                                                                                     " + "\n" +
	/*i14*/ "RESOURCE-CFSR            :id, ns                                                                                     " + "\n" +
	/*i15*/ "RESOURCE-CFSW            :id, ns                                                                                     " + "\n" +
	/*i16*/ "RESOURCE-CNETR           :id, ns                                                                                     " + "\n" +
	/*i17*/ "RESOURCE-CNETT           :id, ns                                                                                     " + "\n" +
	/*i18*/ "RESOURCE-VOLAVAIL        :id                                                                                         " + "\n" +
	/*i19*/ "RESOURCE-VOLCAP          :id                                                                                         " + "\n" +
	/*i20*/ "RESOURCE-VOLUSD          :id                                                                                         " + "\n" +
	/*i21*/ "RESOURCE-NTEMP           :id                                                                                         " + "\n" +
	/*i22*/ "RESOURCE-NTEMPCH         :id                                                                                         " + "\n" +
	/*i23*/ "RESOURCE-NTEMPAV         :id                                                                                         " + "\n" +
	/*i24*/ "RESOURCE-NPROCS          :id                                                                                         " + "\n" +
	/*i25*/ "RESOURCE-NCORES          :id                                                                                         " + "\n" +
	/*i26*/ "RESOURCE-NMEM            :id                                                                                         " + "\n" +
	/*i27*/ "RESOURCE-NMEMTOT         :id                                                                                         " + "\n" +
	/*i28*/ "RESOURCE-NDISKR          :id                                                                                         " + "\n" +
	/*i29*/ "RESOURCE-NDISKW          :id                                                                                         " + "\n" +
	/*i30*/ "RESOURCE-NNETR           :id                                                                                         " + "\n" +
	/*i31*/ "RESOURCE-NNETT           :id                                                                                         " + "\n" +
	/*i32*/ "RESOURCE-NDISKWT         :id                                                                                         " + "\n" +
	/*j00*/ "APPLY-REGSEC             :id, ns                                                                                     " + "\n" +
	/*j01*/ "APPLY-DIST               :id, ns, repoaddr, regaddr                                                                  " + "\n" +
	/*j02*/ "APPLY-CRTOPSSRC          :id, ns, repoaddr, regaddr                                                                  " + "\n" +
	/*j03*/ "APPLY-RESTART            :id, ns, resource, resourcenm                                                               " + "\n" +
	/*j04*/ "APPLY-ROLLBACK           :id, ns, resource, resourcenm                                                               " + "\n" +
	/*j05*/ "APPLY-KILL               :id, ns, resource, resourcenm                                                               " + "\n" +
	/*j06*/ "APPLY-NETRESH            :id                                                                                         " + "\n" +
	/*j07*/ "APPLY-HPA                :id, ns, resource, resourcenm                                                               " + "\n" +
	/*j08*/ "APPLY-HPAUN              :id, ns, resource, resourcenm                                                               " + "\n" +
	/*j09*/ "APPLY-QOS                :id, ns, resource, resourcenm                                                               " + "\n" +
	/*j10*/ "APPLY-QOSUN              :id, ns, resource, resourcenm                                                               " + "\n" +
	/*j11*/ "APPLY-INGR               :id, ns, hostnm, svcnm                                                                      " + "\n" +
	/*j12*/ "APPLY-INGRUN             :id, ns, hostnm, svcnm                                                                      " + "\n" +
	/*j13*/ "APPLY-NDPORT             :id, ns, svcnm                                                                              " + "\n" +
	/*j14*/ "APPLY-NDPORTUN           :id, ns, svcnm                                                                              " + "\n" +
	//"EXIT                     :id                                                           "
	""

func _CONSTRUCT_API_STD() API_STD {

	apistd := make(API_STD)

	sanitized_def := strings.ReplaceAll(API_DEFINITION, " ", "")

	def_list := strings.Split(sanitized_def, "\n")

	for i := 0; i < len(def_list); i++ {

		if def_list[i] == "" || def_list[i] == " " || def_list[i] == "\n" {
			continue
		}

		raw_record := def_list[i]

		record_list := strings.SplitN(raw_record, ":", 2)

		value_list := strings.Split(record_list[1], ",")

		key := record_list[0]

		apistd[key] = value_list

	}

	return apistd

}

var ASgi = _CONSTRUCT_API_STD()
