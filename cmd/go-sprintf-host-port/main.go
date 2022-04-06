package main

import (
	"flag"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/stbenjam/go-sprintf-host-port/pkg/analyzer"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin

func main() {
	// Don't use it: just to not crash on -unsafeptr flag from go vet
	flag.Bool("unsafeptr", false, "")

	singlechecker.Main(analyzer.Analyzer)
}
