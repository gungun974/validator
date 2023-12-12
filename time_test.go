package validator_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gungun974/validator"
)

type FakeTrueTimeValidator struct{}

func (v FakeTrueTimeValidator) Validate(_ time.Time) error {
	return nil
}

type FakeErrorTimeValidator struct{}

func (v FakeErrorTimeValidator) Validate(_ time.Time) error {
	return errors.New("this Time validator always fail")
}

func timeValidatorTests() {
	Describe("ValidateTime", func() {
		It("should return a Time", func() {
			// arrange
			value := time.Date(2023, 12, 24, 5, 0, 0, 0, time.UTC)

			// act
			result, err := validator.ValidateTime(value, validator.TimeValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(
				time.Date(2023, 12, 24, 5, 0, 0, 0, time.UTC),
			))
		})

		It("should not return an Time when input is garbage", func() {
			// arrange
			value := 42

			// act
			_, err := validator.ValidateTime(value, validator.TimeValidators{})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is not a time"))
		})

		It("should return a Time when value is an yy-mm-dd string", func() {
			// arrange
			value := "2023-12-24"

			// act
			result, err := validator.ValidateTime(value, validator.TimeValidators{})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(
				time.Date(2023, 12, 24, 0, 0, 0, 0, time.UTC),
			))
		})

		It("should return a Time when all rules are satisfied", func() {
			// arrange
			value := time.Date(2023, 12, 24, 5, 0, 0, 0, time.UTC)

			// act
			result, err := validator.ValidateTime(value, validator.TimeValidators{
				FakeTrueTimeValidator{},
				FakeTrueTimeValidator{},
				FakeTrueTimeValidator{},
			})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(
				time.Date(2023, 12, 24, 5, 0, 0, 0, time.UTC),
			))
		})

		It("should not return a Time when one rule is not satisfied", func() {
			// arrange
			value := time.Date(2023, 12, 24, 5, 0, 0, 0, time.UTC)

			// act
			_, err := validator.ValidateTime(value, validator.TimeValidators{
				FakeTrueTimeValidator{},
				FakeErrorTimeValidator{},
				FakeTrueTimeValidator{},
			})

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("this Time validator always fail"))
		})

		It("should return a Time when date is before 2023/12/24", func() {
			// arrange
			value := time.Date(2023, 12, 23, 5, 0, 0, 0, time.UTC)

			// act
			result, err := validator.ValidateTime(value, validator.TimeValidators{
				validator.TimeMaxValidator{Max: time.Date(2023, 12, 24, 0, 0, 0, 0, time.UTC)},
			})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(
				time.Date(2023, 12, 23, 5, 0, 0, 0, time.UTC),
			))
		})

		It("should return a Time when date is after 2023/12/24", func() {
			// arrange
			value := time.Date(2023, 12, 25, 5, 0, 0, 0, time.UTC)

			// act
			result, err := validator.ValidateTime(value, validator.TimeValidators{
				validator.TimeMinValidator{Min: time.Date(2023, 12, 24, 0, 0, 0, 0, time.UTC)},
			})

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(
				time.Date(2023, 12, 25, 5, 0, 0, 0, time.UTC),
			))
		})
	})
}
