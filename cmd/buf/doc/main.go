package main

import (
	"github.com/bufbuild/buf/private/buf/cmd/buf"
	"github.com/bufbuild/buf/private/pkg/app"
	"os"
)

func main() {
	if err := buf.Docs("buf"); err != nil {
		os.Exit(app.GetExitCode(err))
	}
}
