package main

import (
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"my-auth-service/internal/server"
)

func main() {
	cmd := &cobra.Command{Use: "my-auth-service"}

	cmd.AddCommand(&cobra.Command{Use: "serve", Run: server.CMD})

	if err := cmd.Execute(); err != nil {
		log.Fatalln(errors.Wrap(err, "Execute my-auth-service failed"))
	}
}
