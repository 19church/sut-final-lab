package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("PositiveTest", func(t *testing.T) {
		c := Customer{
			Name:       "Jojo",
			Email:      "jojo@gmail.com",
			CustomerID: "L6304249",
		}
		ok, err := govalidator.ValidateStruct(c)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
