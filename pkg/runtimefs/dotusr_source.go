package runtimefs

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	bsrc "github.com/OKESTRO-AIDevOps/nkia/pkg/builtinresource"
	goya "github.com/goccy/go-yaml"
)

func CreateHPASource(resourcenm string, resource_key string) (string, error) {

	src_found := 0

	file_byte, err := os.ReadFile(".usr/ops_src.yaml")

	if err != nil {
		return "", fmt.Errorf("failed to create hpa src: %s", err.Error())
	}

	cmd := exec.Command("kubectl", "get", "nodes", "-o", "yaml")

	out, err := cmd.Output()

	var yaml_items []interface{}

	yaml_str := string(out)

	yaml_path_items := "$.items"

	ypath, _ := goya.PathString(yaml_path_items)

	err = ypath.Read(strings.NewReader(yaml_str), &yaml_items)

	if err != nil {
		return "", fmt.Errorf("failed to create hpa src: %s", err.Error())
	}

	min_repl := 0

	max_repl := 0

	prev_top := 0

	pods := 0

	for _, val := range yaml_items {

		yaml_if := make(map[interface{}]interface{})

		resource_b, err := goya.Marshal(val)

		err = goya.Unmarshal(resource_b, &yaml_if)

		if err != nil {
			return "", fmt.Errorf("failed to create hpa src: %s", err.Error())
		}

		pods_str := yaml_if["status"].(map[string]interface{})["allocatable"].(map[string]interface{})["pods"].(string)

		pods, err = strconv.Atoi(pods_str)

		if err != nil {
			return "", fmt.Errorf("failed to create hpa src: %s", err.Error())
		}

		if pods > prev_top {
			prev_top = pods
		}

	}

	pods = prev_top

	min_repl = int(float64(pods) * 0.02)

	max_repl = int(float64(pods) * 0.1)

	head_metadataName := "hpa-deployment-" + resourcenm
	apiVersion := ""
	kind := resource_key
	metadata_name := resourcenm

	var out_hpa bsrc.HorizontalPodAutoscaler

	file_str := string(file_byte)

	file_str_list := strings.Split(file_str, "---\n")

	for _, content := range file_str_list {

		if content == "\n" || content == "" {
			continue
		}

		yaml_if := make(map[interface{}]interface{})

		c_byte := []byte(content)

		if err != nil {
			return "", fmt.Errorf("failed to create hpa src: %s", err.Error())
		}

		err = goya.Unmarshal(c_byte, &yaml_if)

		if err != nil {
			return "", fmt.Errorf("failed to create hpa src: %s", err.Error())
		}

		if yaml_if["kind"] == resource_key && yaml_if["metadata"].(map[string]interface{})["name"] == resourcenm {

			src_found = 1
			apiVersion = yaml_if["apiVersion"].(string)
			kind = yaml_if["kind"].(string)
			metadata_name = yaml_if["metadata"].(map[string]interface{})["name"].(string)

			break
		}

	}

	if src_found == 0 {
		return "", fmt.Errorf("failed to create hpa src: %s", "matching key not found")
	}

	if min_repl == 0 {
		min_repl = 1
	}

	if min_repl > max_repl {
		max_repl = min_repl + 1
	}

	out_hpa.APIVersion = "autoscaling/v1"
	out_hpa.Kind = "HorizontalPodAutoscaler"
	out_hpa.Metadata.Name = head_metadataName
	out_hpa.Spec.ScaleTargetRef.APIVersion = apiVersion
	out_hpa.Spec.ScaleTargetRef.Kind = kind
	out_hpa.Spec.ScaleTargetRef.Name = metadata_name
	out_hpa.Spec.MinReplicas = min_repl
	out_hpa.Spec.MaxReplicas = max_repl
	out_hpa.Spec.TargetCPUUtilizationPercentage = 50

	out_hpa_byte, err := goya.Marshal(out_hpa)

	if err != nil {
		return "", fmt.Errorf("failed to create hpa src: %s", err.Error())
	}

	err = os.WriteFile(".usr/hpa_src.yaml", out_hpa_byte, 0644)

	if err != nil {
		return "", fmt.Errorf("failed to create hpa src: %s", err.Error())
	}

	return ".usr/hpa_src.yaml", nil

}

