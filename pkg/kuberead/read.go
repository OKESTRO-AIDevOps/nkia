package kuberead

import (
	"encoding/json"
	"fmt"

	"os/exec"

	metq "github.com/OKESTRO-AIDevOps/nkia/pkg/metricquery"
)

func ReadPod(main_ns string) ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "pods")

	out, err := cmd.Output()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func ReadPodLog(main_ns string, pod_name string) ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "logs", "-n", main_ns, pod_name)

	out, err := cmd.Output()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func ReadService(main_ns string) ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "services")

	out, err := cmd.Output()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func ReadDeployment(main_ns string) ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "deployments")

	out, err := cmd.Output()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func ReadNode(main_ns string) ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "get", "nodes")

	out, err := cmd.Output()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func ReadEvent(main_ns string) ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "events")

	out, err := cmd.Output()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func ReadResource(main_ns string) ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "all")

	out, err := cmd.Output()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func ReadNamespace(main_ns string) ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "get", "namespaces")

	out, err := cmd.Output()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = out

	return ret_byte, nil

}

func ReadIngress(main_ns string) ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "ingress")

	stdout, err := cmd.Output()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = stdout

	return ret_byte, nil

}

func ReadNodePort(main_ns string) ([]byte, error) {

	var ret_byte []byte

	cmd := exec.Command("kubectl", "-n", main_ns, "get", "services")

	stdout, err := cmd.Output()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())
	}

	ret_byte = stdout

	return ret_byte, nil
}

func ReadImageList(main_ns string) ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil
}

func ReadProjectProbe(main_ns string) ([]byte, error) {

	var ret_byte []byte

	return ret_byte, nil
}

func ReadPodScheduled(main_ns string) ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_PodScheduled(main_ns)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadPodUnscheduled(main_ns string) ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_PodUnscheduled(main_ns)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadContainerCPUUsage(main_ns string) ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_ContainerCPUUsage(main_ns)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadContainerMemUsage(main_ns string) ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_ContainerMemUsage(main_ns)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadContainerFSRead(main_ns string) ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_ContainerFSRead(main_ns)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadContainerFSWrite(main_ns string) ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_ContainerFSWrite(main_ns)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadContainerNetworkReceive(main_ns string) ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_ContainerNetworkReceive(main_ns)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadContainerNetworkTransmit(main_ns string) ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_ContainerNetworkTransmit(main_ns)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadKubeletVolumeAvailable() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_KubeletVolumeAvailable()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadKubeletVolumeCapacity() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_KubeletVolumeCapacity()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadKubeletVolumeUsed() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_KubeletVolumeUsed()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeTemperatureCelsius() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeTemperatureCelsius()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeTemperatureCelsiusChange() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeTemperatureCelsiusChange()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeTemperatureCelsiusAverage() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeTemperatureCelsiusAverage()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeProcessRunning() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeProcessRunning()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeCPUCores() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeCPUCores()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeMemActive() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeMemActive()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeMemTotal() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeMemTotal()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeDiskRead() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeDiskRead()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeDiskWrite() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeDiskWrite()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeNetworkReceive() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeNetworkReceive()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeNetworkTransmit() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeNetworkTransmit()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}

func ReadNodeDiskWrittenTotal() ([]byte, error) {

	var ret_byte []byte

	metq_out, err := metq.PQ_NodeDiskWrittenTotal()

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	ret_byte, err = json.Marshal(metq_out)

	if err != nil {

		return ret_byte, fmt.Errorf(": %s", err.Error())

	}

	return ret_byte, nil

}
