package custom

import "github.com/labstack/echo/v4"

type (
	errorMessage struct {
		Message string `json:"message"`
	}
)

func Error(pctx echo.Context, statusCode int, err error) error {
	return pctx.JSON(statusCode, &errorMessage{Message: err.Error()})
}