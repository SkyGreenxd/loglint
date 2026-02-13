package rules

import (
	"go/token"
	"unicode"
)

const LowercaseName = "lowercase"

type LowercaseRule struct {
	BaseRule
}

func init() {
	RegisterRule(LowercaseName, NewLowercaseRule)
}

func NewLowercaseRule(cfg RuleConfig) (Rule, error) {
	rule := &LowercaseRule{
		BaseRule: NewBaseRule(LowercaseName),
	}
	if err := rule.Configure(cfg); err != nil {
		return nil, err
	}

	return rule, nil
}

// CheckRune проверяет первый символ сообщения.
// Если это буква, она должна быть в нижнем регистре.
// Цифры, спецсимволы и пробелы в начале сообщения игнорируются.
func (l *LowercaseRule) CheckRune(r rune, pos token.Pos) error {
	if pos != 0 || !unicode.IsLetter(r) {
		return nil
	}

	if unicode.IsUpper(r) {
		return ErrNotLowercase
	}

	return nil
}
