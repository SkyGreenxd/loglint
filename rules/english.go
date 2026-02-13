package rules

import (
	"go/token"
	"unicode"
)

const EnglishName = "english"

type EnglishRule struct {
	BaseRule
}

func init() {
	RegisterRule(EnglishName, NewEnglishRule)
}

func NewEnglishRule(cfg RuleConfig) (Rule, error) {
	rule := &EnglishRule{
		BaseRule: NewBaseRule(EnglishName),
	}
	if err := rule.Configure(cfg); err != nil {
		return nil, err
	}

	return rule, nil
}

// CheckRune проверяет каждый символ.
// Если символ является буквой, он должен принадлежать набору ASCII (английский алфавит).
func (e *EnglishRule) CheckRune(r rune, pos token.Pos) error {
	if !unicode.IsLetter(r) {
		return nil
	}

	if r <= 127 {
		return nil
	}

	return ErrNotEnglish
}
