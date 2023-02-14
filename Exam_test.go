package ExamTest

import (

	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Employee struct {

	gorm.Model
	Name string `json:"Name" valid:"required~Name cannot be blank"`
	NumberID string `json:"NumerID" valid:"matches(^[E]\\d{6}$)"`
}

func TestNameValiDate (t *testing.T){
	g := NewGomegaWithT(t)
	n := Employee {
		Name: "",
		NumberID: "E123456",
	}

	ok, err := govalidator.ValidateStruct(n)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name cannot be blank"))

}

func TestNumberIDValiDate (t *testing.T){
	g := NewGomegaWithT(t)
	n := Employee {
		Name: "Emy",
		NumberID: "E123",
	}

	ok, err := govalidator.ValidateStruct(n)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("NumerID: E123 does not validate as matches(^[E]\\d{6}$)"))

}

func TestTrueValiDate (t *testing.T){
	g := NewGomegaWithT(t)
	n := Employee {
		Name: "Emy",
		NumberID: "E123456",
	}

	ok, err := govalidator.ValidateStruct(n)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())

}