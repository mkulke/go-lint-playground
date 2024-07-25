package rules_test

import (
	"go-lint-playground/internal/rules"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNoShortHandInit(t *testing.T) {
	testDataDir := analysistest.TestData()
	analysistest.Run(t, testDataDir, rules.NoShortHandInit, "./src/noshorthandinit/")
}
