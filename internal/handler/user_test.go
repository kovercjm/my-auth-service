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

var _ = Describe("Test handler [User]", func() {
	var engine *gin.Engine

	BeforeEach(func() {
		Ω(fx.New(fx.Module(
			"handler suite test",
			fx.Provide(server.NewLogger, repository.New, handler.New, server.New),
			fx.WithLogger(kFx.FxLogger),
			fx.Invoke(func(srv *server.Server) { engine = srv.GinEngine }),
		)).Start(context.Background())).Should(Succeed())

		_ = request(engine, "POST", "/users/", `{"name":"Doe","password":"123456"}`)
	})

	Context("When create user success", func() {
		var response *httptest.ResponseRecorder
		BeforeEach(func() {
			response = request(engine, "POST", "/users/", `{"name":"John","password":"123456"}`)
		})

		It("should response with status code 200", func() {
			Ω(response.Code).Should(Equal(http.StatusOK))
		})

		It("should response with id in body", func() {
			Ω(response.Body.String()).Should(Equal(`{"id":"John"}`))
		})
	})

	Context("When repeatedly create user, the second call should fail", func() {
		var response *httptest.ResponseRecorder
		BeforeEach(func() {
			response = request(engine, "POST", "/users/", `{"name":"Doe","password":"123456"}`)
		})

		It("should response with status code 409", func() {
			Ω(response.Code).Should(Equal(http.StatusConflict))
		})

		It("should response with message in body", func() {
			Ω(response.Body.String()).Should(Equal(`{"code":"Conflict","message":"User already exists"}`))
		})
	})

	Context("When delete user success", func() {
		var response *httptest.ResponseRecorder
		BeforeEach(func() {
			response = request(engine, "DELETE", "/users/Doe", "")
		})

		It("should response with status code 204", func() {
			Ω(response.Code).Should(Equal(http.StatusNoContent))
		})

		It("should have no response body", func() {
			Ω(response.Body.Len()).Should(Equal(0))
		})
	})

	Context("When repeatedly delete user, the second call should fail", func() {
		var response *httptest.ResponseRecorder
		BeforeEach(func() {
			_ = request(engine, "DELETE", "/users/Doe", "")
			response = request(engine, "DELETE", "/users/Doe", "")
		})

		It("should response with status code 404", func() {
			Ω(response.Code).Should(Equal(http.StatusNotFound))
		})

		It("should response with message in body", func() {
			Ω(response.Body.String()).Should(Equal(`{"code":"Not Found","message":"User not exists"}`))
		})
	})
})
