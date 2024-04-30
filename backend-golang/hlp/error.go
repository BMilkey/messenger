package hlp

import (
	"fmt"
	"os"
	"reflect"
)

func ProcessError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func CheckStructEmptyFields(value reflect.Value, prefix string) error {
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := value.Type().Field(i).Name
		fieldValue := field.Interface()

		if field.Kind() == reflect.Struct {
			err := CheckStructEmptyFields(field, prefix+fieldName+".")
			if err != nil {
				return err
			}
		} else if fieldValue == "" {
			return fmt.Errorf("%s is empty", fieldName)
		}
	}
	return nil
}
