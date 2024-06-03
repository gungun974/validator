package validator_test

import (
	"errors"

	"github.com/gungun974/validator"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type FakeTrueFloatValidator struct{}

func (v FakeTrueFloatValidator) Validate(_ float64) error {
	return nil
}

type FakeErrorFloatValidator struct{}

func (v FakeErrorFloatValidator) Validate(_ float64) error {
	return errors.New("this float validator always fail")
}

func floatValidatorTests() {
	Describe("ValidateFloat", func() {
		It("should return a float", func() {
			// arrange
			value := 5.0

			// act
			result, err := validator.ValidateFloat(value, validator.FloatValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(5.0))
		})

		It("should return convert int to float", func() {
			// arrange
			value := 5

			// act
			result, err := validator.ValidateFloat(value, validator.FloatValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(5.0))
		})

		It("should not return a float when input is garbage", func() {
			// arrange
			value := "garbage"

			// act
			_, err := validator.ValidateFloat(value, validator.FloatValidators{})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is not a number"))
		})

		It("should return a float when all rules are satisfied", func() {
			// arrange
			value := 42.0

			// act
			result, err := validator.ValidateFloat(value, validator.FloatValidators{
				FakeTrueFloatValidator{},
				FakeTrueFloatValidator{},
				FakeTrueFloatValidator{},
			})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(42.0))
		})

		It("should not return a float when one rule is not satisfied", func() {
			// arrange
			value := 42.0

			// act
			_, err := validator.ValidateFloat(value, validator.FloatValidators{
				FakeTrueFloatValidator{},
				FakeErrorFloatValidator{},
				FakeTrueFloatValidator{},
			})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("this float validator always fail"))
		})
	})

	Describe("CoerceAndValidateFloat", func() {
		It("should return a float", func() {
			// arrange
			value := 5.0

			// act
			result, err := validator.CoerceAndValidateFloat(value, validator.FloatValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(5.0))
		})

		It("should return a float when input is a string with decimal", func() {
			// arrange
			value := "5.1"

			// act
			result, err := validator.CoerceAndValidateFloat(value, validator.FloatValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(5.1))
		})

		It("should return a float when input is a string with no decimal", func() {
			// arrange
			value := "5"

			// act
			result, err := validator.CoerceAndValidateFloat(value, validator.FloatValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(5.0))
		})

		It("should not return a float when input is garbage", func() {
			// arrange
			value := "garbage"

			// act
			_, err := validator.CoerceAndValidateFloat(value, validator.FloatValidators{})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is not a number"))
		})

		It("should return a float when all rules are satisfied", func() {
			// arrange
			value := "42.0"

			// act
			result, err := validator.CoerceAndValidateFloat(value, validator.FloatValidators{
				FakeTrueFloatValidator{},
				FakeTrueFloatValidator{},
				FakeTrueFloatValidator{},
			})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(42.0))
		})

		It("should not return a float when one rule is not satisfied", func() {
			// arrange
			value := "42.0"

			// act
			_, err := validator.CoerceAndValidateFloat(value, validator.FloatValidators{
				FakeTrueFloatValidator{},
				FakeErrorFloatValidator{},
				FakeTrueFloatValidator{},
			})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("this float validator always fail"))
		})
	})
}
