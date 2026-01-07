package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeCode(t *testing.T) {
	g := NewGomegaWithT(t)

	employees := Employees{
		Name:         "Holger Rune",
		Salary:       30000,
		EmployeeCode: "A123",
	}

	ok, err := govalidator.ValidateStruct(employees)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))

}