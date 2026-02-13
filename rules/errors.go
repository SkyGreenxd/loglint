package rules

import "errors"

var (
	ErrRuleNotFound    = errors.New("rule not found")
	ErrSensitiveData   = errors.New("log message contains sensitive data (potential password, key or token)")
	ErrNotEnglish      = errors.New("log message contains non-english characters")
	ErrNotLowercase    = errors.New("log message must start with a lowercase letter")
	ErrInvalidSymbol   = errors.New("log message contains forbidden special characters or emojis")
	ErrRuleRegistered  = errors.New("rule already registered")
	ErrDecoderCreation = errors.New("failed to create decoder")
	ErrDecodeSettings  = errors.New("failed to decode settings")
	ErrDecodeOptions   = errors.New("failed to decode rule options")
	ErrInvalidRegex    = errors.New("invalid regular expression")
)
