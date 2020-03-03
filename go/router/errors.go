package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiError struct {
	Code    errorCode `json:"code"`
	Message string    `json:"message"`
}

func (c *apiError) Error() string {
	return fmt.Sprintf("%s: %s", c.Code, c.Message)
}

type errorCode string

const (
	malformedRequestBody   = errorCode("MALFORMED_REQUEST_BODY")
	tooManySolutions       = errorCode("TOO_MANY_SOLUTIONS")
	noSolution             = errorCode("NO_SOLUTION")
	errorCodeInternalError = errorCode("INTERNAL_ERROR")
)

func bindError(err error) *apiError {
	if asJSONError, ok := err.(*json.UnmarshalTypeError); ok {
		return &apiError{
			Code:    malformedRequestBody,
			Message: fmt.Sprintf("Request body has malformed JSON field '%s'", asJSONError.Field),
		}
	}

	return &apiError{
		Code:    malformedRequestBody,
		Message: "Malformed request body",
	}
}

func handleErrors(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		if apiError, ok := err.Err.(*apiError); ok {
			c.JSON(http.StatusBadRequest, apiError)
			return
		}
	}

	if len(c.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, &apiError{
			Code:    errorCodeInternalError,
			Message: "Internal server error",
		})
	}
}
