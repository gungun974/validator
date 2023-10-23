package validator_test

import (
	"github.com/gungun974/validator"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func uuidValidatorTests() {
	Describe("ValidateUUID", func() {
		It("should return an UUID when input is an UUID", func() {
			// arrange
			value := uuid.New()

			// act
			result, err := validator.ValidateUUID(value)

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(value))
		})

		It("should return an UUID when input string is a valid UUID", func() {
			// arrange
			target := uuid.New()
			value := target.String()

			// act
			result, err := validator.ValidateUUID(value)

			// assert
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).To(Equal(target))
		})

		It("should not return an UUID when input string is not a valid UUID", func() {
			// arrange
			value := "myFakeUUID"

			// act
			_, err := validator.ValidateUUID(value)

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is an invalid UUID"))
		})

		It("should not return an UUID when input is garbage", func() {
			// arrange
			value := 42

			// act
			_, err := validator.ValidateUUID(value)

			// assert
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("value is not a string or UUID"))
		})
	})
}
