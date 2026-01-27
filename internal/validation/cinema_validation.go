package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func TranslateValidationError(err error) error {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return err
	}

	for _, fe := range ve {
		switch fe.Field() {
		case "Name":
			switch fe.Tag() {
			case "required":
				return fmt.Errorf("Tên rạp chiếu không được để trống!")
			case "min":
				return fmt.Errorf("Tên rạp chiếu phải có ít nhất %s ký tự!", fe.Param())
			case "max":
				return fmt.Errorf("Tên rạp chiếu không được vượt quá %s ký tự!", fe.Param())
			}
		}
	}

	return fmt.Errorf("Dữ liệu không hợp lệ!")
}
