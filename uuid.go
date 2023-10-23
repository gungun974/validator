package validator

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func ValidateMapUUID(name string, value map[string]any) (uuid.UUID, error) {
	rawValue, ok := value[name]
	if !ok {
		return uuid.UUID{}, errors.New(fmt.Sprintf("missing key \"%v\"", name))
	}
	return ValidateUUID(rawValue)
}

func ValidateUUID(value any) (uuid.UUID, error) {
	uuidValue, uuidOk := value.(uuid.UUID)

	if uuidOk {
		return uuidValue, nil
	}

	stringValue, stringOk := value.(string)

	if !stringOk {
		return uuid.UUID{}, errors.New("value is not a string or UUID")
	}

	uuidValue, err := uuid.Parse(stringValue)
	if err != nil {
		return uuid.UUID{}, errors.New("value is an invalid UUID")
	}

	return uuidValue, nil
}
