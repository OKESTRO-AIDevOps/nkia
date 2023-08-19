package builtinresource

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func HorizontalPodAutoscaler_test() error {

	var test_hpa HorizontalPodAutoscaler

	yaml_byte, err := yaml.Marshal(test_hpa)

	if err != nil {

		return fmt.Errorf("failed to marshal yaml: %s", err.Error())

	}

	err = os.WriteFile("tmp/hpa_test.yaml", yaml_byte, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())

	}

	return nil
}

func Ingress_test() error {

	var test_ingress Ingress

	yaml_byte, err := yaml.Marshal(test_ingress)

	if err != nil {

		return fmt.Errorf("failed to marshal yaml: %s", err.Error())

	}

	err = os.WriteFile("tmp/ingress_test.yaml", yaml_byte, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())

	}

	return nil
}

func Service_test() error {

	var test_service Service

	yaml_byte, err := yaml.Marshal(test_service)

	if err != nil {

		return fmt.Errorf("failed to marshal yaml: %s", err.Error())

	}

	err = os.WriteFile("tmp/service_test.yaml", yaml_byte, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())

	}

	return nil

}

func Deployment_test() error {

	var test_deployment Deployment

	yaml_byte, err := yaml.Marshal(test_deployment)

	if err != nil {

		return fmt.Errorf("failed to marshal yaml: %s", err.Error())

	}

	err = os.WriteFile("tmp/deployment_test.yaml", yaml_byte, 0644)

	if err != nil {

		return fmt.Errorf("failed to write file: %s", err.Error())

	}

	return nil

}
