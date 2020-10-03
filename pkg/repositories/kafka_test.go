package repositories

import (
	"github.com/Kalimaha/ginkgo/reporter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestHello(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Ciao Suite", []Reporter{reporter.New()})
}

var _ = Describe("Ciao", func() {
	var (
		message string
	)

	BeforeEach(func() {
		message = Ciao()
	})

	It("returns 'Ciao!'", func() {
		Expect(message).To(Equal("Ciao!"))
	})
})
