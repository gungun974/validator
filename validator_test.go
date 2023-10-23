package validator_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validator")
}

var _ = Describe("Validator", func() {
	Describe("IntValidator", intValidatorTests)
	Describe("StringValidator", stringValidatorTests)
	Describe("UUIDValidator", uuidValidatorTests)
})
