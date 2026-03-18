package main

import (
	"log"

	"github.com/SkyGreenxd/loglint/analyzer"
	_ "github.com/SkyGreenxd/loglint/loggers"
	"github.com/SkyGreenxd/loglint/rules"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	runner := rules.NewRunner()
	if err := runner.Init(nil); err != nil {
		log.Fatalf("failed to initialize %s: %v", analyzer.AnalyzerName, err)
	}

	singlechecker.Main(analyzer.New(runner))
}
