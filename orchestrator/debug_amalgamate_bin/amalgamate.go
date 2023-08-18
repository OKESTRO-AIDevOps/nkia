package main

import (
	"encoding/hex"
	"os"

	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
)

func main() {

	server_sym_key, err := modules.RandomHex(16)

	if err != nil {
		panic(err.Error())
	}

	file_byte, err := os.ReadFile("config.amld.raw")

	if err != nil {
		panic(err.Error())
	}

	enc_file_b, err := modules.EncryptWithSymmetricKey([]byte(server_sym_key), file_byte)

	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile("okey", []byte(server_sym_key), 0644)

	if err != nil {
		panic(err.Error())
	}

	enc_file_str := hex.EncodeToString(enc_file_b)

	err = os.WriteFile("config.amld", []byte(enc_file_str), 0644)

	if err != nil {
		panic(err.Error())
	}

}
