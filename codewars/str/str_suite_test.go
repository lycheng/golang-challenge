package str

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Str Suite")
}

var _ = Describe("Basic tests", func() {
	t := [...][2]string{
		{"test case", "TestCase"},
		{"camel case method", "CamelCaseMethod"},
		{"say hello ", "SayHello"},
		{" camel case word", "CamelCaseWord"},
		{"", ""},
	}

	for _, v := range t {
		It(fmt.Sprintf("Testing input \"%s\"", v[0]), func() { Expect(CamelCase(v[0])).To(Equal(v[1])) })
	}
})

func dotest(st string, exp []string) {
	var ans = Capitalize(st)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Example tests", func() {
	It("It should work for basic tests", func() {
		dotest("abcdef", []string{"AbCdEf", "aBcDeF"})
		dotest("codewars", []string{"CoDeWaRs", "cOdEwArS"})
		dotest("abracadabra", []string{"AbRaCaDaBrA", "aBrAcAdAbRa"})
		dotest("codewarriors", []string{"CoDeWaRrIoRs", "cOdEwArRiOrS"})
		dotest("indexinglessons", []string{"InDeXiNgLeSsOnS", "iNdExInGlEsSoNs"})
		dotest("codingisafunactivity", []string{"CoDiNgIsAfUnAcTiViTy", "cOdInGiSaFuNaCtIvItY"})
	})
})
