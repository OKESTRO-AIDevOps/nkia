package admin

import (
	"fmt"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
)

func RequestHandler(api_input apistandard.API_INPUT) error {

	ASgi := apistandard.ASgi

	api_out, err := ASgi.Run(api_input)

	if err != nil {
		return fmt.Errorf("request handler: %s", err.Error())
	}

	body := api_out.BODY

	fmt.Println("----------MESSAGE----------")
	fmt.Println(body)

	return nil

}
