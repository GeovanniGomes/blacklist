package util

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

func ValidateOrGenerateID(id *string) (*string, error) {
	newValue := uuid.NewV4().String()
	if id != nil {
		value, err := uuid.FromString(*id)
		if err != nil {
			return nil, errors.New("value id invalid")
		}
		strValue := value.String()
		return &strValue, nil
	}
	return &newValue, nil
}