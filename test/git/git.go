package git

import (
	"fmt"
	"os/exec"
	"strings"

	cicmd "github.com/OKESTRO-AIDevOps/nkia/test/cmd"
)

func GetRepo(opts cicmd.NKIA_CI_OPTIONS) error {

	repo_addr := opts["repo"]

	repo_id := opts["id"]

	repo_pw := opts["pw"]

	repo_nm := opts["name"]

	if strings.HasPrefix(repo_addr, "http://") {

		return fmt.Errorf("plain http:// unsupported, use https:// instead")

	}

	if strings.HasPrefix(repo_addr, "https://") {

		repo_addr = strings.Replace(repo_addr, "https://", "", 1)

	}

	insert := "%s:%s@"

	repo_addr = insert + repo_addr

	repo_addr = fmt.Sprintf(repo_addr, repo_id, repo_pw)

	repo_addr += "https://" + repo_addr

	cmd := exec.Command("git", "clone", repo_addr, repo_nm)

	if err := cmd.Run(); err != nil {

		return fmt.Errorf("cmd run failed: %s", err.Error())

	}

	return nil
}
