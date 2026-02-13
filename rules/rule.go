package rules

import (
	"go/token"
	"strings"
)

type Severity int

const (
	SeverityInfo Severity = iota
	SeverityWarning
	SeverityError
)

// Issue описывает найденную проблему
type Issue struct {
	RuleName string
	Message  string
	Severity Severity
	// Suggestion string предложение исправления
	Pos token.Pos
}

// Rule — контракт для любой проверки
type Rule interface {
	Name() string                                // возвращает имя правила
	Check(message string, pos token.Pos) []Issue // Проверяет сообщение на соответствие правилу
	Enabled() bool                               // Возвращает true если правило включено, false если отключено
	Configure(config RuleConfig) error           // Настраивает правило из конфигурации
	NewIssue(pos token.Pos, message string) Issue
}

// RuneChecker — интерфейс для правил, которые умеют проверять один символ
type RuneChecker interface {
	Rule
	// CheckRune принимает индекс token.Pos(i) относительно начала строки (0, 1, 2...).
	CheckRune(r rune, pos token.Pos) error
}

type BaseRule struct {
	name   string
	config RuleConfig
}

func NewBaseRule(name string) BaseRule {
	return BaseRule{
		name:   name,
		config: RuleConfig{Enabled: true},
	}
}

func (b *BaseRule) Name() string {
	return b.name
}

func (b *BaseRule) Enabled() bool {
	return b.config.Enabled
}

func (b *BaseRule) Configure(config RuleConfig) error {
	b.config = config
	return nil
}

// Check метод заглушка для RuneChecker
func (b *BaseRule) Check(message string, pos token.Pos) []Issue {
	return nil
}

// GetSeverity возвращает типизированный уровень ошибки из конфига
func (b *BaseRule) GetSeverity() Severity {
	switch strings.ToUpper(b.config.Severity) {
	case "INFO":
		return SeverityInfo
	case "ERROR":
		return SeverityError
	default:
		return SeverityWarning
	}
}

func (b *BaseRule) NewIssue(pos token.Pos, message string) Issue {
	return Issue{
		RuleName: b.name,
		Message:  message,
		Severity: b.GetSeverity(),
		Pos:      pos,
	}
}
