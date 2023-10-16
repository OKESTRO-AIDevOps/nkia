package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"

	nkctlclient "github.com/OKESTRO-AIDevOps/nkia/nokubectl/client"
)

func RunClientInteractive() {

	var email string

	var in_raw_query string

	var err error

	jar, err := cookiejar.New(nil)

	if err != nil {

		fmt.Println(err.Error())

		return

	}

	fmt.Println("email: ")

	fmt.Scanln(&email)

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

	if len(os.Args) <= 1 {
		fmt.Println("error: no args specified")
		return
	}

	MODE_INTERACTIVE := 0

	// MODE_ADMIN := 0

	for i := 0; i < len(os.Args); i++ {

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
