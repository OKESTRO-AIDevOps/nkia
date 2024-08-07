package builtinresource

type HelmPromStackDefaultValue struct {
	Alertmanager struct {
		AlertmanagerSpec struct {
			Storage struct {
				VolumeClaimTemplate struct {
					Spec struct {
						StorageClassName string   `yaml:"storageClassName"`
						AccessModes      []string `yaml:"accessModes"`
						Resources        struct {
							Requests struct {
								Storage string `yaml:"storage"`
							} `yaml:"requests"`
						} `yaml:"resources"`
					} `yaml:"spec"`
				} `yaml:"volumeClaimTemplate"`
			} `yaml:"storage"`
		} `yaml:"alertmanagerSpec"`
	} `yaml:"alertmanager"`
	Grafana struct {
		Persistence struct {
			Enabled          bool     `yaml:"enabled"`
			Type             string   `yaml:"type"`
			StorageClassName string   `yaml:"storageClassName"`
			AccessModes      []string `yaml:"accessModes"`
			Size             string   `yaml:"size"`
		} `yaml:"persistence"`
	} `yaml:"grafana"`
	Prometheus struct {
		PrometheusSpec struct {
			StorageSpec struct {
				VolumeClaimTemplate struct {
					Spec struct {
						StorageClassName string   `yaml:"storageClassName"`
						AccessModes      []string `yaml:"accessModes"`
						Resources        struct {
							Requests struct {
								Storage string `yaml:"storage"`
							} `yaml:"requests"`
						} `yaml:"resources"`
					} `yaml:"spec"`
				} `yaml:"volumeClaimTemplate"`
			} `yaml:"storageSpec"`
		} `yaml:"prometheusSpec"`
	} `yaml:"prometheus"`
}

type KanikoBuilder struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Containers    []KanikoBuilder_Container `yaml:"containers"`
		RestartPolicy string                    `yaml:"restartPolicy"`
		Volumes       []KanikoBuilder_Volume    `yaml:"volumes"`
	} `yaml:"spec"`
}

type KanikoBuilder_Container struct {
	Name         string                                `yaml:"name"`
	Image        string                                `yaml:"image"`
	Args         []string                              `yaml:"args"`
	VolumeMounts []KanikoBuilder_Container_VolumeMount `yaml:"volumeMounts"`
	Env          []KanikoBuilder_Container_Env         `yaml:"env"`
}

type KanikoBuilder_Container_VolumeMount struct {
	Name      string `yaml:"name"`
	MountPath string `yaml:"mountPath"`
}

type KanikoBuilder_Container_Env struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type KanikoBuilder_Volume struct {
	Name   string `yaml:"name"`
	Secret struct {
		SecretName string                      `yaml:"secretName"`
		Items      []KanikoBuilder_Volume_Item `yaml:"items"`
	} `yaml:"secret"`
}

type KanikoBuilder_Volume_Item struct {
	Key  string `yaml:"key"`
	Path string `yaml:"path"`
}
