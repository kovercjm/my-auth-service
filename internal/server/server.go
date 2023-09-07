package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	kFx "github.com/kovercjm/tool-go/dependency_injection/fx"
	kLog "github.com/kovercjm/tool-go/logger"
	kServer "github.com/kovercjm/tool-go/server"
	kGin "github.com/kovercjm/tool-go/server/gin"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"my-auth-service/internal/handler"
	"my-auth-service/internal/handler/gen"
	"my-auth-service/internal/middleware/dependency"
	"my-auth-service/internal/repository"
)

var CMD = func(cmd *cobra.Command, args []string) {
	fx.New(fx.Module(
		"my-auth-service",
		fx.Provide(newLogger, repository.New, handler.New, newServer),
		fx.WithLogger(kFx.FxLogger),
		fx.Invoke(func(lifecycle fx.Lifecycle, server *Server) {
			lifecycle.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					address := fmt.Sprintf(":%d", server.config.Port)
					server.HTTPServer = &http.Server{
						Addr:    address,
						Handler: server.GinEngine,
					}
					go func() {
						server.logger.Info("gin gen server starting", "listening", address)
						if err := server.HTTPServer.ListenAndServe(); err != nil {
							server.logger.Error("gin gen server failed to serve", "error", err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					server.logger.Info("gin gen server is shutting down")

					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					if err := server.HTTPServer.Shutdown(ctx); err != nil {
						server.logger.Error("gin gen server shutdown failed", "error", err)
					}

					server.logger.Info("gin gen server stopped gracefully")
					return nil
				},
			})
		}),
	)).Run()
}

type Server struct {
	GinEngine  *gin.Engine
	HTTPServer *http.Server

	config *kServer.APIConfig
	logger kLog.Logger
}

func newServer(logger kLog.Logger, handler *handler.Handler, repository dependency.Repository) (*Server, error) {
	s := &Server{}
	s.config = &kServer.APIConfig{Port: 4201} // TODO hard code port
	s.logger = logger

	s.GinEngine = gin.New()
	s.GinEngine.Use(
		kGin.APILogging(s.logger.NoCaller()),
		kGin.ErrorFormatter(),
		kGin.PanicRecovery(s.logger.NoCaller()),
		AuthenticateMiddleware(
			repository,
			handler.UsersPost,
			handler.UsersIdDelete,
			handler.RolesPost,
			handler.RolesIdDelete,
			handler.UsersUserIDRolesRoleIDPost,
			handler.AuthLoginPost,
		),
	)

	group := s.GinEngine.Group("/")
	gen.RegisterUserAPI(group, handler)
	gen.RegisterRoleAPI(group, handler)
	gen.RegisterAuthAPI(group, handler)

	return s, nil
}
