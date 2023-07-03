package controller

import (
	"github.com/OKESTRO-AIDevOps/npia-server/server/modules"
)

type APIMessageRequest struct {
	Query string `json:"query"`
}

type APIMessageResponse struct {
	ServerMessage string `json:"server_message"`

	QueryResult string `json:"query_result"`
}

// challenge_id : ASK, ANS
// response     : NOPE

type AuthChallenge struct {
	ChallengeID      string                 `json:"challenge_id"`
	ChallengeMessage string                 `json:"challenge_message"`
	ChallengeData    modules.ChallengRecord `json:"challenge_data"`
	ChallengeKey     modules.KeyRecord      `json:"challenge_key"`
}
