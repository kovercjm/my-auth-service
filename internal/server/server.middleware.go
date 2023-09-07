package server

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"

	"my-auth-service/internal/middleware/dependency"
	"my-auth-service/internal/service"
)

func AuthenticateMiddleware(repository dependency.Repository, ignoreHandlers ...interface{}) gin.HandlerFunc {
	ignoreFuncs := map[string]struct{}{}
	for _, handler := range ignoreHandlers {
		funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		pieces := strings.Split(funcName, ".")
		funcName = pieces[len(pieces)-1]
		ignoreFuncs[funcName] = struct{}{}
	}
	return func(ctx *gin.Context) {
		funcName := runtime.FuncForPC(reflect.ValueOf(ctx.Handler()).Pointer()).Name()
		pieces := strings.Split(funcName, ".")
		funcName = pieces[len(pieces)-1]

		// TODO handle not registered path
		if strings.Contains(funcName, "func") {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		if _, ok := ignoreFuncs[funcName]; ok {
			ctx.Next()
			return
		}

		// TODO handle token in cookie
		bearerToken := strings.TrimSpace(ctx.GetHeader("Authorization"))
		bearerToken = strings.TrimPrefix(bearerToken, "Bearer ")
		if bearerToken == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, err := service.ParseToken(repository, bearerToken)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("token", bearerToken)
		ctx.Set("userID", claims.UserID) // TODO to extended

		ctx.Next()
	}
}
