package validator

import (
	"errors"
	"fmt"
	"time"
)

type TimeValidators []timeValidator

type timeValidator interface {
	Validate(value time.Time) error
}

type TimeMaxValidator struct {
	Max time.Time
}

func (v TimeMaxValidator) Validate(value time.Time) error {
	if value.After(v.Max) {
		return fmt.Errorf("value must before %v", v.Max)
	}
	return nil
}

type TimeMinValidator struct {
	Min time.Time
}

func (v TimeMinValidator) Validate(value time.Time) error {
	if value.Before(v.Min) {
		return fmt.Errorf("value must after %v", v.Min)
	}
	return nil
}

func ValidateMapTime(name string, value map[string]any, rules TimeValidators) (time.Time, error) {
	rawValue, ok := value[name]
	if !ok {
		return time.Time{}, fmt.Errorf("missing key \"%v\"", name)
	}
	return ValidateTime(rawValue, rules)
}

func ValidateMapTimeOrNil(
	name string,
	value map[string]any,
	rules TimeValidators,
) (*time.Time, error) {
	rawValue, ok := value[name]
	if !ok {
		return nil, nil
	}

	val, err := ValidateTime(rawValue, rules)

	return &val, err
}

func ValidateTime(value any, rules TimeValidators) (time.Time, error) {
	timeValue, timeOk := value.(time.Time)
	stringValue, stringOk := value.(string)
	if !timeOk && !stringOk {
		return time.Time{}, errors.New("value is not a time")
	}

	if stringOk {
		date, err := time.Parse("2006-01-02", stringValue)
		if err != nil {
			return time.Time{}, errors.New("value is not a time")
		}
		timeValue = date
	}

	for _, rule := range rules {
		if err := rule.Validate(timeValue); err != nil {
			return time.Time{}, err
		}
	}

	return timeValue, nil
}
