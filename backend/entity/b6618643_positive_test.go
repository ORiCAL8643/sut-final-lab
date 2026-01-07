package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	employees := Employees{
		Name:         "Holger Rune",
		Salary:       30000,
		EmployeeCode: "AB-1234",
	}

	ok, err := govalidator.ValidateStruct(employees)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())

}
