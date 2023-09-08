package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	kFx "github.com/kovercjm/tool-go/dependency_injection/fx"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"

	"my-auth-service/internal/handler"
	"my-auth-service/internal/repository"
	"my-auth-service/internal/server"
)

var _ = Describe("Test handler [Auth]", func() {
	var engine *gin.Engine

	BeforeEach(func() {
		Ω(fx.New(fx.Module(
			"handler suite test",
			fx.Provide(server.NewLogger, repository.New, handler.New, server.New),
			fx.WithLogger(kFx.FxLogger),
			fx.Invoke(func(srv *server.Server) { engine = srv.GinEngine }),
		)).Start(context.Background())).Should(Succeed())

		_ = request(engine, "POST", "/users/", `{"name":"John","password":"123456"}`)
		_ = request(engine, "POST", "/roles/", `{"name":"platform-user"}`)
	})

	Context("When grant user role success", func() {
		var response *httptest.ResponseRecorder
		BeforeEach(func() {
			response = request(engine, "POST", "/users/John/roles/platform-user", "")
		})

		It("should response with status code 204", func() {
			Ω(response.Code).Should(Equal(http.StatusNoContent))
		})

		It("should have no response body", func() {
			Ω(response.Body.Len()).Should(Equal(0))
		})
	})

	Context("When repeatedly grant user role success", func() {
		var response *httptest.ResponseRecorder

		BeforeEach(func() {
			_ = request(engine, "POST", "/users/John/roles/platform-user", "")
			response = request(engine, "POST", "/users/John/roles/platform-user", "")
		})

		It("should response with status code 204", func() {
			Ω(response.Code).Should(Equal(http.StatusNoContent))
		})

		It("should have no response body", func() {
			Ω(response.Body.Len()).Should(Equal(0))
		})
	})

	Context("When list user roles success", func() {
		var (
			response *httptest.ResponseRecorder
			token    string
		)

		BeforeEach(func() {
			response = request(engine, "POST", "/auth/login", `{"name":"John","password":"123456"}`)
			resp := response.Body.String()
			var body struct {
				Token string `json:"token"`
			}
			Ω(json.Unmarshal([]byte(resp), &body)).Should(Succeed())
			token = body.Token

			_ = request(engine, "POST", "/users/John/roles/platform-user", "")
			response = requestWithAuthorization(engine, "GET", "/users/me/roles", "", token)
		})

		It("should response with status code 200", func() {
			Ω(response.Code).Should(Equal(http.StatusOK))
		})

		It("response with message in body", func() {
			Ω(response.Body.String()).Should(Equal(`{"roles":[{"id":"platform-user","name":"platform-user"}]}`))
		})
	})

	Context("When get user role success", func() {
		var (
			response *httptest.ResponseRecorder
			token    string
		)

		BeforeEach(func() {
			response = request(engine, "POST", "/auth/login", `{"name":"John","password":"123456"}`)
			resp := response.Body.String()
			var body struct {
				Token string `json:"token"`
			}
			Ω(json.Unmarshal([]byte(resp), &body)).Should(Succeed())
			token = body.Token

			_ = request(engine, "POST", "/users/John/roles/platform-user", "")
			response = requestWithAuthorization(engine, "GET", "/users/me/roles/platform-user", "", token)
		})

		It("should response with status code 204", func() {
			Ω(response.Code).Should(Equal(http.StatusNoContent))
		})

		It("response with message in body", func() {
			Ω(response.Body.Len()).Should(Equal(0))
		})
	})

	Context("When user sign in success", func() {
		var response *httptest.ResponseRecorder

		BeforeEach(func() {
			response = request(engine, "POST", "/auth/login", `{"name":"John","password":"123456"}`)
		})

		It("should response with status code 200", func() {
			Ω(response.Code).Should(Equal(http.StatusOK))
		})

		It("should response with token in body", func() {
			resp := response.Body.String()
			var body struct {
				token string `json:"token"`
			}
			Ω(json.Unmarshal([]byte(resp), &body)).Should(Succeed())
		})
	})

	Context("When user sign out success", func() {
		var (
			response *httptest.ResponseRecorder
			token    string
		)

		BeforeEach(func() {
			response = request(engine, "POST", "/auth/login", `{"name":"John","password":"123456"}`)
			resp := response.Body.String()
			var body struct {
				Token string `json:"token"`
			}
			Ω(json.Unmarshal([]byte(resp), &body)).Should(Succeed())
			token = body.Token

			response = requestWithAuthorization(engine, "POST", "/auth/logout", "", token)
		})

		It("should response with status code 204", func() {
			Ω(response.Code).Should(Equal(http.StatusNoContent))
		})

		It("should have no response body", func() {
			Ω(response.Body.Len()).Should(Equal(0))
		})
	})
})
