package customvalidator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetErrorsMessage(err error) []map[string]string {
	listErrMsg := []map[string]string{}
	for _, v := range err.(validator.ValidationErrors) {
		errMsg := map[string]string{"message": fmt.Sprintf("field '%v' have tag %v", v.Field(), v.Tag())}
		listErrMsg = append(listErrMsg, errMsg)
	}
	return listErrMsg
}
