package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/fx"
)

func Lifecycle(lifecycle fx.Lifecycle, server *Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%d", server.Config.Port)
			server.HTTPServer = &http.Server{
				Addr:    address,
				Handler: server.GinEngine,
			}
			go func() {
				server.Logger.Info("gin gen server starting", "listening", address)
				if err := server.HTTPServer.ListenAndServe(); err != nil {
					server.Logger.Error("gin gen server failed to serve", "error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Logger.Info("gin gen server is shutting down")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.HTTPServer.Shutdown(ctx); err != nil {
				server.Logger.Error("gin gen server shutdown failed", "error", err)
			}

			server.Logger.Info("gin gen server stopped gracefully")
			return nil
		},
	})
}
