package main

import (
	"github.com/Kalimaha/ginkgo/reporter"
	"github.com/aws/aws-lambda-go/events"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestHello(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Hello API Suite", []Reporter{reporter.New()})
}

var _ = Describe("Hello", func() {
	var message *events.APIGatewayProxyResponse

	BeforeEach(func() {
		message, _ = handler(events.APIGatewayProxyRequest{})
	})

	It("returns HTTP 200 - OK", func() {
		Expect(message.StatusCode).To(Equal(200))
	})

	It("returns 'Ciao!' in the body", func() {
		Expect(message.Body).To(Equal("Ciao!"))
	})
})
