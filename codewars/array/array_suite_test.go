package array

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Array Suite")
}
