package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegative_Name(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("PositiveTest", func(t *testing.T) {
		c := Customer{
			Name:       "",
			Email:      "jojo@gmail.com",
			CustomerID: "L6304249",
		}
		ok, err := govalidator.ValidateStruct(c)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name can not be blank"))
	})
}
