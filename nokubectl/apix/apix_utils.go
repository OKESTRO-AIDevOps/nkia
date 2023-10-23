package apix

import (
	"os"
	"strings"

	apistd "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
)

var NN = "\n\n"

var N = "\\\n"

var SS8 = strings.Repeat("&nbsp;", 8)

var SS16 = strings.Repeat("&nbsp;", 16)

func PrintHelp() {

}

func ExportMD() {

	var OUTPUT_MD string

	OUTPUT_MD += "" +

		"# NKIA API eXpression" + NN +

		"**Below are nokubectl specific flags" + NN +

		""

	nkctl_flags := NKCTLFlags()

	OUTPUT_MD += nkctl_flags

	OUTPUT_MD += "" +

		"**Below are all available queries and corresponding required options**" + NN +

		"queries are made of arguments and then joined by the options" + N +
		"ex) nokubectl {arg1} {arg2} {arg3} {--option_name} {option_val}" + NN +

		""

	query_and_options := QueryAndOptions()

	OUTPUT_MD += query_and_options

	_ = os.WriteFile("README.md", []byte(OUTPUT_MD), 0644)

}

func NKCTLFlags() string {

	nkctl_flag := ""

	for k, v := range NKCTLflag {

		per_flag := "- " + k + ": " + v + N

		nkctl_flag += per_flag

	}

	return nkctl_flag

}

func QueryAndOptions() string {

	qa := ""

	for i := 0; i < len(AXid); i++ {

		query := AXid[i]

		qid := AXgi[query]

		options := apistd.ASgi[qid]

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

		qa += per_query + NN

	}

	return qa

}

func ExportJS() {

}

func ExportPY() {

}
