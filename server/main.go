package main

import (
	"fmt"

	"github.com/takaaa220/gocyclo-ide/server/internal/lsp"
)

func main() {
	if err := lsp.StartServer(); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
