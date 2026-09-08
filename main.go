package main

import (
	"github.com/grexcode-ai/opencode/cmd"
	"github.com/grexcode-ai/opencode/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
