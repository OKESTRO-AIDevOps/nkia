package runtimefs

type AppOrigin struct {
	RECORDS []RecordInfo `json:"RECORDS,omitempty"`
	REPOS   []RepoInfo   `json:"REPOS,omitempty"`
	REGS    []RegInfo    `json:"REGS,omitempty"`
}

type RecordInfo struct {
	NS        string `json:"NS,omitempty"`
	REPO_ADDR string `json:"REPO_ADDR,omitempty"`
	REG_ADDR  string `json:"REG_ADDR,omitempty"`
}

type RepoInfo struct {
	REPO_ADDR string `json:"REPO_ADDR,omitempty"`
	REPO_ID   string `json:"REPO_ID,omitempty"`
	REPO_PW   string `json:"REPO_PW,omitempty"`
}

type RegInfo struct {
	REG_ADDR string `json:"REG_ADDR,omitempty"`
	REG_ID   string `json:"REG_ID,omitempty"`
	REG_PW   string `json:"REG_PW,omitempty"`
}
