package runtimefs

import (
	"fmt"
	"os"

	btsrc "github.com/OKESTRO-AIDevOps/nkia/pkg/builtinresource"

	goya "github.com/goccy/go-yaml"
)

func CreateDefaultStorageSource() (string, error) {

	var ret_path string

	var dsc btsrc.StorageClass

	dsc.Kind = "StorageClass"

	dsc.APIVersion = "storage.k8s.io/v1"

	dsc.Metadata.Annotations.StorageclassKubernetesIoIsDefaultClass = "true"

	dsc.Metadata.Name = "nfs-default-storageclass"

	dsc.Provisioner = "cluster.local/nfs-subdir-external-provisioner"

	dsc.ReclaimPolicy = "Delete"

	dsc.AllowVolumeExpansion = true

	dsc.VolumeBindingMode = "Immediate"

	result_b, err := goya.Marshal(dsc)

	if err != nil {

		return ret_path, fmt.Errorf("failed to create default storage class source: %s", err.Error())

	}

	err = os.WriteFile(".npia/default-storage-class.yaml", result_b, 0644)

	if err != nil {
		return "", fmt.Errorf("failed to create default sc: %s", err.Error())
	}

	ret_path = ".npia/default-storage-class.yaml"

	return ret_path, nil

}

func CreateDefaultHelmValueSource() (string, error) {

	var ret_path string

	var dhvs btsrc.HelmPromStackDefaultValue

	dhvs.Alertmanager.AlertmanagerSpec.Storage.VolumeClaimTemplate.Spec.StorageClassName = "nfs-default-storageclass"

	dhvs.Alertmanager.AlertmanagerSpec.Storage.VolumeClaimTemplate.Spec.AccessModes = append(dhvs.Alertmanager.AlertmanagerSpec.Storage.VolumeClaimTemplate.Spec.AccessModes, "ReadWriteOnce")

	dhvs.Alertmanager.AlertmanagerSpec.Storage.VolumeClaimTemplate.Spec.Resources.Requests.Storage = "4Gi"

	dhvs.Grafana.Persistence.Enabled = true

	dhvs.Grafana.Persistence.Type = "pvc"

	dhvs.Grafana.Persistence.StorageClassName = "nfs-default-storageclass"

	dhvs.Grafana.Persistence.AccessModes = append(dhvs.Grafana.Persistence.AccessModes, "ReadWriteOnce")

	dhvs.Grafana.Persistence.Size = "4Gi"

	dhvs.Prometheus.PrometheusSpec.StorageSpec.VolumeClaimTemplate.Spec.StorageClassName = "nfs-default-storageclass"

	dhvs.Prometheus.PrometheusSpec.StorageSpec.VolumeClaimTemplate.Spec.AccessModes = append(dhvs.Prometheus.PrometheusSpec.StorageSpec.VolumeClaimTemplate.Spec.AccessModes, "ReadWriteOnce")

	dhvs.Prometheus.PrometheusSpec.StorageSpec.VolumeClaimTemplate.Spec.Resources.Requests.Storage = "4Gi"

	result_b, err := goya.Marshal(dhvs)

	if err != nil {

		return ret_path, fmt.Errorf("failed to create default helm kubeprom value: %s", err.Error())

	}

	err = os.WriteFile(".npia/default-helm-kubeprom-value.yaml", result_b, 0644)

	if err != nil {
		return "", fmt.Errorf("failed to create default helm kpv: %s", err.Error())
	}

	ret_path = ".npia/default-helm-kubeprom-value.yaml"

	return ret_path, nil

}
