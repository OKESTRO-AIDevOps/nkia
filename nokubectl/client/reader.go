package client

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
)

func Read_APIX_Store_Override() {

	var HEAD apistandard.API_METADATA

	head_b, err := os.ReadFile(".npia/_apix_o/HEAD")

	if err != nil {
		panic(err.Error())
	}

	body_b, err := os.ReadFile(".npia/_apix_o/BODY")

	err = json.Unmarshal(head_b, &HEAD)

	if err != nil {
		panic(err.Error())
	}

	if !HEAD.JSON {

		body_str := string(body_b)

		fmt.Println(body_str)

	} else {
		fmt.Println("implementing...")
	}

}
