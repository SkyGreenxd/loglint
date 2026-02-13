package rules

// LinterConfig определяет глобальные настройки линтера
type LinterConfig struct {
	Loggers []string              `yaml:"loggers"`
	Rules   map[string]RuleConfig `yaml:"rules"`
}

// RuleConfig определяет конфигурацию конкретного правила
type RuleConfig struct {
	Enabled  bool           `yaml:"enabled"`
	Severity string         `yaml:"severity"` // "INFO", "WARN", "ERROR"
	Options  map[string]any `yaml:"options"`
}
