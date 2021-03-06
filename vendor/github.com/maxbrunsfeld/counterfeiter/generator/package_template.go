package generator

import (
	"strings"
	"text/template"
)

var packageFuncs template.FuncMap = template.FuncMap{
	"ToLower":  strings.ToLower,
	"UnExport": unexport,
	"Replace":  strings.Replace,
	"Generate": func() string { return "go:generate" }, // yes, this seems insane but ensures that we can use `go generate ./...` from the main package
}

const packageTemplate string = `// Code generated by counterfeiter. DO NOT EDIT.
package {{.DestinationPackage}}

import (
	{{- range .Imports}}
	{{.Alias}} "{{.Path}}"
	{{- end}}
)

//{{.Generate}} counterfeiter . {{.Name}}

// {{.Name}} is a generated interface representing the exported functions
// in the {{.TargetPackage}} package.
type {{.Name}} interface {
  {{- range .Methods}}
  {{.Name}}({{.Params.AsNamedArgsWithTypes}}) {{.Returns.AsReturnSignature}}
  {{- end}}
}

type {{.Name}}Shim struct {}

{{- range .Methods}}
func (p *{{.FakeName}}Shim) {{.Name}}({{.Params.AsNamedArgsWithTypes}}) {{.Returns.AsReturnSignature}} {
  {{if .Returns.HasLength}}return {{end}}{{.FakePackage}}.{{.Name}}({{.Params.AsNamedArgsForInvocation}})
}
{{end}}
var _ {{.Name}} = new({{.Name}}Shim)
`
