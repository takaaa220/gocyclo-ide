package internal

import (
	"fmt"

	"github.com/fzipp/gocyclo"
)

func calculateFunctionComplexities(filePath string) []string {
	stats := gocyclo.Analyze([]string{filePath}, nil)
	shownStats := stats.SortAndFilter(-1, -1)

	result := make([]string, len(shownStats))
	for i, stat := range shownStats {
		result[i] = fmt.Sprintf("%d %s %s", stat.Complexity, stat.FuncName, stat.Pos)
	}

	return result
}
