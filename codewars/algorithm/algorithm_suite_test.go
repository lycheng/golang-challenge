package algorithm

import (
	"math"
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

var _ = Describe("Example test", func() {
	It("Example test case", func() {
		Expect(FindOutlier([]int{2, 6, 8, -10, 3})).To(Equal(3))
		Expect(FindOutlier([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781})).To(Equal(206847684))
		Expect(FindOutlier([]int{math.MaxInt32, 0, 1})).To(Equal(0))
	})
})

var _ = Describe("Test Vasya Clerk", func() {
	It("should work for some basic cases", func() {
		Expect(Tickets([]int{25, 25, 50})).To(Equal("YES"))
		Expect(Tickets([]int{25, 100})).To(Equal("NO"))
	})
})

func dotest2(v1, v2, g int, exp [3]int) {
	var ans = Race(v1, v2, g)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Tortoise Racing", func() {
	It("should handle basic cases", func() {
		dotest2(720, 850, 70, [3]int{0, 32, 18})
		dotest2(820, 81, 550, [3]int{-1, -1, -1})
		dotest2(80, 91, 37, [3]int{3, 21, 49})
	})
})

func dotest3(n int, exp [][]int) {
	var ans = Solequa(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest3(5, [][]int{{3, 1}})
		dotest3(12, [][]int{{4, 1}})
		dotest3(13, [][]int{{7, 3}})
		dotest3(9005, [][]int{{4503, 2251}, {903, 449}})
		dotest3(9008, [][]int{{1128, 562}})
		dotest3(90002, [][]int{})
	})
})

func dotest4(k, start, nd int, exp []int) {
	var ans = Gap(k, start, nd)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Gap In Primes", func() {
	It("should handle basic cases", func() {
		dotest4(2, 100, 110, []int{101, 103})
		dotest4(4, 100, 110, []int{103, 107})
		dotest4(6, 100, 110, nil)
		dotest4(8, 300, 400, []int{359, 367})
		dotest4(10, 300, 400, []int{337, 347})
	})
})
