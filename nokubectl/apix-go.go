package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"github.com/OKESTRO-AIDevOps/nkia/nokubectl/apix"
	nkctlclient "github.com/OKESTRO-AIDevOps/nkia/nokubectl/client"
	"github.com/OKESTRO-AIDevOps/nkia/nokubectl/config"
)

func RunClientX(apix_id string, apix_options apix.API_X_OPTIONS) {

	var email string

	var err error

	jar, err := cookiejar.New(nil)

	if err != nil {

		fmt.Println(err.Error())

		return

	}

	email = config.EMAIL

	client := &http.Client{
		Jar: jar,
	}

	c, err := nkctlclient.KeyAuthConn(client, email)

	if err != nil {

		fmt.Println(err.Error())

		return
	}

	oreq, err := apix.AXgi.BuildOrchRequest(apix_id, apix_options)

	if err != nil {

		fmt.Printf("failed: %s\n", err.Error())

		return
	}

	nkctlclient.RequestHandler_APIX_Once_PrintOnly(c, oreq)

}
