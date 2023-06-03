package utils

import "net/http"

type Body struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HttpResponse struct {
	Body   Body `json:"body"`
	Status int  `json:"status"`
}

func Ok(message string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Body: Body{
			Message: message,
			Data:    data,
		},
		Status: http.StatusOK,
	}
}

func Created(message string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Body: Body{
			Message: message,
			Data:    data,
		},
		Status: http.StatusCreated,
	}
}

func NotFound(message string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Body: Body{
			Message: message,
			Data:    data,
		},
		Status: http.StatusNotFound,
	}
}

func BadRequest(message string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Body: Body{
			Message: message,
			Data:    data,
		},
		Status: http.StatusBadRequest,
	}
}

func Unauthorized(message string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Body: Body{
			Message: message,
			Data:    data,
		},
		Status: http.StatusUnauthorized,
	}
}

func Forbidden(message string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Body: Body{
			Message: message,
			Data:    data,
		},
		Status: http.StatusForbidden,
	}
}

func InternalError(message string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Body: Body{
			Message: message,
			Data:    data,
		},
		Status: http.StatusInternalServerError,
	}
}
