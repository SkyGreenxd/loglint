package main

import (
	"github.com/SkyGreenxd/loglint/analyzer"
	_ "github.com/SkyGreenxd/loglint/loggers"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.New(nil))
}
