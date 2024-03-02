package validator

import (
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

type GooKitValidator struct {
	StopOnError bool
}

const (
	uuidNotValidError    = "is not valid for UUID type"
	requiredErrorMessage = "needs to be on request"
	requiredUuidKey      = "required_uuid"
)

func NewGooKitValidator() *GooKitValidator {
	return &GooKitValidator{
		StopOnError: false,
	}
}

// ValidateStruct valida que la interface cumpla con los requisitos descritos por el tag validate. Si no lo hace se regresa
// una lista de errores que se puede regresar al front. Más detalles ver la documentación del paquete gookit/validate.
//
// Se cuenta con un validador custom:
//   - required_uuid
//
// Este fue creado con el proposito de validar que un tipo uuid.UUID del packete de google se encuentra en un requets. Por
// el momento aun es experimental, pero funciona para lo que se requiere.
func (c *GooKitValidator) ValidateStruct(toValidate interface{}) error {
	v := c.configuratedValidator(toValidate)
	if v.Validate() {
		return nil
	}
	return &ValidationError{Errors: v.Errors}
}

func (c *GooKitValidator) Validate(i interface{}) error {
	return c.ValidateStruct(i)
}

func (c *GooKitValidator) configuratedValidator(toValidate interface{}) *validate.Validation {
	v := validate.Struct(toValidate) //nolint:varnamelen
	v.StopOnError = c.StopOnError
	v.AddMessages(map[string]string{
		"uuid":          uuidNotValidError,
		"isUUID":        uuidNotValidError,
		"required":      requiredErrorMessage,
		requiredUuidKey: requiredErrorMessage,
	})
	v.AddValidator(requiredUuidKey, func(val any) bool {
		return val != uuid.Nil
	})
	return v
}