func CreateQOSSource(resourcenm string, resource_key string) (string, error) {

	src_found := 0

	file_byte, err := os.ReadFile(".usr/ops_src.yaml")

	if err != nil {
		return "", fmt.Errorf("failed to create qos src: %s", err.Error())
	}

	cmd := exec.Command("kubectl", "get", "nodes", "-o", "yaml")

	out, err := cmd.Output()

	var yaml_items []interface{}

	yaml_str := string(out)

	yaml_path_items := "$.items"

	ypath, _ := goya.PathString(yaml_path_items)

	err = ypath.Read(strings.NewReader(yaml_str), &yaml_items)

	if err != nil {
		return "", fmt.Errorf("failed to create qos src: %s", err.Error())
	}

	polled_node_index := 0

	prev_top := 0

	pods := 0

	for i, val := range yaml_items {

		yaml_if := make(map[interface{}]interface{})

		resource_b, err := goya.Marshal(val)

		err = goya.Unmarshal(resource_b, &yaml_if)

		if err != nil {
			return "", fmt.Errorf("failed to create qos src: %s", err.Error())
		}

		pods_str := yaml_if["status"].(map[string]interface{})["allocatable"].(map[string]interface{})["pods"].(string)

		pods, err = strconv.Atoi(pods_str)

		if err != nil {
			return "", fmt.Errorf("failed to create qos src: %s", err.Error())
		}

		if pods > prev_top {
			prev_top = pods
			polled_node_index = i

		}

	}

	pods = prev_top

	polled_cpu := yaml_items[polled_node_index].(map[string]interface{})["status"].(map[string]interface{})["allocatable"].(map[string]interface{})["cpu"].(string)

	polled_mem := yaml_items[polled_node_index].(map[string]interface{})["status"].(map[string]interface{})["allocatable"].(map[string]interface{})["memory"].(string)

	pods_float := float64(pods)

	polled_cpu_float, err := strconv.ParseFloat(polled_cpu, 64)

	if err != nil {
		return "", fmt.Errorf("failed to create qos src: %s", err.Error())
	}

	polled_mem_str := strings.ReplaceAll(polled_mem, "Ki", "")

	polled_mem_float, err := strconv.ParseFloat(polled_mem_str, 64)

	if err != nil {
		return "", fmt.Errorf("failed to create qos src: %s", err.Error())
	}

	cpu_limit_per_node := (polled_cpu_float / pods_float) * 8.0

	mem_limit_per_node := (polled_mem_float / pods_float) * 16.0

	qos_cpu_high := strconv.FormatFloat(cpu_limit_per_node*0.8, 'f', -1, 64) + "m"

	qos_mem_high := strconv.FormatFloat(mem_limit_per_node*0.8, 'f', -1, 64) + "Ki"

	qos_cpu_middle := strconv.FormatFloat(cpu_limit_per_node*0.5, 'f', -1, 64) + "m"

	qos_mem_middle := strconv.FormatFloat(mem_limit_per_node*0.5, 'f', -1, 64) + "Ki"

	cpu_limits := qos_cpu_high

	mem_limits := qos_mem_high

	cpu_requests := qos_cpu_middle

	mem_requests := qos_mem_middle

	var out_qos []byte

	file_str := string(file_byte)

	file_str_list := strings.Split(file_str, "---\n")

	for _, content := range file_str_list {

		if content == "\n" || content == "" {
			continue
		}

		yaml_if := make(map[interface{}]interface{})

		c_byte := []byte(content)

		if err != nil {
			return "", fmt.Errorf("failed to create qos src: %s", err.Error())
		}

		err = goya.Unmarshal(c_byte, &yaml_if)

		if err != nil {
			return "", fmt.Errorf("failed to create qos src: %s", err.Error())
		}

		if yaml_if["kind"] == resource_key && yaml_if["metadata"].(map[string]interface{})["name"] == resourcenm {

			src_found = 1

			c_count := len(yaml_if["spec"].(map[string]interface{})["template"].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{}))

			for j := 0; j < c_count; j++ {

				rsc := map[string]map[string]string{
					"limits": {
						"cpu":    cpu_limits,
						"memory": mem_limits,
					},
					"requests": {
						"cpu":    cpu_requests,
						"memory": mem_requests,
					},
				}

				yaml_if["spec"].(map[string]interface{})["template"].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{})[j].(map[string]interface{})["resources"] = rsc

			}

			b_yaml_if, err := goya.Marshal(yaml_if)

			if err != nil {
				return "", fmt.Errorf("failed to create qos src: %s", err.Error())
			}

			out_qos = b_yaml_if

			break
		}

	}

	if src_found == 0 {
		return "", fmt.Errorf("failed to create qos src: %s", "matching key not found")
	}

	err = os.WriteFile(".usr/qos_src.yaml", out_qos, 0644)

	if err != nil {
		return "", fmt.Errorf("failed to create qos src: %s", err.Error())
	}

	return ".usr/qos_src.yaml", nil

}

