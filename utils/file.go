package utils

import (
	"os"
	"strings"
	"text/template"
)

func ReplaceVars(s string, data any) string {
	tmpl, err := template.New("replace").Parse(s)
	if err != nil {
		return s
	}
	var b strings.Builder
	tmpl.Execute(&b, data)
	return b.String()
}

func RenderTemplate(tplPath string, outPath string, data any) error {
	t, err := template.ParseFiles(tplPath)
	if err != nil {
		return err
	}
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, data)
}
