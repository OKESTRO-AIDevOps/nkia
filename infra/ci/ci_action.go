package ci

import (
	"fmt"

	cigit "github.com/OKESTRO-AIDevOps/nkia/infra/git"
	"gopkg.in/yaml.v3"
)

func ActionV1(targets *CITargetsCtl, tidx int) {

	vid := targets.CI_TARGETS[tidx].CI_VERSION_ID

	tid := targets.CI_TARGETS[tidx].CI_TARGET_ID

	var target_ci TargetV1

	log_line := ""

	err := yaml.Unmarshal(targets.CI_TARGETS[tidx].CI_PLAYBOOK, &target_ci)

	if err != nil {

		log_line = fmt.Sprintf("failed to unmarshal playbook: %s", err.Error())

		LogErrAbortV1(targets, tidx, log_line)

		return
	}

	log_line = fmt.Sprintf("successfully unmarshalled playbook for: version: %d, target: %d", vid, tid)

	LogV1(targets, tidx, log_line)

	repo_addr := target_ci.GitPackage.Address

	repo_id := targets.CI_TARGETS[tidx].CI_USER_ID

	repo_pw := targets.CI_TARGETS[tidx].CI_USER_PW

	repo_nm := target_ci.GitPackage.Name

	err = cigit.GetDestRepo(repo_addr, repo_id, repo_pw, repo_nm)

	if err != nil {

		log_line = fmt.Sprintf("failed to get dest repo: %s", err.Error())

		LogErrAbortV1(targets, tidx, log_line)

		return

	}

	log_line = "successfully fetched destination repo"

	LogV1(targets, tidx, log_line)

	err = QueryV1(&target_ci)

	if err != nil {

		log_line = fmt.Sprintf("failed to execute query v1: %s", err.Error())

		LogErrAbortV1(targets, tidx, log_line)

		return

	}

	log_line = "successfully executed query defined in CI description"

	LogV1(targets, tidx, log_line)

	targets.CI_TARGETS[tidx].CI_STAT_FROM <- SIG_SUCCESS

	return

}

func LogV1(targets *CITargetsCtl, tidx int, log_line string) {

	CILOCK.Lock()

	targets.CI_TARGETS[tidx].CI_LOG = append(targets.CI_TARGETS[tidx].CI_LOG, log_line)

	CILOCK.Unlock()

}

func LogErrAbortV1(targets *CITargetsCtl, tidx int, log_line string) {

	CILOCK.Lock()

	targets.CI_TARGETS[tidx].CI_ERRLOG = append(targets.CI_TARGETS[tidx].CI_ERRLOG, log_line)

	CILOCK.Unlock()

	targets.CI_TARGETS[tidx].CI_STAT_FROM <- SIG_FAILED

}
