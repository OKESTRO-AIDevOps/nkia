package ci

import (
	"fmt"

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

		CILOCK.Lock()

		targets.CI_TARGETS[tidx].CI_ERRLOG = append(targets.CI_TARGETS[tidx].CI_ERRLOG, log_line)

		CILOCK.Unlock()

		targets.CI_TARGETS[tidx].CI_STAT_FROM <- -1

		return
	}

	log_line = fmt.Sprintf("successfully unmarshalled playbook for: version: %d, target: %d", vid, tid)

	CILOCK.Lock()

	targets.CI_TARGETS[tidx].CI_LOG = append(targets.CI_TARGETS[tidx].CI_LOG, log_line)

	CILOCK.Unlock()
	//CILOCK.Unlock()

	if tid > 1 {
		targets.CI_TARGETS[tidx].CI_STAT_FROM <- SIG_SUCCESS
	} else {

		targets.CI_TARGETS[tidx].CI_STAT_FROM <- SIG_FAILED

	}

	//time.Sleep(1000 * time.Second)
}
