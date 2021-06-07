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
