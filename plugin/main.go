package plugin

import (
	"fmt"

	"github.com/SkyGreenxd/loglint/analyzer"
	_ "github.com/SkyGreenxd/loglint/loggers"
	"github.com/SkyGreenxd/loglint/rules"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglint", newPlugin)
}

func newPlugin(conf any) (register.LinterPlugin, error) {
	confMap, ok := conf.(map[string]interface{})
	if !ok && conf != nil {
		return nil, fmt.Errorf("settings must be a map[string]interface{}, got %T", conf)
	}

	runner := rules.NewRunner()
	if err := runner.Init(confMap); err != nil {
		return nil, fmt.Errorf("failed to initialize loglint rules: %w", err)
	}

	return &loglintPlugin{runner: runner}, nil
}

type loglintPlugin struct {
	runner *rules.Runner
}

func (p *loglintPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.New(p.runner)}, nil
}

func (p *loglintPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
