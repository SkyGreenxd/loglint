package rules

import "errors"

var (
	ErrRuleNotFound  = errors.New("rule not found")
	ErrSensitiveData = errors.New("log message contains sensitive data (potential password, key or token)")
	ErrNotEnglish    = errors.New("log message contains non-english characters")
	ErrNotLowercase  = errors.New("log message must start with a lowercase letter")
	ErrInvalidSymbol = errors.New("log message contains forbidden special characters or emojis")
)
