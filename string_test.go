package validator_test

import (
	"errors"

	"github.com/gungun974/validator"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type FakeTrueStringValidator struct{}

func (v FakeTrueStringValidator) Validate(_ string) error {
	return nil
}

type FakeErrorStringValidator struct{}

func (v FakeErrorStringValidator) Validate(_ string) error {
	return errors.New("this String validator always fail")
}

func stringValidatorTests() {
	Describe("ValidateString", func() {
		It("should return an String", func() {
			// arrange
			value := "5"

			// act
			result, err := validator.ValidateString(value, validator.StringValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal("5"))
		})

		It("should not return an String when input is garbage", func() {
			// arrange
			value := 42

			// act
			_, err := validator.ValidateString(value, validator.StringValidators{})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is not a string"))
		})

		It("should return an String when all rules are satisfied", func() {
			// arrange
			value := "result"

			// act
			result, err := validator.ValidateString(value, validator.StringValidators{
				FakeTrueStringValidator{},
				FakeTrueStringValidator{},
				FakeTrueStringValidator{},
			})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal("result"))
		})

		It("should not return an String when one rule is not satisfied", func() {
			// arrange
			value := "invalid"

			// act
			_, err := validator.ValidateString(value, validator.StringValidators{
				FakeTrueStringValidator{},
				FakeErrorStringValidator{},
				FakeTrueStringValidator{},
			})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("this String validator always fail"))
		})

		Describe("StringEmailValidator", func() {
			It("should not return a String when input is not a supported golang email", func() {
				// arrange
				value := "hoi"

				// act
				_, err := validator.ValidateString(value, validator.StringValidators{
					validator.StringEmailValidator{},
				})

				// assert
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(Equal("value is not an email"))
			})

			It("should return a String when input is a supported golang email", func() {
				// arrange
				value := "john@doe.com"

				// act
				result, err := validator.ValidateString(value, validator.StringValidators{
					validator.StringEmailValidator{},
				})

				// assert
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).To(Equal("john@doe.com"))
			})

			DescribeTable("should return a String when input is a weird but valid email account",
				func(email string) {
					// arrange
					value := email

					// act
					result, err := validator.ValidateString(value, validator.StringValidators{
						validator.StringEmailValidator{},
					})

					// assert
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result).To(Equal(email))
				},
				Entry("t'challa@ is a valid email", "t'challa@me.com"),
				Entry("rocket+groot@ is a valid email", "rocket+groot@gmail.com"),
				Entry("\"Bruce Banner\"@ is a valid email", "\"Bruce Banner\"@batman.com"),
				Entry("'@ is a valid email", "'@hoi.com"),
				Entry("-@ is a valid email", "-@github.com"),
				Entry("_@ is a valid email", "_@golang.com"),
			)
		})
	})
}
