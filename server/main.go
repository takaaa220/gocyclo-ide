package main

import (
	"fmt"

	"github.com/takaaa220/gocyclo-ide/server/internal"
)

func main() {
	if err := internal.StartServer(); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
