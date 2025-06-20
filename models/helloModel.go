package model

type HelloModelRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type HelloModelReponse struct {
	Message string `json:"message"`
}
