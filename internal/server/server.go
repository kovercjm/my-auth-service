package server

import (
	"net/http"

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
		fx.Provide(NewLogger, repository.New, handler.New, New),
		fx.WithLogger(kFx.FxLogger),
		fx.Invoke(Lifecycle),
	)).Run()
}

type Server struct {
	GinEngine  *gin.Engine
	HTTPServer *http.Server

	Config *kServer.APIConfig
	Logger kLog.Logger
}

func New(logger kLog.Logger, handler *handler.Handler, repository dependency.Repository) (*Server, error) {
	s := &Server{}
	s.Config = &kServer.APIConfig{Port: 4201} // TODO hard-coded port
	s.Logger = logger

	s.GinEngine = gin.New()
	s.GinEngine.Use(
		kGin.APILogging(s.Logger.NoCaller()),
		kGin.ErrorFormatter(),
		kGin.PanicRecovery(s.Logger.NoCaller()),
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
