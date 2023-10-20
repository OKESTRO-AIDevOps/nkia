package main

import (
	"github.com/OKESTRO-AIDevOps/nkia/orch.apix/apix"
)

func main() {

	test := apix.Admin_InstallAnotherControlPlaneOnRemote{}

	apix.AXgi.ConvertToOrchIOSpec_Test(test)

}
