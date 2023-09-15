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
