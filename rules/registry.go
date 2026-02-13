package rules

import (
	"sync"

	"github.com/SkyGreenxd/loglint/pkg/e"
)

// Глобальный реестр правил
var (
	globalRuleRegistry = &Registry{
		rules: make(map[string]Factory),
		mu:    sync.RWMutex{},
	}
)

// Factory конвертирует RuleConfig в правило Rule
type Factory func(cfg RuleConfig) (Rule, error)

type Registry struct {
	rules map[string]Factory
	mu    sync.RWMutex
}

// RegisterRule — функция, через которую каждое правило "записывается" в реестр
func RegisterRule(name string, factory Factory) {
	globalRuleRegistry.mu.Lock()
	defer globalRuleRegistry.mu.Unlock()

	if _, ok := globalRuleRegistry.rules[name]; ok {
		panic(e.Wrap(name, ErrRuleRegistered))
	}

	globalRuleRegistry.rules[name] = factory
}

// GetRule создает экземпляр правила по имени
func (r *Registry) GetRule(name string, cfg RuleConfig) (Rule, error) {
	r.mu.RLock()
	factory, exists := r.rules[name]
	r.mu.RUnlock()

	if !exists {
		return nil, ErrRuleNotFound
	}

	return factory(cfg)
}
