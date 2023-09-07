package repository_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my-auth-service/internal/domain"
	"my-auth-service/internal/middleware/dependency"
)

var _ = Describe("Repository user test", func() {
	Context("When create user", func() {
		It("First time create should success", func() {
			Ω(r.CreateUser(&domain.User{
				ID:       gofakeit.Email(),
				Name:     gofakeit.Name(),
				Password: gofakeit.Password(true, true, true, true, true, 16),
			})).Should(Succeed())
		})

		It("Duplicated create should failed", func() {
			id := gofakeit.Email()
			Ω(r.CreateUser(&domain.User{
				ID:       id,
				Name:     gofakeit.Name(),
				Password: gofakeit.Password(true, true, true, true, true, 16),
			})).Should(Succeed())
			Ω(r.CreateUser(&domain.User{
				ID:       id,
				Name:     gofakeit.Name(),
				Password: gofakeit.Password(true, true, true, true, true, 16),
			})).Should(Equal(dependency.AlreadyExistsError))
		})
	})

	Context("When check user password", func() {
		It("Should success", func() {
			id, password := gofakeit.Email(), gofakeit.Password(true, true, true, true, true, 16)
			Ω(r.CreateUser(&domain.User{
				ID:       id,
				Password: password,
			})).Should(Succeed())
			Ω(r.CheckUserPassword(&domain.User{
				ID:       id,
				Password: password,
			})).Should(BeTrue())
		})
		It("Should failed", func() {
			id, password := gofakeit.Email(), gofakeit.Password(true, true, true, true, true, 16)
			Ω(r.CreateUser(&domain.User{
				ID:       id,
				Password: password,
			})).Should(Succeed())
			Ω(r.CheckUserPassword(&domain.User{
				ID:       id,
				Password: password + "something",
			})).Should(BeFalse())
		})
	})

	Context("When delete user", func() {
		It("Should success", func() {
			userID := gofakeit.Email()
			Ω(r.CreateUser(&domain.User{
				ID:       userID,
				Password: gofakeit.Password(true, true, true, true, true, 16),
			})).Should(Succeed())
			roleID := gofakeit.UUID()
			Ω(r.CreateRole(&domain.Role{
				ID:   roleID,
				Name: roleID,
			}))
			Ω(r.GrantUserRole(&domain.UserRole{
				User:  &domain.User{ID: userID},
				Roles: []*domain.Role{{ID: roleID}},
			})).Should(Succeed())
			Ω(r.DeleteUser(userID)).Should(Succeed())
		})
		It("Should failed", func() {
			Ω(r.DeleteUser(gofakeit.Email())).Should(Equal(dependency.NotFoundError))
		})
	})
})
