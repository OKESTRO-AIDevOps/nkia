package libinterface

import (
	"fmt"
	"path/filepath"
	//. "github.com/goccy/go-yaml"
)

func (libif LibIface) GetLibComponentAddress(iface_name string, component string) (string, error) {

	iface_path, okay := libif[iface_name]

	if !okay {

		return "", fmt.Errorf("failed to get address: %s", "no such iface")
	}

	iface_component_path := filepath.Join(iface_path, component)

	return iface_component_path, nil

}
