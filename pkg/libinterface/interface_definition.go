package libinterface

import (
	"fmt"
	"os"
	"path/filepath"
)

type LibIface map[string]string

func ConstructLibIface() (LibIface, error) {

	tmp_libif := make(LibIface)

	_ROOT, err := filepath.Abs("lib")

	if err != nil {

		return tmp_libif, fmt.Errorf("construct iface failed: %s", err.Error())

	}

	_DIR, err := os.Open(_ROOT)

	if err != nil {

		return tmp_libif, fmt.Errorf("construct iface failed: %s", err.Error())

	}

	_FILES, err := _DIR.Readdir(-1)

	if err != nil {

		return tmp_libif, fmt.Errorf("construct iface failed: %s", err.Error())

	}

	for _, f := range _FILES {

		iface_name := f.Name()

		full_path := filepath.Join(_ROOT, iface_name)

		tmp_libif[iface_name] = full_path
	}

	_DIR.Close()

	return tmp_libif, nil
}
