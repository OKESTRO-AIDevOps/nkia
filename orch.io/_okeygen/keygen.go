package main

import (
	"fmt"
	"os"

	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
)

func ToFile(cs *modules.CertSet) error {

	err := os.WriteFile("certs_server/ca.crt", cs.RootCertPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())

	}

	err = os.WriteFile("certs_server/ca.priv", cs.RootKeyPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())
	}

	err = os.WriteFile("certs_server/ca.pub", cs.RootPubPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())
	}

	err = os.WriteFile("certs/ca.crt", cs.RootCertPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())

	}

	err = os.WriteFile("certs/ca.priv", cs.RootKeyPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())
	}

	err = os.WriteFile("certs/ca.pub", cs.RootPubPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())
	}

	err = os.WriteFile("certs_server/server.crt", cs.ServCertPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())
	}

	err = os.WriteFile("certs_server/server.priv", cs.ServKeyPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())
	}

	err = os.WriteFile("certs_server/server.pub", cs.ServPubPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())
	}

	err = os.WriteFile("certs/client.crt", cs.ClientCertPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())
	}

	err = os.WriteFile("certs/client.priv", cs.ClientKeyPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())
	}

	err = os.WriteFile("certs/client.pub", cs.ClientPubPEM, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())
	}

	return nil
}

func EncryptExport() {

}

func main() {

	var cs *modules.CertSet

	cs = modules.NewCertsPipeline()

	err := ToFile(cs)

	if err != nil {

		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}

}
