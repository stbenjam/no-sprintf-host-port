package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/stbenjam/go-sprintf-host-port/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
