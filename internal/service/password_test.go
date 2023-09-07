package service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my-auth-service/internal/service"
)

var _ = Describe("Service password test", func() {
	Context("Should success", func() {
		It("Should success", func() {
			Ω(service.Hash("password")).Should(Succeed())
		})
		It("Should failed", func() {
			hash, err := service.Hash("password")
			Ω(err).Should(Succeed())
			Ω(hash).ShouldNot(Equal("password"))
		})
	})
})
