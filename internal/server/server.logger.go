package server

import (
	kInit "github.com/kovercjm/tool-go/init"
	kLog "github.com/kovercjm/tool-go/logger"
)

func NewLogger() (kLog.Logger, error) {
	return kInit.NewLogger(&kLog.Config{
		Debug:       true,
		Development: true,
	})
}
