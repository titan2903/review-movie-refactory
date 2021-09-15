package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Data interface{} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

type ResponseError struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
	ErrorValidation interface{} `json:"error_validation"`
}

type ResponseLogin struct {
	Data interface{} `json:"data"`
}

type ResponseGeneral struct {
	Data interface{} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {

	jsonResponse := Response{
		Data: data,
		Message: message,
		Status:  status,
	}

	return jsonResponse
}

func ApiResponseGeneral(message string, status string, data interface{}) ResponseGeneral {

	jsonResponse := ResponseGeneral{
		Data: data,
		Message: message,
		Status:  status,
	}

	return jsonResponse
}

func ApiResponseError(message string, code int, status string, error interface{}) ResponseError {
	jsonResponse := ResponseError{
		Code: code,
		Message: message,
		Status:  status,
		ErrorValidation: error,
	}

	return jsonResponse
}

func ApiResponseLogin(data interface{}) ResponseLogin {
	jsonResponse := ResponseLogin{
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) { //!mnegubah terlebih dahulu menjadi validation error
		errors = append(errors, e.Error()) //! menambahkan nilai errornya
	}

	return errors
}