package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	cicmd "github.com/OKESTRO-AIDevOps/nkia/infra/cmd"
	"golang.org/x/term"
)

func GetSourceRepo(opts cicmd.NKIA_CI_OPTIONS) error {

	repo_addr := opts["repo"]

	repo_id := opts["id"]

	repo_pw := opts["token"]

	repo_nm := opts["name"]

	repo_nm = ".npia.infra/" + repo_nm

	if _, err := os.Stat(repo_nm); err == nil {

		cmd_rm := exec.Command("rm", "-rf", repo_nm)

		_ = cmd_rm.Run()

	}

	if strings.HasPrefix(repo_addr, "http://") {

		return fmt.Errorf("plain http:// unsupported, use https:// instead")

	}

	if strings.HasPrefix(repo_addr, "https://") {

		repo_addr = strings.Replace(repo_addr, "https://", "", 1)

	}

	if repo_pw == "-" {

		fmt.Printf("\ntype repo pw: ")

		pw_b, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return fmt.Errorf("get repo error: %s", err.Error())
		}

		repo_pw = string(pw_b)

	}

	insert := "%s:%s@"

	repo_addr = insert + repo_addr

	repo_addr = fmt.Sprintf(repo_addr, repo_id, repo_pw)

	repo_addr = "https://" + repo_addr

	cmd := exec.Command("git", "clone", repo_addr, repo_nm)

	err := cmd.Start()

	if err != nil {

		return fmt.Errorf("failed to start cloning: %s", err.Error())
	}

	fmt.Println("\nCloning source repo...")

	err = cmd.Wait()

	if err != nil {

		return fmt.Errorf("failed to clone: %s", err.Error())
	}

	fmt.Printf("successfuly cloned source repo at: %s\n", repo_nm)

	return nil
}

func GetDestRepo(repo_addr string, repo_id string, repo_pw string, repo_nm string) error {

	repo_nm = ".npia.infra/" + repo_nm

	if strings.HasPrefix(repo_addr, "http://") {

		return fmt.Errorf("plain http:// unsupported, use https:// instead")

	}

	if strings.HasPrefix(repo_addr, "https://") {

		repo_addr = strings.Replace(repo_addr, "https://", "", 1)

	}

	insert := "%s:%s@"

	repo_addr = insert + repo_addr

	repo_addr = fmt.Sprintf(repo_addr, repo_id, repo_pw)

	repo_addr = "https://" + repo_addr

	cmd := exec.Command("git", "clone", repo_addr, repo_nm)

	if err := cmd.Run(); err != nil {

		return fmt.Errorf("cmd run failed: %s", err.Error())

	}

	return nil
}

func CommitDestRepo() error {

	return nil

}

func PushDestRepo() error {

	return nil

}
