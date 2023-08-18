package controller

import (
	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
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

type OrchestratorRequest struct {
	RequestOption string `json:"request_option"`

	RequestTarget string `json:"request_target"`

	Query string `json:"query"`
}

type OrchestratorResponse struct {
	ServerMessage string `json:"server_message"`

	QueryResult []byte `json:"query_result"`
}
