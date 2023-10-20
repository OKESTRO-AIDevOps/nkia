package apix

import (
	"fmt"
	"reflect"
	"strings"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/nokubelet/controller"
)

func (axgi API_X) ConvertToOrchIOSpec(v interface{}) (ctrl.OrchestratorRequest, error) {

	// var tmp map[string]string

	oreq := ctrl.OrchestratorRequest{}

	v_type := reflect.TypeOf(v)

	if v_type == nil {
		return oreq, fmt.Errorf("failed to convert to orch io spec: %s", "nil type")
	}

	v_type_str := v_type.String()

	v_type_list := strings.Split(v_type_str, ".")

	idx := len(v_type_list) - 1

	key := v_type_list[idx]

	val := axgi[key]

	_ = val

	// org_struct :=

	return oreq, nil

}

func (axgi API_X) ConvertToOrchIOSpec_Test(v interface{}) {
	v_type := reflect.TypeOf(v)

	if v_type == nil {
		fmt.Printf("failed to convert to orch io spec: %s", "nil type")
		return
	}

	v_type_str := v_type.String()

	v_type_list := strings.Split(v_type_str, ".")

	idx := len(v_type_list) - 1

	key := v_type_list[idx]

	val := axgi[key]

	fmt.Printf("key: %s\n", key)

	fmt.Printf("val: %s\n", val)

}
