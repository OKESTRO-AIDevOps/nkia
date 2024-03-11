package ciconfig

type CITargets map[string][]byte

type CIV1 struct {
	TargetV1 []TargetV1 `yaml:"target.v1"`
}

type TargetV1 struct {
	GitPackage struct {
		Address  string   `yaml:"address"`
		Name     string   `yaml:"name"`
		Strategy string   `yaml:"strategy"`
		Lock     []string `yaml:"lock"`
	} `yaml:"gitPackage"`
	Root   string           `yaml:"root"`
	Select []TargetV1Select `yaml:"select"`
}

type TargetV1Select struct {
	What []string `yaml:"what"`
	From string   `yaml:"from"`
	Not  []string `yaml:"not"`
	As   string   `yaml:"as"`
}
