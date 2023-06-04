package utils

import (
	"net/http"

	"github.com/gustavoteixeira8/url-shortener/src/cerrors"
)

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

func GetErrorInHttpFormat(err error) *HttpResponse {
	if err == nil {
		return nil
	}

	switch err {
	case cerrors.ErrCantPingUrl:
		fallthrough
	case cerrors.ErrInvalidURLFormat:
		fallthrough
	case cerrors.ErrNameIsRequired:
		fallthrough
	case cerrors.ErrUrlIsRequired:
		fallthrough
	case cerrors.ErrUrlAlreadyExists:
		fallthrough
	case cerrors.ErrNameAlreadyExists:
		return BadRequest(err.Error(), nil)
	case cerrors.ErrUrlNotFound:
		return NotFound(err.Error(), nil)
	}

	return InternalError(err.Error(), nil)
}
