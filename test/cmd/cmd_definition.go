package cmd

import "strings"

var NKIACI_FLAGS = "" +

	"repo     : repository address" + "\n" +
	"id       : repository id, if" + "\n" +
	"token    : repository password (or token)," + "\n" +
	"name     : repository name for using as a workspace " + "\n" +
	"from-yaml: repo, id, token, name from yaml" + "\n" +
	""

type NKIA_CI_OPTIONS map[string]string

func _CONSTRUCT_NKIA_CI_FLAG_OPTIONS() NKIA_CI_OPTIONS {

	nkiaci_opts := make(NKIA_CI_OPTIONS)

	def_list := strings.Split(NKIACI_FLAGS, "\n")

	for i := 0; i < len(def_list); i++ {

		if def_list[i] == "" || def_list[i] == " " || def_list[i] == "\n" {
			continue
		}

		raw_record := def_list[i]

		record_list := strings.SplitN(raw_record, ":", 2)

		key := record_list[0]

		key = strings.ReplaceAll(key, " ", "")

		nkiaci_opts[key] = record_list[1]

	}

	return nkiaci_opts
}

var NKIACIopts = _CONSTRUCT_NKIA_CI_FLAG_OPTIONS()
