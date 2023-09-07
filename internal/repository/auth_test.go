package repository_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my-auth-service/internal/domain"
)

var _ = Describe("Repository auth test", func() {
	Context("When grant user role", func() {
		It("One time create should success", func() {
			userID, roleID := gofakeit.Email(), gofakeit.UUID()
			Ω(r.CreateUser(&domain.User{
				ID:       userID,
				Name:     gofakeit.Name(),
				Password: gofakeit.Password(true, true, true, true, true, 16),
			})).Should(Succeed())
			Ω(r.CreateRole(&domain.Role{
				ID:   roleID,
				Name: gofakeit.Name(),
			}))
			Ω(r.GrantUserRole(&domain.UserRole{
				User: &domain.User{
					ID:       userID,
					Name:     gofakeit.Name(),
					Password: gofakeit.Password(true, true, true, true, true, 16),
				},
				Roles: []*domain.Role{{
					ID:   roleID,
					Name: gofakeit.Name(),
				}},
			})).Should(Succeed())
		})
		It("Duplicated grant should success", func() {
			userID, roleID := gofakeit.Email(), gofakeit.UUID()
			Ω(r.CreateUser(&domain.User{
				ID:       userID,
				Name:     gofakeit.Name(),
				Password: gofakeit.Password(true, true, true, true, true, 16),
			})).Should(Succeed())
			Ω(r.CreateRole(&domain.Role{
				ID:   roleID,
				Name: gofakeit.Name(),
			}))
			Ω(r.GrantUserRole(&domain.UserRole{
				User: &domain.User{ID: userID},
				Roles: []*domain.Role{{
					ID:   roleID,
					Name: gofakeit.Name(),
				}},
			})).Should(Succeed())
			Ω(r.GrantUserRole(&domain.UserRole{
				User: &domain.User{ID: userID},
				Roles: []*domain.Role{{
					ID:   roleID,
					Name: gofakeit.Name(),
				}},
			})).Should(Succeed())
		})
		It("Should failed", func() {
			Ω(r.GrantUserRole(nil)).ShouldNot(Succeed())
		})
	})

	Context("When get user roles", func() {
		It("Should success with no role", func() {
			userID := gofakeit.Email()
			Ω(r.CreateUser(&domain.User{
				ID:       userID,
				Name:     gofakeit.Name(),
				Password: gofakeit.Password(true, true, true, true, true, 16),
			})).Should(Succeed())
			userRoles, err := r.GetUserRoles(&domain.User{ID: userID})
			Ω(err).Should(Succeed())
			Ω(len(userRoles.Roles)).Should(Equal(0))
		})
		It("Should success with one role", func() {
			userID, roleID := gofakeit.Email(), gofakeit.UUID()
			Ω(r.CreateUser(&domain.User{
				ID:       userID,
				Name:     gofakeit.Name(),
				Password: gofakeit.Password(true, true, true, true, true, 16),
			})).Should(Succeed())
			Ω(r.CreateRole(&domain.Role{
				ID:   roleID,
				Name: gofakeit.Name(),
			}))
			Ω(r.GrantUserRole(&domain.UserRole{
				User: &domain.User{
					ID:       userID,
					Name:     gofakeit.Name(),
					Password: gofakeit.Password(true, true, true, true, true, 16),
				},
				Roles: []*domain.Role{{
					ID:   roleID,
					Name: gofakeit.Name(),
				}},
			})).Should(Succeed())
			userRoles, err := r.GetUserRoles(&domain.User{ID: userID})
			Ω(err).Should(Succeed())
			Ω(len(userRoles.Roles)).Should(Equal(1))
			Ω(userRoles.User.ID).Should(Equal(userID))
			Ω(userRoles.Roles[0].ID).Should(Equal(roleID))
		})
		It("Should failed", func() {
			Ω(r.GetUserRoles(nil)).ShouldNot(Succeed())
		})
	})

	Context("When create token", func() {
		It("Should success", func() {
			token := gofakeit.UUID()
			Ω(r.CreateToken(token)).Should(Succeed())
		})
		It("Should failed", func() {
			Ω(r.CreateToken("")).ShouldNot(Succeed())
		})
	})

	Context("When check token", func() {
		It("Should success", func() {
			token := gofakeit.UUID()
			Ω(r.CreateToken(token)).Should(Succeed())
			Ω(r.CheckToken(token)).Should(Succeed())
		})
		It("Should failed", func() {
			Ω(r.CheckToken("")).ShouldNot(Succeed())
		})
	})

	Context("When delete token", func() {
		It("Should success", func() {
			token := gofakeit.UUID()
			Ω(r.CreateToken(token)).Should(Succeed())
			Ω(r.DeleteToken(token)).Should(Succeed())
		})
		It("Should failed", func() {
			Ω(r.DeleteToken("")).ShouldNot(Succeed())
		})
	})
})
