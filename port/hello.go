package port

import (
	_model "kp_api/models"
)

type HelloInferface interface {
	HelloWeb(request _model.HelloModelRequest) _model.HelloModelReponse
}