func CreateDelQOSSource(resourcenm string, resource_key string) (string, error) {

	src_found := 0

	file_byte, err := os.ReadFile(".usr/ops_src.yaml")

	if err != nil {
		return "", fmt.Errorf("failed to create del qos src: %s", err.Error())
	}

	var out_del_qos []byte

	file_str := string(file_byte)

	file_str_list := strings.Split(file_str, "---\n")

	for _, content := range file_str_list {

		if content == "\n" || content == "" {
			continue
		}

		yaml_if := make(map[interface{}]interface{})

		c_byte := []byte(content)

		if err != nil {
			return "", fmt.Errorf("failed to create del qos src: %s", err.Error())
		}

		err = goya.Unmarshal(c_byte, &yaml_if)

		if err != nil {
			return "", fmt.Errorf("failed to create del qos src: %s", err.Error())
		}

		if yaml_if["kind"] == resource_key && yaml_if["metadata"].(map[string]interface{})["name"] == resourcenm {

			src_found = 1

			b_yaml_if, err := goya.Marshal(yaml_if)

			if err != nil {
				return "", fmt.Errorf("failed to create del qos src: %s", err.Error())
			}

			out_del_qos = b_yaml_if

			break
		}

	}

	if src_found == 0 {
		return "", fmt.Errorf("failed to create qos src: %s", "matching key not found")
	}

	err = os.WriteFile(".usr/del_qos_src.yaml", out_del_qos, 0644)

	if err != nil {
		return "", fmt.Errorf("failed to create del qos src: %s", err.Error())
	}

	return ".usr/del_qos_src.yaml", nil

}

