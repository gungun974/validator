package validator_test

import (
	"errors"

	"github.com/gungun974/validator"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type FakeTrueBoolValidator struct{}

func (v FakeTrueBoolValidator) Validate(_ bool) error {
	return nil
}

type FakeErrorBoolValidator struct{}

func (v FakeErrorBoolValidator) Validate(_ bool) error {
	return errors.New("this bool validator always fail")
}

func boolValidatorTests() {
	Describe("ValidateBool", func() {
		It("should return a bool", func() {
			// arrange
			value := true

			// act
			result, err := validator.ValidateBool(value, validator.BoolValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(true))
		})

		It("should not return an bool when input is garbage", func() {
			// arrange
			value := "garbage"

			// act
			_, err := validator.ValidateBool(value, validator.BoolValidators{})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is not a bool"))
		})

		It("should return a bool when input string is \"true\"", func() {
			// arrange
			value := "true"

			// act
			result, err := validator.ValidateBool(value, validator.BoolValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(true))
		})

		It("should return a bool when input string is \"false\"", func() {
			// arrange
			value := "false"

			// act
			result, err := validator.ValidateBool(value, validator.BoolValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(false))
		})

		It("should return a bool when all rules are satisfied", func() {
			// arrange
			value := false

			// act
			result, err := validator.ValidateBool(value, validator.BoolValidators{
				FakeTrueBoolValidator{},
				FakeTrueBoolValidator{},
				FakeTrueBoolValidator{},
			})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(false))
		})

		It("should not return a bool when one rule is not satisfied", func() {
			// arrange
			value := true

			// act
			_, err := validator.ValidateBool(value, validator.BoolValidators{
				FakeTrueBoolValidator{},
				FakeErrorBoolValidator{},
				FakeTrueBoolValidator{},
			})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("this bool validator always fail"))
		})
	})
}
