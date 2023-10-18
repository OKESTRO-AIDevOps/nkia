package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"

	nkctlclient "github.com/OKESTRO-AIDevOps/nkia/nokubectl/client"
)

func InitCtl() error {

	var priv_loc string

	var file_b []byte

	cmd := exec.Command("mkdir", "-p", "srv")

	err := cmd.Run()

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	fmt.Println("enter privkey filepath: ")

	fmt.Scanln(&priv_loc)

	file_b, err = os.ReadFile(priv_loc)

	if err != nil {
		return fmt.Errorf("failed to init: %s", err.Error())
	}

	err = os.WriteFile("srv/.priv", file_b, 0644)

	if err != nil {
		return fmt.Errorf("failed to init: %s", err.Error())
	}

	return nil
}

func RunClientInteractive() {

	var email string

	var in_raw_query string

	var err error

	jar, err := cookiejar.New(nil)

	if err != nil {

		fmt.Println(err.Error())

		return

	}

	email = nkctlclient.EMAIL

	client := &http.Client{
		Jar: jar,
	}

	c, err := nkctlclient.KeyAuthConn(client, email)

	if err != nil {

		fmt.Println(err.Error())

		return
	}

	for {

		fmt.Println("query: ")

		fmt.Scanln(&in_raw_query)

		switch in_raw_query {

		case "exit":

			return

		default:

			var target string

			var option string

			fmt.Println("target: ")

			fmt.Scanln(&target)

			fmt.Println("option: ")

			fmt.Scanln(&option)

			nkctlclient.RequestHandler_LinearInstruction_PrintOnly(c, target, option, in_raw_query)

		}

	}

}

func RunClient() {

}

func main() {

	err := InitCtl()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(os.Args) <= 1 {
		fmt.Println("error: no args specified")
		return
	}

	MODE_INTERACTIVE := 0

	// MODE_ADMIN := 0

	for i := 1; i < len(os.Args); i++ {

		flag := os.Args[i]

		if flag == "-i" || flag == "--interactive" {

			MODE_INTERACTIVE = 1

		}

	}

	//if (MODE_INTERACTIVE) > 1 {
	//	fmt.Println("error: more than one option used together")
	//	return
	//}

	if MODE_INTERACTIVE == 1 {

		RunClientInteractive()

	} else {

		RunClient()

	}

	return
}
