package utils

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func Bdate(fl validator.FieldLevel) bool {
	bdate := fl.Field().String()

	_, err := time.Parse("2006-01-02", bdate)
	return err == nil
}
