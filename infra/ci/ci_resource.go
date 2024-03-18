package ci

import "sync"

type CITargets map[string][]interface{}

type CITargetsCtl struct {
	CI_ONGOING []CITargetID
	CI_SUCCESS []CITargetID
	CI_FAILED  []CITargetID
	CI_TARGETS []CITargetCtl
}

type CICredStore struct {
	CI_CRED []CITargetCred
}

type CITargetCred struct {
	CI_VERSION_ID int
	CI_TARGET_ID  int
	USER_ID       string
	USER_PW       string
	USER_EMAIL    string
}

type CITargetID struct {
	CI_VERSION_ID int
	CI_TARGET_ID  int
}

type CITargetCtl struct {
	CI_VERSION_ID int
	CI_TARGET_ID  int
	CI_USER_ID    string
	CI_USER_PW    string
	CI_USER_EMAIL string
	CI_CMD_TO     chan int
	CI_STAT_FROM  chan int
	CI_DONE       int
	CI_PLAYBOOK   []byte
	CI_LOG_PTR    int
	CI_LOG        []string
	CI_ERRLOG_PTR int
	CI_ERRLOG     []string
}

const (
	SIG_SUCCESS = 2
	SIG_FORWARD = 1
	SIG_FAILED  = -1
)

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

var CILOCK sync.Mutex
