package internal

import (
	"go/token"

	"github.com/fzipp/gocyclo"
)

type functionComplexity struct {
	Complexity int
	FuncName   string
	Pos        token.Position
}

func calculateFunctionComplexities(filePath string) []functionComplexity {
	stats := gocyclo.Analyze([]string{filePath}, nil)
	shownStats := stats.SortAndFilter(-1, -1)

	result := make([]functionComplexity, len(shownStats))
	for i, stat := range shownStats {
		result[i] = functionComplexity{
			Complexity: stat.Complexity,
			FuncName:   stat.FuncName,
			Pos:        stat.Pos,
		}
	}

	return result
}
