package custom

import (
	"sync"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	EchoRequest interface {
		Bind(i interface{}) error
	}
	echoRequestImpl struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

var (
	once sync.Once
	validatorInstance *validator.Validate
)

func NewEchoRequest(echoRequest echo.Context) EchoRequest {
	once.Do(func(){
		validatorInstance = validator.New()
	})
	
	return &echoRequestImpl{
		ctx:       echoRequest,
		validator: validatorInstance,
	}
}

func (req *echoRequestImpl) Bind(i interface{}) error {
	if err := req.ctx.Bind(i); err != nil {
		return err
	}

	if err := req.validator.Struct(i); err != nil {
		return err
	}

	return nil
}