package apix

import (
	"os"
	"strings"

	apistd "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
)

var NN = "\n\n\n\n"

var N = " \n\n"

var SS8 = strings.Repeat(" ", 8)

var SS16 = strings.Repeat(" ", 16)

var STARTBLOCK = "\n\n```yaml\n"

var ENDBLOCK = "\n```\n\n"

func PrintHelp() {

}

func ExportMD() {

	var OUTPUT_MD string

	OUTPUT_MD += "" +

		"# NKIA API eXpression" + NN +

		"## Below are nokubectl specific flags" + NN +

		"**nokudectl specific flags cannot be used in combination with each other or other options,**" + N +
		"**aka mutually exclusive**" + NN +

		"**ex) nokubectl {flag}**" + NN +

		""

	nkctl_flags := NKCTLFlags()

	OUTPUT_MD += STARTBLOCK + nkctl_flags + ENDBLOCK

	OUTPUT_MD += "" +

		"## Below are orch.io request specific options**" + NN +

		"**ex) nokubectl {arg1} {arg2} {arg3} {--req_option_name} {req_option_val}**" + NN +

		""

	orch_req_flag := OrchRequestFlags()

	OUTPUT_MD += STARTBLOCK + orch_req_flag + ENDBLOCK

	OUTPUT_MD += "" +

		"## Below are all available queries and corresponding required options" + NN +

		"**queries are made of arguments and then joined by the options**" + N +
		"**ex) nokubectl {arg1} {arg2} {arg3} {--option_name} {option_val}**" + NN +

		""

	query_and_options := QueryAndOptions()

	OUTPUT_MD += STARTBLOCK + query_and_options + ENDBLOCK

	_ = os.WriteFile("README.md", []byte(OUTPUT_MD), 0644)

}

func NKCTLFlags() string {

	nkctl_flag := ""

	for k, v := range NKCTLflag {

		per_flag := "- " + k + ": " + N

		per_flag += SS8 + "description: " + v + N

		nkctl_flag += per_flag

	}

	nkctl_flag += NN + ""

	return nkctl_flag

}

func OrchRequestFlags() string {

	orchreq_flag := ""

	for k, v := range AXflag {

		if k == "to" {

			per_flag := "- " + k + ": " + N

			per_flag += SS8 + "description: " + v + N

			orchreq_flag += per_flag

		} else if k == "as" {

			per_flag := "- " + k + ": " + N

			per_flag += SS8 + "description: " + v + N

			orchreq_flag += per_flag

		} else {
			continue
		}

	}

	return orchreq_flag

}

func QueryAndOptions() string {

	qa := ""

	for i := 0; i < len(AXid); i++ {

		admin_flag := 0

		query := AXid[i]

		qid := AXgi[query]

		options := apistd.ASgi[qid]

		if !strings.Contains(query, "-") {
			admin_flag = 1
		}

		per_query := "- " + strings.ReplaceAll(query, "-", " ") + ": " + N

		per_query += SS8 + "description: " + AXcmd[query] + N

		per_query += SS8 + "options:" + N

		for j := 0; j < len(options); j++ {

			oid := options[j]

			if oid == "id" {
				continue
			}

			per_query += SS16 + "--" + options[j] + ": " + AXflag[oid] + N

		}

		if admin_flag == 1 {
			per_query += SS16 + "# this is admin query, which is used" + N
			per_query += SS16 + "# with [ --as admin ] " + N
		} else {
			per_query += SS16 + "# this is nkia query, which is used" + N
			per_query += SS16 + "# with [ --to $CLUSTER_ID ] usually " + N
		}

		qa += per_query + NN

	}

	return qa

}

func ExportJS() {

}

func ExportPY() {

}
