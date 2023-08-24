package promquery

import (
	"fmt"
	"strings"
)

func PQ_PodScheduled(ns string) ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	if err := SanitizePQ(ns); err != nil {
		return ret, fmt.Errorf("error handling prom query: %s", err.Error())
	}

	ns = "'" + ns + "'"

	query = strings.ReplaceAll(Q_POD_SCHEDULED, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, POD)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_PodUnscheduled(ns string) ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	if err := SanitizePQ(ns); err != nil {
		return ret, fmt.Errorf("error handling prom query: %s", err.Error())
	}

	ns = "'" + ns + "'"

	query = strings.ReplaceAll(Q_POD_UNSCHEDULED, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, POD)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_ContainerCPUUsage(ns string) ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	if err := SanitizePQ(ns); err != nil {
		return ret, fmt.Errorf("error handling prom query: %s", err.Error())
	}

	ns = "'" + ns + "'"

	query = strings.ReplaceAll(Q_CONTAINER_CPU_USAGE, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, POD)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_ContainerMemUsage(ns string) ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	if err := SanitizePQ(ns); err != nil {
		return ret, fmt.Errorf("error handling prom query: %s", err.Error())
	}

	ns = "'" + ns + "'"

	query = strings.ReplaceAll(Q_CONTAINER_MEM_USAGE, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, POD)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_ContainerFSRead(ns string) ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	if err := SanitizePQ(ns); err != nil {
		return ret, fmt.Errorf("error handling prom query: %s", err.Error())
	}

	ns = "'" + ns + "'"

	query = strings.ReplaceAll(Q_CONTAINER_FS_READ, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, POD)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_ContainerFSWrite(ns string) ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	if err := SanitizePQ(ns); err != nil {
		return ret, fmt.Errorf("error handling prom query: %s", err.Error())
	}

	ns = "'" + ns + "'"

	query = strings.ReplaceAll(Q_CONTAINER_FS_WRITE, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, POD)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_ContainerNetworkReceive(ns string) ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	if err := SanitizePQ(ns); err != nil {
		return ret, fmt.Errorf("error handling prom query: %s", err.Error())
	}

	ns = "'" + ns + "'"

	query = strings.ReplaceAll(Q_CONTAINER_NETWORK_RECEIVE, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, POD)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_ContainerNetworkTransmit(ns string) ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	if err := SanitizePQ(ns); err != nil {
		return ret, fmt.Errorf("error handling prom query: %s", err.Error())
	}

	ns = "'" + ns + "'"

	query = strings.ReplaceAll(Q_CONTAINER_NETWORK_TRANSMIT, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, POD)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_KubeletVolumeAvailable() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	query = Q_KUBELET_VOLUME_AVAILABLE

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, PVC)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_KubeletVolumeCapacity() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	query = Q_KUBELET_VOLUME_CAPACITY

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, PVC)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_KubeletVolumeUsed() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	query = Q_KUBELET_VOLUME_USED

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, PVC)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeTemperatureCelsius() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_TEMPERATURE_CELSIUS, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, SENSOR)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeTemperatureCelsiusChange() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_TEMPERATURE_CELSIUS_CHANGE, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, SENSOR)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeTemperatureCelsiusAverage() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_TEMPERATURE_CELSIUS_AVERAGE, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, SENSOR)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeProcessRunning() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_PROCESS_RUNNING, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, INSTANCE)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeCPUCores() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	query = Q_NODE_CPU_CORES

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, NODE)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeMemActive() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_MEM_ACTIVE, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, INSTANCE)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeMemTotal() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_MEM_TOTAL, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, INSTANCE)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeDiskRead() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_DISK_READ, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, INSTANCE)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeDiskWrite() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_DISK_WRITE, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, INSTANCE)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeNetworkReceive() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_NETWORK_RECEIVE, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, INSTANCE)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeNetworkTransmit() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_NETWORK_TRANSMIT, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, INSTANCE)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}

func PQ_NodeDiskWrittenTotal() ([]PQOutputFormat, error) {

	var query string

	var ret []PQOutputFormat

	ns := "'" + "default" + "'"

	query = strings.ReplaceAll(Q_NODE_DISK_WRITTEN_TOTAL, "***", ns)

	body_bytes, err := PromQueryPost(query)

	if err != nil {

		return ret, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	ret_data, err := PromQueryStandardizer(body_bytes, INSTANCE)

	if err != nil {

		return ret_data, fmt.Errorf("error handling prom query: %s", err.Error())

	}

	return ret_data, nil

}
