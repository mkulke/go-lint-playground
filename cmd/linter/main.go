package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"go-lint-playground/internal/rules"
)

func main() {
	singlechecker.Main(rules.NoShortHandInit)
}
