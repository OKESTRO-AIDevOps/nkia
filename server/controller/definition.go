package controller

import (
	"github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"
)

type JSONMessageRequest struct {
	Query string `json:"query"`
}

type JSONMessageResponse struct {
	ServerMessage string `json:"server_message"`

	QueryResult apistandard.API_OUTPUT `json:"query_result"`
}
