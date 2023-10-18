package main

import (
	"os"

	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/modules"
)

func main() {

	server_sym_key, err := modules.RandomHex(16)

	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile("okey", []byte(server_sym_key), 0644)

	if err != nil {
		panic(err.Error())
	}

}
