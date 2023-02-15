package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegative_Customer(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("NegativeTest_CustomerID", func(t *testing.T) {
		c := Customer{
			Name:       "Jojo",
			Email:      "jojo@gmail.com",
			CustomerID: "B6304249",
		}
		ok, err := govalidator.ValidateStruct(c)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("CustomerID: B6304249 does not validate as matches(^[LMH]\\d{7}$)"))
	})
}