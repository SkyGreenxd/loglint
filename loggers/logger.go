package loggers

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

// Logger описывает контракт для детектора логгера
type Logger interface {
	Name() string
	Matches(packagePath string, methodName string) bool
}

// BaseLogger базовая реализация логгера
type BaseLogger struct {
	name        string
	packagePath string
	methods     map[string]bool
}

func NewBaseLogger(name, packagePath string, methods []string) BaseLogger {
	methodsMap := make(map[string]bool, len(methods))
	for _, m := range methods {
		methodsMap[m] = true
	}

	return BaseLogger{
		name:        name,
		packagePath: packagePath,
		methods:     methodsMap,
	}
}

func (b *BaseLogger) Name() string {
	return b.name
}

// Matches выполняет проверку вызова. Поддерживает сопоставление по полному пути
// или по суффиксу пакета (например, "uber-go/zap" совпадет с "go.uber.org/zap").
func (b *BaseLogger) Matches(packagePath string, methodName string) bool {
	if packagePath != b.packagePath && !hasSuffix(packagePath, b.packagePath) {
		return false
	}

	return b.methods[methodName]
}

// IsMethodCall проверяет, является ли выражение вызовом метода (SelectorExpr).
func IsMethodCall(call *ast.CallExpr) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}

	return sel.Sel.Name, true
}

// GetPackagePath извлекает путь к пакету, которому принадлежит вызываемый метод.
func GetPackagePath(pass *analysis.Pass, sel *ast.SelectorExpr) string {
	if obj, ok := pass.TypesInfo.Uses[sel.Sel]; ok {
		if pkg := obj.Pkg(); pkg != nil {
			return pkg.Path()
		}
	}

	if tv, ok := pass.TypesInfo.Types[sel.X]; ok {
		return extractPathFromType(tv.Type)
	}

	return ""
}

// extractPathFromType рекурсивно разворачивает тип (указатели, именованные типы),
// чтобы найти путь к исходному пакету.
func extractPathFromType(t types.Type) string {
	switch typ := t.(type) {
	case *types.Named:
		if pkg := typ.Obj().Pkg(); pkg != nil {
			return pkg.Path()
		}
	case *types.Pointer:
		return extractPathFromType(typ.Elem())
	case *types.Interface:
	}

	return ""
}

// hasSuffix безопасно проверяет, является ли пакет суффиксом пути.
// Учитывает границы папок, чтобы "my/zap" не совпало с "notza
func hasSuffix(s, suffix string) bool {
	if len(s) < len(suffix) {
		return false
	}

	if s[len(s)-len(suffix):] != suffix {
		return false
	}

	if len(s) == len(suffix) {
		return true
	}

	return s[len(s)-len(suffix)-1] == '/'
}
