package rest_errors


import (
	"net/http"
	"errors"
	"fmt"
	"encoding/json"
)

type RestErr interface {
        Message() string
        Status() int
        Error() string
        Causes() []interface{}
}

type restErr struct {
	message string        `json:"message"`
	status  int           `json:"code"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e restErr) Message() string {
	return e.message
}

func (e restErr) Status() int {
	return e.status
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %y ]",
                e.message, e.status, e.error, e.causes)
}

func (e restErr) Causes() []interface{} {
	return e.causes
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr {
		message: message,
		status: status,
		error: err,
		causes: causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) RestErr {
	return restErr {
		message: message,
		status: http.StatusBadRequest,
		error: "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
        return restErr {
                message: message,
                status: http.StatusNotFound,
                error: "not_found",
        }
}

func NewUnauthorizedError(message string) RestErr {
        return restErr {
                message: message,
                status: http.StatusUnauthorized,
                error: "unauthorized",
        }
}

func NewConflictError(message string) RestErr {
	return restErr {
		message: message,
		status: http.StatusConflict,
		error: "conflict",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr {
		message: message,
		status: http.StatusInternalServerError,
		error: "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}
