package validations

import (
	"playing-with-golang-on-k8s/models/mresponse"

	"github.com/asaskevich/govalidator"
)

// ValidateRequest validates a resuest from the API clients
func ValidateRequest(obj interface{}) *mresponse.ErrorResponse {

	_, er := govalidator.ValidateStruct(obj)

	if er != nil {
		e := mresponse.ErrorResponse{}
		e.HttpCode = 400
		e.Code = INVALID_REQUEST
		e.Response = ResponseMessageErrorsMapper[INVALID_REQUEST]

		details := []mresponse.ErrorDetail{}
		for k, v := range govalidator.ErrorsByField(er) {
			errorDetail := mresponse.ErrorDetail{
				Property: k,
				Message:  v,
			}
			details = append(details, errorDetail)
		}

		e.Errors = details

		return &e
	}

	return nil
}
