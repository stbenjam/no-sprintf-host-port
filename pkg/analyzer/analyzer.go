package analyzer

import (
	"go/ast"
	"go/token"
	"regexp"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gosprintfhostport",
	Doc:      "Checks that sprintf is not used to construct a host:port combination in a URL.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		callExpr := node.(*ast.CallExpr)

		selector, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		pkg, ok := selector.X.(*ast.Ident)
		if !ok {
			return
		}
		if pkg.Name != "fmt" || selector.Sel.Name != "Sprintf" {
			return
		}

		if len(callExpr.Args) < 2 {
			return
		}

		// Let's see if our format string is a string literal.
		fsRaw, ok := callExpr.Args[0].(*ast.BasicLit)
		if !ok {
			return
		}
		if fsRaw.Kind != token.STRING {
			return
		}

		// Remove quotes
		fs := fsRaw.Value[1 : len(fsRaw.Value)-1]

		regexes := []*regexp.Regexp{
			// These check to see if it looks like a URI with a port, basically scheme://%s:<something else>,
			// or scheme://user:pass@%s:<something else>.
			// Matching requirements:
			//		- Scheme as per RFC3986 is ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
			//		- A format string substitution in the host portion, preceded by an optional username/password@
			//  	- A colon indicating a port will be specified
			regexp.MustCompile(`^[a-zA-Z0-9+-.]*://%s:[^@]*$`),
			regexp.MustCompile(`^[a-zA-Z0-9+-.]*://[^/]*@%s:.*$`),
		}

		for _, re := range regexes {
			if re.MatchString(fs) {
				pass.Reportf(node.Pos(), "host:port in url should be constructed with net.JoinHostPort and not directly with fmt.Sprintf")
				break
			}
		}
	})

	return nil, nil
}