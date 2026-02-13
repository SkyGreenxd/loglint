package rules

import (
	"go/token"
	"unicode"
)

const SymbolsName = "symbols"

type SymbolsRule struct {
	BaseRule
}

func init() {
	RegisterRule(SymbolsName, NewSymbolsRule)
}

func NewSymbolsRule(cfg RuleConfig) (Rule, error) {
	rule := &SymbolsRule{
		BaseRule: NewBaseRule(SymbolsName),
	}
	if err := rule.Configure(cfg); err != nil {
		return nil, err
	}

	return rule, nil
}

func (s *SymbolsRule) CheckRune(r rune, pos token.Pos) error {
	if isSymbol(r) {
		return ErrInvalidSymbol
	}

	return nil
}

// isSymbol определяет, является ли символ запрещенным (спецсимволы, пунктуация, эмодзи)
func isSymbol(r rune) bool {
	return unicode.IsSymbol(r) || unicode.IsPunct(r)
}
