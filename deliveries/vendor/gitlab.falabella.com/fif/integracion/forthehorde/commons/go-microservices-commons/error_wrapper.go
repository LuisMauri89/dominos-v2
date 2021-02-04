package commons

import (
	"net/http"
)

//DOCSURL variable para setear DOCS de response
var DOCSURL string = ""

//Error2WrapperFunc tipo de funcion error2Wrapper
type Error2WrapperFunc func(err error) (status int, errBody interface{})

//Error2WrapperMiddleware middleware de error2WrapperFunc
type Error2WrapperMiddleware func(Error2WrapperFunc) Error2WrapperFunc

//ErrorWrapper wrapper de response error
type ErrorWrapper struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Docs    []string `json:"docs,omitempty"`
}

//NewErrorWrapper crea un nuevo error wrapper
func NewErrorWrapper(code, message string, docs ...string) ErrorWrapper {

	if len(docs) < 1 {
		docs = []string{DOCSURL}
	}

	return ErrorWrapper{
		Code:    code,
		Message: message,
		Docs:    docs,
	}
}

//Error2Wrapper Convierte error a status code y error wrapper
func Error2Wrapper(err error) (status int, errBody interface{}) {
	switch err.(type) {
	case *InputError:
		return http.StatusBadRequest, NewErrorWrapper("400", err.Error())
	case *NotFoundError:
		return http.StatusNotFound, NewErrorWrapper("404", err.Error())
	case *GatewayError:
		return http.StatusServiceUnavailable, NewErrorWrapper("503", err.Error())
	case *BackendCodedError:
		beErr := err.(*BackendCodedError)
		return http.StatusBadGateway, NewErrorWrapper(beErr.Code, beErr.Message)
	case *CustomError:
		ceErr := err.(*CustomError)
		return ceErr.StatusCode, NewErrorWrapper(ceErr.Code, ceErr.Message)
	default:
		return http.StatusInternalServerError, NewErrorWrapper("500", err.Error())
	}
}
