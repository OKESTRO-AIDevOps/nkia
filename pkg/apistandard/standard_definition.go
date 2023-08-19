package apistandard

import "strings"

type API_OUTPUT struct {
	HEAD API_METADATA

	BODY string
}

type API_METADATA map[string]map[string]string

type API_INPUT map[string]string

type API_STD map[string][]string

var API_DEFINITION string = "" +
	//            id          :       keys
	//	"ADMIN-ADMRMTCHK          :id, hostaddr, hostport, usernm, userpw, machinerole, acd     " + "\n" +
	//	"ADMIN-ADMRMTLDHA         :id, hostaddr, hostport, usernm, userpw, remoteip             " + "\n" +
	//	"ADMIN-ADMRMTLDMV         :id, hostaddr, hostport, usernm, userpw                       " + "\n" +
	//	"ADMIN-ADMRMTMSR          :id, hostaddr, hostport, usernm, userpw, localip, token       " + "\n" +
	//	"ADMIN-ADMRMTLDWRK        :id, hostaddr, hostport, usernm, userpw, localip, token       " + "\n" +
	//	"ADMIN-ADMRMTWRK          :id, hostaddr, hostport, usernm, userpw, localip, token       " + "\n" +
	//	"ADMIN-ADMRMTSTR          :id, hostaddr, hostport, usernm, userpw, localip, token       " + "\n" +
	//	"ADMIN-ADMRMTLOG          :id, hostaddr, hostport, usernm, userpw                       " + "\n" +
	//	"ADMIN-ADMRMTSTATUS       :id, hostaddr, hostport, usernm, userpw                       " + "\n" +
	//	"ADMIN-LEAD               :id, token, targetip                                          " + "\n" +
	//	"ADMIN-MSR                :id, token, targetip                                          " + "\n" +
	//	"ADMIN-LDVOL              :id, token, targetip                                          " + "\n" +
	//	"ADMIN-WRK                :id, token, targetip                                          " + "\n" +
	//	"ADMIN-STR                :id, token, targetip                                          " + "\n" +
	//	"ADMIN-LOG                :id                                                           " + "\n" +
	//	"ADMIN-STATUS             :id, token, targetip                                          " + "\n" +
	//	"ADMIN-UP                 :id, token                                                    " + "\n" +
	//	"ADMIN-DOWN               :id, token                                                    " + "\n" +
	//	"DELND                    :id                                                           " + "\n" +
	"ADMIN-INIT               :id                                                           " + "\n" +
	"ADMIN-INITLOG            :id                                                           " + "\n" +
	"SETTING-CRTNS            :id, ns, repoaddr, regaddr                                    " + "\n" +
	"SETTING-SETREPO          :id, ns, repoaddr, repoid, repopw                             " + "\n" +
	"SETTING-SETREG           :id, ns, regaddr, regid, regpw                                " + "\n" +
	// "SETTING-CRTNSVOL          :id, ns, volserver                                            " + "\n" +
	// "SETTING-CRTVOL            :id, ns, volserver                                            " + "\n" +
	"SETTING-CRTMON           :id                                                           " + "\n" +
	//  "SETTING-DELNS            :id, ns                                                       " + "\n" +
	//  "SUBMIT                   :id                                                           " + "\n" +
	//	"CALLME                   :id                                                           " + "\n" +
	//  "GITLOG                   :id, ns, repoaddr                                             " + "\n" +
	//	"PIPEHIST                 :id, ns                                                       " + "\n" +
	//	"PIPE                     :id, ns, repoaddr, regaddr                                    " + "\n" +
	//	"PIPELOG                  :id                                                           " + "\n" +
	"BUILD                    :id, ns, repoaddr, regaddr                                    " + "\n" +
	"BUILDLOG                 :id                                                           " + "\n" +
	"RESOURCE-NDS             :id, ns                                                       " + "\n" +
	"RESOURCE-PDS             :id, ns                                                       " + "\n" +
	"RESOURCE-PLOG            :id, ns, podnm                                                " + "\n" +
	"RESOURCE-SVC             :id, ns                                                       " + "\n" +
	"RESOURCE-DPL             :id, ns                                                       " + "\n" +
	//	"RESOURCE-IMGLI           :id, ns                                                       " + "\n" +
	"RESOURCE-EVNT            :id, ns                                                       " + "\n" +
	"RESOURCE-RSRC            :id, ns                                                       " + "\n" +
	"RESOURCE-NSPC            :id, ns                                                       " + "\n" +
	"RESOURCE-INGR            :id, ns                                                       " + "\n" +
	"RESOURCE-NDPORT          :id, ns                                                       " + "\n" +
	//	"RESOURCE-PRJPRB          :id, ns                                                       " + "\n" +
	"RESOURCE-PSCH            :id, ns                                                       " + "\n" +
	"RESOURCE-PUNSCH          :id, ns                                                       " + "\n" +
	"RESOURCE-CCPU            :id, ns                                                       " + "\n" +
	"RESOURCE-CMEM            :id, ns                                                       " + "\n" +
	"RESOURCE-CFSR            :id, ns                                                       " + "\n" +
	"RESOURCE-CFSW            :id, ns                                                       " + "\n" +
	"RESOURCE-CNETR           :id, ns                                                       " + "\n" +
	"RESOURCE-CNETT           :id, ns                                                       " + "\n" +
	"RESOURCE-VOLAVAIL        :id                                                           " + "\n" +
	"RESOURCE-VOLCAP          :id                                                           " + "\n" +
	"RESOURCE-VOLUSD          :id                                                           " + "\n" +
	"RESOURCE-NTEMP           :id                                                           " + "\n" +
	"RESOURCE-NTEMPCH         :id                                                           " + "\n" +
	"RESOURCE-NTEMPAV         :id                                                           " + "\n" +
	"RESOURCE-NPROCS          :id                                                           " + "\n" +
	"RESOURCE-NCORES          :id                                                           " + "\n" +
	"RESOURCE-NMEM            :id                                                           " + "\n" +
	"RESOURCE-NMEMTOT         :id                                                           " + "\n" +
	"RESOURCE-NDISKR          :id                                                           " + "\n" +
	"RESOURCE-NDISKW          :id                                                           " + "\n" +
	"RESOURCE-NNETR           :id                                                           " + "\n" +
	"RESOURCE-NNETT           :id                                                           " + "\n" +
	"RESOURCE-NDISKWT         :id                                                           " + "\n" +
	"APPLY-REGSEC             :id, ns                                                       " + "\n" +
	"APPLY-DIST               :id, ns, repoaddr, regaddr                                    " + "\n" +
	"APPLY-CRTOPSSRC          :id, ns, repoaddr, regaddr                                    " + "\n" +
	"APPLY-RESTART            :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-ROLLBACK           :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-KILL               :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-NETRESH            :id                                                           " + "\n" +
	"APPLY-HPA                :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-HPAUN              :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-QOS                :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-QOSUN              :id, ns, resource, resourcenm                                 " + "\n" +
	"APPLY-INGR               :id, ns, hostnm, svcnm                                        " + "\n" +
	"APPLY-INGRUN             :id, ns, hostnm, svcnm                                        " + "\n" +
	"APPLY-NDPORT             :id, ns, svcnm                                                " + "\n" +
	"APPLY-NDPORTUN           :id, ns, svcnm                                                "
	//"EXIT                     :id                                                           "

func _CONSTRUCT_API_INPUT() API_STD {

	apistd := make(API_STD)

	sanitized_def := strings.ReplaceAll(API_DEFINITION, " ", "")

	def_list := strings.Split(sanitized_def, "\n")

	for i := 0; i < len(def_list); i++ {

		raw_record := def_list[i]

		record_list := strings.SplitN(raw_record, ":", 2)

		value_list := strings.Split(record_list[1], ",")

		key := record_list[0]

		apistd[key] = value_list

	}

	return apistd

}

var ASgi = _CONSTRUCT_API_INPUT()
