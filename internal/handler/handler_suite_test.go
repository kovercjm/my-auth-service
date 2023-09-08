package handler_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

func request(engine *gin.Engine, method, path string, body string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, bytes.NewBuffer([]byte(body)))
	Ω(err).Should(Succeed())
	response := httptest.NewRecorder()
	engine.ServeHTTP(response, req)
	return response
}

func requestWithAuthorization(engine *gin.Engine, method, path, body, token string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, bytes.NewBuffer([]byte(body)))
	Ω(err).Should(Succeed())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	response := httptest.NewRecorder()
	engine.ServeHTTP(response, req)
	return response
}
