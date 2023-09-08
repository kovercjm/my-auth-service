package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	kFx "github.com/kovercjm/tool-go/dependency_injection/fx"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"

	"my-auth-service/internal/handler"
	"my-auth-service/internal/repository"
	"my-auth-service/internal/server"
)

var _ = Describe("Test handler [Role]", func() {
	var engine *gin.Engine

	BeforeEach(func() {
		Ω(fx.New(fx.Module(
			"handler suite test",
			fx.Provide(server.NewLogger, repository.New, handler.New, server.New),
			fx.WithLogger(kFx.FxLogger),
			fx.Invoke(func(srv *server.Server) { engine = srv.GinEngine }),
		)).Start(context.Background())).Should(Succeed())

		_ = request(engine, "POST", "/roles/", `{"name":"platform-admin"}`)
	})

	Context("When create role success", func() {
		var response *httptest.ResponseRecorder
		BeforeEach(func() {
			response = request(engine, "POST", "/roles/", `{"name":"platform-user"}`)
		})

		It("should response with status code 200", func() {
			Ω(response.Code).Should(Equal(http.StatusOK))
		})

		It("should response with id in body", func() {
			Ω(response.Body.String()).Should(Equal(`{"id":"platform-user"}`))
		})

		AfterEach(func() {
			_ = request(engine, "DELETE", "/roles/platform-user", "")
		})
	})

	Context("When repeatedly create role, the second call should fail", func() {
		var response *httptest.ResponseRecorder
		BeforeEach(func() {
			response = request(engine, "POST", "/roles/", `{"name":"platform-admin"}`)
		})

		It("should response with status code 409", func() {
			Ω(response.Code).Should(Equal(http.StatusConflict))
		})

		It("should response with message in body", func() {
			Ω(response.Body.String()).Should(Equal(`{"code":"Conflict","message":"Role already exists"}`))
		})

		AfterEach(func() {
			_ = request(engine, "DELETE", "/roles/platform-admin", "")
		})
	})

	Context("When delete role success", func() {
		var response *httptest.ResponseRecorder
		BeforeEach(func() {
			response = request(engine, "DELETE", "/roles/platform-admin", "")
		})

		It("should response with status code 204", func() {
			Ω(response.Code).Should(Equal(http.StatusNoContent))
		})

		It("should have no response body", func() {
			Ω(response.Body.Len()).Should(Equal(0))
		})

		AfterEach(func() {
			_ = request(engine, "POST", "/roles/platform-admin", "")
		})
	})

	Context("When repeatedly delete role, the second call should fail", func() {
		var response *httptest.ResponseRecorder
		BeforeEach(func() {
			_ = request(engine, "DELETE", "/roles/platform-admin", "")
			response = request(engine, "DELETE", "/roles/platform-admin", "")
		})

		It("should response with status code 404", func() {
			Ω(response.Code).Should(Equal(http.StatusNotFound))
		})

		It("should response with message in body", func() {
			Ω(response.Body.String()).Should(Equal(`{"code":"Not Found","message":"Role not exists"}`))
		})

		AfterEach(func() {
			_ = request(engine, "POST", "/roles/platform-admin", "")
		})
	})
})
