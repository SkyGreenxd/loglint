package rules

import (
	"go/token"

	"github.com/SkyGreenxd/loglint/loggers"
	"github.com/SkyGreenxd/loglint/pkg/e"
	"github.com/mitchellh/mapstructure"
)

// Runner управляет жизненным циклом правил и запускает их проверку для сообщений логов.
type Runner struct {
	runeCheckers []RuneChecker // Правила, проверяющие каждый символ по отдельности
	fullCheckers []Rule        // Правила, анализирующие всю строку целиком (например, регулярки)
	loggers      []loggers.Logger
}

func NewRunner() *Runner {
	return &Runner{}
}

// Init настраивает раннер на основе переданной карты настроек (из YAML)
func (runner *Runner) Init(settings map[string]any) error {
	const defaultTagName = "yaml"

	var cfg LinterConfig
	decoderCfg := &mapstructure.DecoderConfig{
		Result:           &cfg,
		WeaklyTypedInput: true,
		TagName:          defaultTagName,
	}

	decoder, err := mapstructure.NewDecoder(decoderCfg)
	if err != nil {
		return e.Wrap(ErrDecoderCreation.Error(), err)
	}

	if err := decoder.Decode(settings); err != nil {
		return e.Wrap(ErrDecodeSettings.Error(), err)
	}

	// Если логгеры не указаны в конфиге, подтягиваются все известные
	loggerRegistry := loggers.GetRegistry()
	if len(cfg.Loggers) == 0 {
		runner.loggers = loggerRegistry.GetAll()
	} else {
		runner.loggers = loggerRegistry.GetByNames(cfg.Loggers)
	}

	for name, ruleCfg := range cfg.Rules {
		if !ruleCfg.Enabled {
			continue
		}

		rule, err := globalRuleRegistry.GetRule(name, ruleCfg)
		if err != nil {
			return e.Wrap(name, err)
		}

		runner.Register(rule)
	}

	return nil
}

// Register распределяет правила по категориям
// runeCheckers проверяют символы
// fullCheckers проверяют строку целиком
func (runner *Runner) Register(rule Rule) {
	if !rule.Enabled() {
		return
	}

	if rc, ok := rule.(RuneChecker); ok {
		runner.runeCheckers = append(runner.runeCheckers, rc)
	} else {
		runner.fullCheckers = append(runner.fullCheckers, rule)
	}
}

// Run выполняет все проверки для конкретного сообщения лога.
func (runner *Runner) Run(message string, startPos token.Pos) []Issue {
	var issues []Issue

	if len(runner.runeCheckers) > 0 {
		skipped := make(map[string]bool)
		for i, r := range message {
			currentPos := startPos + token.Pos(i)
			for _, rc := range runner.runeCheckers {
				// Для оптимизации, если правило уже нашло ошибку в этой строке, пропускаем его
				if skipped[rc.Name()] {
					continue
				}

				if err := rc.CheckRune(r, token.Pos(i)); err != nil {
					issues = append(issues, rc.NewIssue(currentPos+1, err.Error()))
					skipped[rc.Name()] = true
				}
			}

			// Если все посимвольные правила уже сработали, выходим из цикла по строке
			if len(skipped) == len(runner.runeCheckers) {
				break
			}
		}
	}

	for _, fc := range runner.fullCheckers {
		issues = append(issues, fc.Check(message, startPos+1)...)
	}

	return issues
}

func (runner *Runner) GetLoggers() []loggers.Logger {
	return runner.loggers
}
