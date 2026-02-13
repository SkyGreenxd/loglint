package analyzer

import (
	"testing"

	"github.com/SkyGreenxd/loglint/rules"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestGroups(t *testing.T) {
	tests := []struct {
		name     string
		settings map[string]any
		pkg      string
	}{
		{
			name: "Only Lowercase Rule",
			settings: map[string]any{
				"rules": map[string]any{
					"lowercase": map[string]any{"enabled": true},
					"sensitive": map[string]any{"enabled": false},
					"symbols":   map[string]any{"enabled": false},
					"english":   map[string]any{"enabled": false},
				},
			},
			pkg: "lowercase",
		},
		{
			name: "Only Symbols Rule",
			settings: map[string]any{
				"rules": map[string]any{
					"lowercase": map[string]any{"enabled": false},
					"sensitive": map[string]any{"enabled": false},
					"symbols":   map[string]any{"enabled": true},
					"english":   map[string]any{"enabled": false},
				},
			},
			pkg: "symbols",
		},
		{
			name: "Only English Rule",
			settings: map[string]any{
				"rules": map[string]any{
					"lowercase": map[string]any{"enabled": false},
					"sensitive": map[string]any{"enabled": false},
					"symbols":   map[string]any{"enabled": false},
					"english":   map[string]any{"enabled": true},
				},
			},
			pkg: "english",
		},
		{
			name: "Only Sensitive Rule",
			settings: map[string]any{
				"rules": map[string]any{
					"lowercase": map[string]any{"enabled": false},
					"sensitive": map[string]any{"enabled": true},
					"symbols":   map[string]any{"enabled": false},
					"english":   map[string]any{"enabled": false},
				},
			},
			pkg: "sensitive",
		},
		{
			name: "All Rules Enabled",
			settings: map[string]any{
				"rules": map[string]any{
					"lowercase": map[string]any{"enabled": true},
					"sensitive": map[string]any{"enabled": true},
					"symbols":   map[string]any{"enabled": true},
					"english":   map[string]any{"enabled": true},
				},
			},
			pkg: "multi",
		},
		{
			name: "Optional Patterns",
			settings: map[string]any{
				"rules": map[string]any{
					"lowercase": map[string]any{"enabled": false},
					"sensitive": map[string]any{
						"enabled": true,
						"options": map[string]any{
							"patterns": []string{
								"test",
								"admin",
								"тест",
								"skygreenxd",
								"user",
							},
						},
					},
					"symbols": map[string]any{"enabled": false},
					"english": map[string]any{"enabled": false},
				},
			},
			pkg: "optional",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := rules.NewRunner()
			err := runner.Init(tt.settings)
			if err != nil {
				t.Fatalf("failed to init runner: %v", err)
			}

			analysistest.Run(t, analysistest.TestData(), New(runner), tt.pkg)
		})
	}
}
