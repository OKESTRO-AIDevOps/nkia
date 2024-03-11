package ciconfig

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadTargetsFromFile() (CITargets, error) {

	ci_targets := make(CITargets)

	var tmp_yaml map[string][]interface{}

	if _, err := os.Stat(".nkia/ci.yaml"); err != nil {

		return ci_targets, fmt.Errorf("nkia ci.yaml not found")

	}

	file_b, err := os.ReadFile(".nkia/ci.yaml")

	if err != nil {

		return ci_targets, fmt.Errorf("failed to read nkia ci.yaml")

	}

	err = yaml.Unmarshal(file_b, &tmp_yaml)

	if err != nil {

		return ci_targets, fmt.Errorf("failed to unmarshal ci.yaml")
	}

	for k, v := range tmp_yaml {

		tv_b, err := yaml.Marshal(v)

		if err != nil {

			return ci_targets, fmt.Errorf("error unmarshalling on key: %s", k)

		}

		ci_targets[k] = tv_b

	}

	return ci_targets, nil
}
