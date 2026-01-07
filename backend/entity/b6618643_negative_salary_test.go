package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestSalary(t *testing.T) {
	g := NewGomegaWithT(t)

	employees := Employees{
		Name:         "Holger Rune",
		Salary:       1000,
		EmployeeCode: "AB-1234",
	}

	ok, err := govalidator.ValidateStruct(employees)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))

}