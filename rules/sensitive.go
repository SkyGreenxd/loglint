package rules

import (
	"fmt"
	"go/token"
	"regexp"

	"github.com/SkyGreenxd/loglint/pkg/e"
	"github.com/mitchellh/mapstructure"
)

const SensitiveName = "sensitive"

// SensitiveOptions содержит список регулярных выражений
type SensitiveOptions struct {
	Patterns []string `mapstructure:"patterns"`
}

type SensitiveRule struct {
	BaseRule
	patterns []*regexp.Regexp
}

func init() {
	RegisterRule(SensitiveName, NewSensitiveRule)
}

func NewSensitiveRule(cfg RuleConfig) (Rule, error) {
	rule := &SensitiveRule{
		BaseRule: NewBaseRule(SensitiveName),
	}

	if err := rule.Configure(cfg); err != nil {
		return nil, err
	}

	return rule, nil
}

// Configure настраивает правило, декодируя параметры из конфигурации и компилируя паттерны
func (r *SensitiveRule) Configure(cfg RuleConfig) error {
	if err := r.BaseRule.Configure(cfg); err != nil {
		return err
	}

	var opts SensitiveOptions
	if err := mapstructure.Decode(cfg.Options, &opts); err != nil {
		return e.Wrap(ErrDecodeOptions.Error(), err)
	}

	// Если паттерны не указаны, используются дефолтные
	patterns := opts.Patterns
	if len(patterns) == 0 {
		patterns = getDefaultPatterns()
	}

	return r.compilePatterns(patterns)
}

// compilePatterns компилирует регулярные выражения
func (r *SensitiveRule) compilePatterns(patterns []string) error {
	compiled := make([]*regexp.Regexp, 0, len(patterns))

	for _, p := range patterns {
		re, err := regexp.Compile(p)
		if err != nil {
			return fmt.Errorf("%w '%s': %w", ErrInvalidRegex, p, err)
		}
		compiled = append(compiled, re)
	}

	r.patterns = compiled
	return nil
}

// getDefaultPatterns возвращает набор стандартных регулярных выражений для поиска секретов
func getDefaultPatterns() []string {
	return []string{
		`(?i)\b(password|passwd|pwd)\b`,
		`(?i)\b(api[_-]?key|apikey)\b`,
		`(?i)\b(token|auth[_-]?token)\b`,
		`(?i)\b(secret|secret[_-]?key)\b`,
		`(?i)\b(access[_-]?key|accesskey)\b`,
		`(?i)\b(credentials?|creds?)\b`,
		`Bearer\s+[A-Za-z0-9\-._~+/]+=*`,
		`[0-9a-f]{32,}`,
	}
}

// Check выполняет поиск всех совпадений по паттернам и возвращает список найденных утечек
func (r *SensitiveRule) Check(message string, pos token.Pos) []Issue {
	var issues []Issue

	for _, re := range r.patterns {
		matches := re.FindAllStringIndex(message, -1)

		for _, match := range matches {
			startIndex := match[0]
			exactPos := pos + token.Pos(startIndex)
			issues = append(issues, r.NewIssue(exactPos, ErrSensitiveData.Error()))
		}
	}

	return issues
}