func CreateIngressSource(ns string, host string, svc string) (string, error) {

	src_found := 0

	file_byte, err := os.ReadFile(".usr/ops_src.yaml")

	if err != nil {
		return "", fmt.Errorf("failed to create ingr src: %s", err.Error())
	}

	var port_number uint64

	var out_ingr bsrc.Ingress

	var ingr_rules bsrc.Ingress_Rules

	var ingr_rules_paths bsrc.Ingress_Rules_Paths

	file_str := string(file_byte)

	file_str_list := strings.Split(file_str, "---\n")

	for _, content := range file_str_list {

		if content == "\n" || content == "" {
			continue
		}

		yaml_if := make(map[interface{}]interface{})

		c_byte := []byte(content)

		if err != nil {
			return "", fmt.Errorf("failed to create ingr src: %s", err.Error())
		}

		err = goya.Unmarshal(c_byte, &yaml_if)

		if err != nil {
			return "", fmt.Errorf("failed to create ingr src: %s", err.Error())
		}

		if yaml_if["kind"] == "Service" && yaml_if["metadata"].(map[string]interface{})["name"] == svc {

			src_found = 1

			port_number = yaml_if["spec"].(map[string]interface{})["ports"].([]interface{})[0].(map[string]interface{})["port"].(uint64)

			break
		}

	}

	if src_found == 0 {
		return "", fmt.Errorf("failed to create ingr src: %s", "matching key not found")
	}

	ingr_rules_paths.Backend.Service.Name = svc
	ingr_rules_paths.Backend.Service.Port.Number = int(port_number)
	ingr_rules_paths.Path = "/"
	ingr_rules_paths.PathType = "Prefix"

	ingr_rules.Host = host
	ingr_rules.HTTP.Paths = append(ingr_rules.HTTP.Paths, ingr_rules_paths)

	out_ingr.APIVersion = "networking.k8s.io/v1"
	out_ingr.Kind = "Ingress"
	out_ingr.Metadata.Name = "ingress-" + ns
	out_ingr.Metadata.Annotations.NginxIngressKubernetesIoProxyBodySize = "0"

	out_ingr.Spec.Rules = append(out_ingr.Spec.Rules, ingr_rules)

	out_ingr_byte, err := goya.Marshal(out_ingr)

	if err != nil {
		return "", fmt.Errorf("failed to create ingr src: %s", err.Error())
	}

	err = os.WriteFile(".usr/ingr_src.yaml", out_ingr_byte, 0644)

	if err != nil {
		return "", fmt.Errorf("failed to create ingr src: %s", err.Error())
	}

	return ".usr/ingr_src.yaml", nil
}

func CreateNodePortSource(ns string, svc string) (string, error) {

	src_found := 0

	file_byte, err := os.ReadFile(".usr/ops_src.yaml")

	if err != nil {
		return "", fmt.Errorf("failed to create ndpt src: %s", err.Error())
	}

	var node_port_number int

	var port_number uint64

	var out_ndpt bsrc.NodePort

	var ndpt_ports bsrc.NodePort_Ports

	file_str := string(file_byte)

	file_str_list := strings.Split(file_str, "---\n")

	for _, content := range file_str_list {

		if content == "\n" || content == "" {
			continue
		}

		yaml_if := make(map[interface{}]interface{})

		c_byte := []byte(content)

		if err != nil {
			return "", fmt.Errorf("failed to create ndpt src: %s", err.Error())
		}

		err = goya.Unmarshal(c_byte, &yaml_if)

		if err != nil {
			return "", fmt.Errorf("failed to create ndpt src: %s", err.Error())
		}

		if yaml_if["kind"] == "Service" && yaml_if["metadata"].(map[string]interface{})["name"] == svc {

			src_found = 1

			port_number = yaml_if["spec"].(map[string]interface{})["ports"].([]interface{})[0].(map[string]interface{})["port"].(uint64)

			break
		}

	}

	if src_found == 0 {
		return "", fmt.Errorf("failed to create ndpt src: %s", "matching key not found")
	}

	rand.Seed(time.Now().UnixNano())
	min := 30000
	max := 32767
	node_port_number = rand.Intn(max-min+1) + min

	ndpt_ports.NodePort = node_port_number
	ndpt_ports.Port = int(port_number)
	ndpt_ports.TargetPort = int(port_number)

	out_ndpt.Spec.Selector.IoKomposeService = svc
	out_ndpt.Spec.Type = "NodePort"
	out_ndpt.Metadata.Name = "nodeport-" + ns
	out_ndpt.APIVersion = "v1"
	out_ndpt.Kind = "Service"

	out_ndpt_byte, err := goya.Marshal(out_ndpt)

	if err != nil {
		return "", fmt.Errorf("failed to create ndpt src: %s", err.Error())
	}

	err = os.WriteFile(".usr/ndpt_src.yaml", out_ndpt_byte, 0644)

	if err != nil {
		return "", fmt.Errorf("failed to create ndpt src: %s", err.Error())
	}

	return ".usr/ndpt_src.yaml", nil
}
