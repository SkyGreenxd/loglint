package analyzer

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"strconv"

	"github.com/SkyGreenxd/loglint/loggers"
	"github.com/SkyGreenxd/loglint/rules"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const AnalyzerName = "loglint"

// New создаёт анализатор, захватывая runner.
func New(runner *rules.Runner) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     AnalyzerName,
		Doc:      "reports suspicious or non-standard log messages, created by SkyGreenxd.",
		Run:      makeRun(runner),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

// makeRun возвращает основную функцию выполнения анализа (run function).
func makeRun(runner *rules.Runner) func(*analysis.Pass) (any, error) {
	return func(pass *analysis.Pass) (any, error) {
		if runner == nil {
			return nil, ErrRunnerNotConfigured
		}

		inspectResult := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		activeLoggers := runner.GetLoggers()
		// Фильтруем узлы AST, оставляя только вызовы функций
		nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}
		inspectResult.Preorder(nodeFilter, func(n ast.Node) {
			call := n.(*ast.CallExpr)

			if !isLogCall(call, pass, activeLoggers) {
				return
			}

			message, pos := extractLogMessage(call, pass)
			if message == "" {
				return
			}

			issues := runner.Run(message, pos)
			for _, issue := range issues {
				reportIssue(pass, issue)
			}
		})

		return nil, nil
	}
}

// isLogCall проверяет, является ли выражение вызовом одного из поддерживаемых логгеров.
func isLogCall(call *ast.CallExpr, pass *analysis.Pass, activeLoggers []loggers.Logger) bool {
	methodName, ok := loggers.IsMethodCall(call)
	if !ok {
		return false
	}

	sel := call.Fun.(*ast.SelectorExpr)
	packagePath := loggers.GetPackagePath(pass, sel)
	if packagePath == "" {
		return false
	}

	for _, logger := range activeLoggers {
		if logger.Matches(packagePath, methodName) {
			return true
		}
	}

	return false
}

// extractLogMessage находит и возвращает строковое содержимое сообщения лога и его позицию.
func extractLogMessage(call *ast.CallExpr, pass *analysis.Pass) (string, token.Pos) {
	if len(call.Args) == 0 {
		return "", token.NoPos
	}

	for _, arg := range call.Args {
		message, ok := buildMessage(arg, pass)
		if !ok || message == "" {
			continue
		}

		return message, arg.Pos()
	}

	return "", token.NoPos
}

// buildMessage рекурсивно собирает строковое значение из выражения,
// учитывая константы, литералы и конкатенацию.
func buildMessage(expr ast.Expr, pass *analysis.Pass) (string, bool) {
	if tv, ok := pass.TypesInfo.Types[expr]; ok && tv.Value != nil && tv.Value.Kind() == constant.String {
		return constant.StringVal(tv.Value), true
	}

	switch e := expr.(type) {
	case *ast.BinaryExpr:
		if e.Op != token.ADD {
			break
		}

		left, lok := buildMessage(e.X, pass)
		right, rok := buildMessage(e.Y, pass)
		if !lok && !rok {
			return "", false
		}

		return left + right, true
	case *ast.BasicLit:
		if e.Kind != token.STRING {
			break
		}

		unquoted, err := strconv.Unquote(e.Value)
		if err != nil {
			return e.Value, true
		}
		return unquoted, true
	}

	if tv, ok := pass.TypesInfo.Types[expr]; ok && tv.Type != nil {
		if types.Identical(tv.Type, types.Typ[types.String]) {
			return "<value>", true
		}
	}

	return "", false
}

// reportIssue отправляет найденную проблему в отчет анализатора pass.
func reportIssue(pass *analysis.Pass, issue rules.Issue) {
	sev := issue.Severity.String()
	pass.Reportf(issue.Pos, "[%s] [%s] %s", issue.RuleName, sev, issue.Message)
}
