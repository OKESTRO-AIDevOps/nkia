package controller

import (
	"github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"
)

type APIMessageRequest struct {
	Query string `json:"query"`
}

type APIMessageResponse struct {
	ServerMessage string `json:"server_message"`

	QueryResult apistandard.API_OUTPUT `json:"query_result"`
}

type ClientBaseRequest struct {
}

type ClientBaseResponse struct {
}
