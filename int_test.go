package validator_test

import (
	"errors"

	"github.com/gungun974/validator"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type FakeTrueIntValidator struct{}

func (v FakeTrueIntValidator) Validate(_ int) error {
	return nil
}

type FakeErrorIntValidator struct{}

func (v FakeErrorIntValidator) Validate(_ int) error {
	return errors.New("this int validator always fail")
}

func intValidatorTests() {
	Describe("ValidateInt", func() {
		It("should return an int", func() {
			// arrange
			value := 5

			// act
			result, err := validator.ValidateInt(value, validator.IntValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(5))
		})

		It("should return an int when input is a float64 with no decimal", func() {
			// arrange
			value := 5.0

			// act
			result, err := validator.ValidateInt(value, validator.IntValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(5))
		})

		It("should not return an int when input is a float64 with decimal", func() {
			// arrange
			value := 5.1

			// act
			_, err := validator.ValidateInt(value, validator.IntValidators{})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is not an int"))
		})

		It("should not return an int when input is garbage", func() {
			// arrange
			value := "garbage"

			// act
			_, err := validator.ValidateInt(value, validator.IntValidators{})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is not a number"))
		})

		It("should return an int when all rules are satisfied", func() {
			// arrange
			value := 42

			// act
			result, err := validator.ValidateInt(value, validator.IntValidators{
				FakeTrueIntValidator{},
				FakeTrueIntValidator{},
				FakeTrueIntValidator{},
			})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(42))
		})

		It("should not return an int when one rule is not satisfied", func() {
			// arrange
			value := 42

			// act
			_, err := validator.ValidateInt(value, validator.IntValidators{
				FakeTrueIntValidator{},
				FakeErrorIntValidator{},
				FakeTrueIntValidator{},
			})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("this int validator always fail"))
		})
	})

	Describe("CoerceAndValidateInt", func() {
		It("should return an int", func() {
			// arrange
			value := 5

			// act
			result, err := validator.CoerceAndValidateInt(value, validator.IntValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(5))
		})

		It("should return an int when input is a valid string", func() {
			// arrange
			value := "5"

			// act
			result, err := validator.CoerceAndValidateInt(value, validator.IntValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(5))
		})

		It("should return an int when input is a string with no decimal", func() {
			// arrange
			value := "5.0"

			// act
			result, err := validator.CoerceAndValidateInt(value, validator.IntValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(5))
		})

		It("should not return an int when input is a string with decimal", func() {
			// arrange
			value := "5.1"

			// act
			_, err := validator.CoerceAndValidateInt(value, validator.IntValidators{})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is not an int"))
		})

		It("should not return an int when input is garbage", func() {
			// arrange
			value := "garbage"

			// act
			_, err := validator.CoerceAndValidateInt(value, validator.IntValidators{})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is not a number"))
		})

		It("should return an int when all rules are satisfied", func() {
			// arrange
			value := "42"

			// act
			result, err := validator.CoerceAndValidateInt(value, validator.IntValidators{
				FakeTrueIntValidator{},
				FakeTrueIntValidator{},
				FakeTrueIntValidator{},
			})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(42))
		})

		It("should not return an int when one rule is not satisfied", func() {
			// arrange
			value := "42"

			// act
			_, err := validator.CoerceAndValidateInt(value, validator.IntValidators{
				FakeTrueIntValidator{},
				FakeErrorIntValidator{},
				FakeTrueIntValidator{},
			})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("this int validator always fail"))
		})
	})
}
