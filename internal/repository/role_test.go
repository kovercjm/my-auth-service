package repository_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my-auth-service/internal/domain"
	"my-auth-service/internal/middleware/dependency"
)

var _ = Describe("Repository role test", func() {
	Context("When create role", func() {
		It("First time create should success", func() {
			Ω(r.CreateRole(&domain.Role{
				ID:   gofakeit.Email(),
				Name: gofakeit.Name(),
			})).Should(Succeed())
		})
		It("Duplicated create should failed", func() {
			id := gofakeit.Email()
			Ω(r.CreateRole(&domain.Role{
				ID:   id,
				Name: gofakeit.Name(),
			})).Should(Succeed())
			Ω(r.CreateRole(&domain.Role{
				ID:   id,
				Name: gofakeit.Name(),
			})).Should(Equal(dependency.AlreadyExistsError))
		})
	})

	Context("When delete role", func() {
		It("Should success", func() {
			roleID := gofakeit.Email()
			Ω(r.CreateRole(&domain.Role{
				ID:   roleID,
				Name: gofakeit.Name(),
			})).Should(Succeed())
			Ω(r.DeleteRole(roleID)).Should(Succeed())
		})
	})
})
