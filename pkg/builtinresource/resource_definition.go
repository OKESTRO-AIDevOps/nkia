package builtinresource

type HorizontalPodAutoscaler struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ScaleTargetRef struct {
			APIVersion string `yaml:"apiVersion"`
			Kind       string `yaml:"kind"`
			Name       string `yaml:"name"`
		} `yaml:"scaleTargetRef"`
		MinReplicas                    int `yaml:"minReplicas"`
		MaxReplicas                    int `yaml:"maxReplicas"`
		TargetCPUUtilizationPercentage int `yaml:"targetCPUUtilizationPercentage"`
	} `yaml:"spec"`
}

type Ingress struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name        string `yaml:"name"`
		Annotations struct {
			NginxIngressKubernetesIoProxyBodySize string `yaml:"nginx.ingress.kubernetes.io/proxy-body-size"`
		} `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec struct {
		Rules []Ingress_Rules `yaml:"rules"`
	} `yaml:"spec"`
}

type Ingress_Rules struct {
	Host string `yaml:"host"`
	HTTP struct {
		Paths []Ingress_Rules_Paths `yaml:"paths"`
	} `yaml:"http"`
}

type Ingress_Rules_Paths struct {
	Path     string `yaml:"path"`
	PathType string `yaml:"pathType"`
	Backend  struct {
		Service struct {
			Name string `yaml:"name"`
			Port struct {
				Number int `yaml:"number"`
			} `yaml:"port"`
		} `yaml:"service"`
	} `yaml:"backend"`
}

type NodePort struct {
	Kind       string `yaml:"kind"`
	APIVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Type     string `yaml:"type"`
		Selector struct {
			IoKomposeService string `yaml:"io.kompose.service"`
		} `yaml:"selector"`
		Ports []NodePort_Ports `yaml:"ports"`
	} `yaml:"spec"`
}

type NodePort_Ports struct {
	NodePort   int `yaml:"nodePort"`
	Port       int `yaml:"port"`
	TargetPort int `yaml:"targetPort"`
}

type Service struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			App string `yaml:"app"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	Spec struct {
		Type     string          `yaml:"type"`
		Ports    []Service_Ports `yaml:"ports"`
		Selector struct {
			App string `yaml:"app"`
		} `yaml:"selector"`
	} `yaml:"spec"`
}

type Service_Ports struct {
	Port       int    `yaml:"port"`
	TargetPort int    `yaml:"targetPort"`
	Protocol   string `yaml:"protocol"`
}

type Deployment struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Selector struct {
			MatchLabels struct {
				App string `yaml:"app"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Replicas int `yaml:"replicas"`
		Template struct {
			Metadata struct {
				Labels struct {
					App string `yaml:"app"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec struct {
				ImagePullSecrets []Deployment_ImagePullSecrets `yaml:"imagePullSecrets"`
				Containers       []Deployment_Containers       `yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}

type Deployment_ImagePullSecrets struct {
	Name string `yaml:"name"`
}

type Deployment_Containers struct {
	Name            string                        `yaml:"name"`
	Image           string                        `yaml:"image"`
	ImagePullPolicy string                        `yaml:"imagePullPolicy"`
	Ports           []Deployment_Containers_Ports `yaml:"ports"`
}

type Deployment_Containers_Ports struct {
	ContainerPort int `yaml:"containerPort"`
}

type PersistentVolume struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			Type string `yaml:"type"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	Spec struct {
		Capacity struct {
			Storage string `yaml:"storage"`
		} `yaml:"capacity"`
		VolumeMode                    string   `yaml:"volumeMode"`
		AccessModes                   []string `yaml:"accessModes"`
		PersistentVolumeReclaimPolicy string   `yaml:"persistentVolumeReclaimPolicy"`
		StorageClassName              string   `yaml:"storageClassName"`
		HostPath                      struct {
			Path string `yaml:"path"`
		} `yaml:"hostPath"`
	} `yaml:"spec"`
}

type PersistentVolumeClaim struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		StorageClassName string   `yaml:"storageClassName"`
		AccessModes      []string `yaml:"accessModes"`
		Resources        struct {
			Requests struct {
				Storage string `yaml:"storage"`
			} `yaml:"requests"`
		} `yaml:"resources"`
	} `yaml:"spec"`
}

type DeploymentVol struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Selector struct {
			MatchLabels struct {
				App string `yaml:"app"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Replicas int `yaml:"replicas"`
		Template struct {
			Metadata struct {
				Labels struct {
					App string `yaml:"app"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec struct {
				ImagePullSecrets []DeploymentVol_ImagePullSecrets `yaml:"imagePullSecrets"`
				Volumes          []DeploymentVol_Volumes          `yaml:"volumes"`
				Containers       []DeploymentVol_Containers       `yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}

type DeploymentVol_ImagePullSecrets struct {
	Name string `yaml:"name"`
}

type DeploymentVol_Volumes struct {
	Name                  string `yaml:"name"`
	PersistentVolumeClaim struct {
		ClaimName string `yaml:"claimName"`
	} `yaml:"persistentVolumeClaim"`
}

type DeploymentVol_Containers struct {
	Name            string                                  `yaml:"name"`
	Image           string                                  `yaml:"image"`
	ImagePullPolicy string                                  `yaml:"imagePullPolicy"`
	Ports           []DeploymentVol_Containers_Ports        `yaml:"ports"`
	VolumeMounts    []DeploymentVol_Containers_VolumeMounts `yaml:"volumeMounts"`
}

type DeploymentVol_Containers_Ports struct {
	ContainerPort int `yaml:"containerPort"`
}

type DeploymentVol_Containers_VolumeMounts struct {
	MountPath string `yaml:"mountPath"`
	Name      string `yaml:"name"`
}

type StorageClass struct {
	Kind       string `yaml:"kind"`
	APIVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Annotations struct {
			StorageclassKubernetesIoIsDefaultClass string `yaml:"storageclass.kubernetes.io/is-default-class"`
		} `yaml:"annotations"`
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Provisioner          string `yaml:"provisioner"`
	ReclaimPolicy        string `yaml:"reclaimPolicy"`
	AllowVolumeExpansion bool   `yaml:"allowVolumeExpansion"`
	VolumeBindingMode    string `yaml:"volumeBindingMode"`
}
