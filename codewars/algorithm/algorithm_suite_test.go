package algorithm

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAlgorithm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Algorithm Suite")
}

var _ = Describe("Breaking simple chockolate bars", func() {
	It("should return number of breaks needed", func() {
		Expect(BreakChocolate(5, 5)).To(Equal(24))
		Expect(BreakChocolate(1, 1)).To(Equal(0))
	})
})

func dotest(listArt []string, listCat []string, exp string) {
	var ans = StockList(listArt, listCat)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests StockList", func() {

	It("should handle basic cases", func() {
		var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
		var c = []string{"A", "B", "C", "D"}
		dotest(b, c, "(A : 0) - (B : 1290) - (C : 515) - (D : 600)")

		b = []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
		c = []string{"A", "B"}
		dotest(b, c, "(A : 200) - (B : 1140)")

		b = []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
		c = []string{"E", "F"}
		dotest(b, c, "(E : 0) - (F : 0)")
	})
})

var _ = Describe("Test Example", func() {
	It("'add' should return the two numbers added together!", func() {
		Expect(Arithmetic(8, 2, "add")).To(Equal(10))
	})
	It("'subtract' should return a minus b!", func() {
		Expect(Arithmetic(8, 2, "subtract")).To(Equal(6))
	})
	It("'multiply' should return a multiplied by b!", func() {
		Expect(Arithmetic(8, 2, "multiply")).To(Equal(16))
	})
	It("'divide' should return a divided by b!", func() {
		Expect(Arithmetic(8, 2, "divide")).To(Equal(4))
	})
})
