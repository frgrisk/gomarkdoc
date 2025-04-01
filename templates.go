// Code generated by gentmpl.sh; DO NOT EDIT.

package gomarkdoc

var templates = map[string]string{
	"doc": `{{- range (iter .Blocks) -}}
	{{- if eq .Entry.Kind "paragraph" -}}
		{{- template "text" .Entry.Spans -}}
	{{- else if eq .Entry.Kind "code" -}}
		{{- codeBlock "" (include "text" .Entry.Spans) -}}
	{{- else if eq .Entry.Kind "header" -}}
		{{- header .Entry.Level (include "text" .Entry.Spans) -}}
    {{- else if eq .Entry.Kind "list" -}}
        {{- template "list" .Entry.List -}}
	{{- end -}}
	{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
{{- end -}}
`,
	"example": `{{- accordionHeader .Title -}}
{{- spacer -}}

{{- template "doc" .Doc -}}
{{- spacer -}}

{{- codeBlock "go" .Code -}}
{{- spacer -}}

{{- if .HasOutput -}}

	{{- header 4 "Output" -}}
	{{- spacer -}}

	{{- codeBlock "" .Output -}}
	{{- spacer -}}

{{- end -}}

{{- accordionTerminator -}}

`,
	"file": `<!-- Code generated by gomarkdoc. DO NOT EDIT -->

{{if .Header -}}
	{{- .Header -}}
	{{- spacer -}}
{{- end -}}

{{- range .Packages -}}
	{{- template "package" . -}}
	{{- spacer -}}
{{- end -}}

{{- if .Footer -}}
	{{- .Footer -}}
	{{- spacer -}}
{{- end -}}

Generated by {{link "gomarkdoc" "https://github.com/frgrisk/gomarkdoc"}}
`,
	"func": `{{- if .Receiver -}}
	{{- rawAnchorHeader .Level (codeHref .Location | link (escape .Name) | printf "func \\(%s\\) %s" (escape .Receiver)) .Anchor -}}
{{- else -}}
	{{- rawAnchorHeader .Level (codeHref .Location | link (escape .Name) | printf "func %s") .Anchor -}}
{{- end -}}
{{- spacer -}}

{{- codeBlock "go" .Signature -}}
{{- spacer -}}

{{- template "doc" .Doc -}}

{{- if len .Examples -}}
	{{- spacer -}}

	{{- range (iter .Examples) -}}
		{{- template "example" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}

`,
	"import": `{{- codeBlock "go" .Import -}}`,
	"index": `{{- if len .Consts -}}

	{{- localHref "Constants" | link "Constants" | listEntry 0 -}}
	{{- inlineSpacer -}}

{{- end -}}

{{- if len .Vars -}}

	{{- localHref "Variables" | link "Variables" | listEntry 0 -}}
	{{- inlineSpacer -}}

{{- end -}}

{{- range .Funcs -}}

	{{- (link .Signature (rawLocalHref .Anchor)) | listEntry 0 -}}
	{{- inlineSpacer -}}

{{- end -}}

{{- range .Types -}}

	{{- (link .Title (rawLocalHref .Anchor)) | listEntry 0 -}}
	{{- inlineSpacer -}}

	{{- range .Funcs -}}
		{{- (link .Signature (rawLocalHref .Anchor)) | listEntry 1 -}}
		{{- inlineSpacer -}}
	{{- end -}}

	{{- range .Methods -}}
		{{- (link .Signature (rawLocalHref .Anchor)) | listEntry 1 -}}
		{{- inlineSpacer -}}
	{{- end -}}

{{- end -}}
`,
	"list": `{{- range (iter .Items) -}}
    {{- if eq .Entry.Kind "ordered" -}}
        {{- .Entry.Number -}}. {{ hangingIndent (include "doc" .Entry) 2 -}}
    {{- else -}}
        - {{ hangingIndent (include "doc" .Entry) 2 -}}
    {{- end -}}

    {{- if (not .Last) -}}
        {{- if $.BlankBetween -}}
            {{- spacer -}}
        {{- else -}}
            {{- inlineSpacer -}}
        {{- end -}}
    {{- end -}}

{{- end -}}`,
	"package": `{{- if eq .Name "main" -}}
	{{- header .Level .Dirname -}}
{{- else -}}
	{{- header .Level .Name -}}
{{- end -}}
{{- spacer -}}

{{- template "import" . -}}
{{- spacer -}}

{{- if len .Doc.Blocks -}}
	{{- template "doc" .Doc -}}
	{{- spacer -}}
{{- end -}}

{{- range (iter .Examples) -}}
	{{- template "example" .Entry -}}
	{{- spacer -}}
{{- end -}}

{{- header (add .Level 1) "Index" -}}
{{- spacer -}}

{{- template "index" . -}}

{{- if len .Consts -}}
	{{- spacer -}}

	{{- header (add .Level 1) "Constants" -}}
	{{- spacer -}}

	{{- range (iter .Consts) -}}
		{{- template "value" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}

{{- end -}}

{{- if len .Vars -}}
	{{- spacer -}}

	{{- header (add .Level 1) "Variables" -}}
	{{- spacer -}}

	{{- range (iter .Vars) -}}
		{{- template "value" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}

{{- end -}}

{{- if len .Funcs -}}
	{{- spacer -}}

	{{- range (iter .Funcs) -}}
		{{- template "func" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}

{{- if len .Types -}}
	{{- spacer -}}

	{{- range (iter .Types) -}}
		{{- template "type" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}
`,
	"text": `{{- range . -}}
	{{- if eq .Kind "text" -}}
		{{- escape .Text -}}
	{{- else if eq .Kind "rawText" -}}
		{{- .Text -}}
	{{- else if eq .Kind "autolink" -}}
		{{- .Text -}}
	{{- else if eq .Kind "link" -}}
		{{- link (escape .Text) .URL -}}
	{{- end -}}
{{- end -}}`,
	"type": `{{- rawAnchorHeader .Level (codeHref .Location | link (escape .Name) | printf "type %s") .Anchor -}}
{{- spacer -}}

{{- template "doc" .Doc -}}
{{- spacer -}}

{{- codeBlock "go" .Decl -}}

{{- if len .Consts -}}
	{{- spacer -}}

	{{- range (iter .Consts) -}}
		{{- template "value" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}

{{- if len .Vars -}}
	{{- spacer -}}

	{{- range (iter .Vars) -}}
		{{- template "value" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}

{{- if len .Examples -}}
	{{- spacer -}}

	{{- range (iter .Examples) -}}
		{{- template "example" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}

{{- if len .Funcs -}}
	{{- spacer -}}

	{{- range (iter .Funcs) -}}
		{{- template "func" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}

{{- if len .Methods -}}
	{{- spacer -}}

	{{- range (iter .Methods) -}}
		{{- template "func" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}

`,
	"value": `{{- anchor .Anchor -}}
{{- template "doc" .Doc -}}
{{- spacer -}}

{{- codeBlock "go" .Decl -}}

`,
}
