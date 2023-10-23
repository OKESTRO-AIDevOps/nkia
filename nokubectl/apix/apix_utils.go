package apix

import (
	"os"
	"strings"

	apistd "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
)

var NN = "\n\n"

var NL = "\\\n"

var SS30 = "                              "

func PrintHelp() {

}

func ExportMD() {

	var OUTPUT_MD string

	OUTPUT_MD += "" +

		"# NKIA API eXpression" + "\n\n" +

		"**Below are all available queries and corresponding required options**" + NN +

		"queries are made of arguments and then joined by the options" + NL +
		"ex) nokubectl {arg1} {arg2} {arg3} {--option_name} {option_val}" + NN +

		""

	query_and_options := QueryAndOptions()

	OUTPUT_MD += query_and_options

	_ = os.WriteFile("README.md", []byte(OUTPUT_MD), 0644)

}

func QueryAndOptions() string {

	qa := ""

	for i := 0; i < len(AXid); i++ {

		query := AXid[i]

		qid := AXgi[query]

		options := apistd.ASgi[qid]

		per_query := strings.ReplaceAll(query, "-", " ") + ": " + AXcmd[query] + NN

		for j := 0; j < len(options); j++ {

			oid := options[j]

			per_query += SS30 + options[j] + ": " + AXflag[oid] + NL

		}

		qa += per_query

	}

	return qa

}

func ExportJS() {

}

func ExportPY() {

}
