package rules

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var NoShortHandInit = &analysis.Analyzer{
	Name: "noshorthandinit",
	Doc:  "check for short-hand struct initialization",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		if _, ok := node.(*ast.CompositeLit); !ok {
			return true
		}

		compLit := node.(*ast.CompositeLit)
		if _, ok := compLit.Type.(*ast.Ident); !ok {
			return true
		}

		if len(compLit.Elts) == 0 {
			return false
		}

		elt := compLit.Elts[0]
		if _, ok := elt.(*ast.KeyValueExpr); ok {
			return false
		}

		pass.Reportf(compLit.Pos(), "avoid short-hand struct initialization")
		return false
	}

	for _, file := range pass.Files {
		ast.Inspect(file, inspect)
	}

	return nil, nil
}
