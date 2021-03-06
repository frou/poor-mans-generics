package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type RunSpec struct {
	CleanOnly           bool
	TemplatePath        string
	OutputPathFmt       string
	OutputPathFmtArgKey string

	TemplateVars map[string]string
}

// https://golang.org/s/generatedcode
const advisoryComment = "// Code generated by github.com/frou/poor-mans-generics ; DO NOT EDIT."

func Run(spec RunSpec) error {
	if spec.CleanOnly {
		return cleanGeneratedFiles(spec.OutputPathFmt)
	}

	for k, v := range spec.TemplateVars {
		if strings.TrimSpace(v) == "" {
			return fmt.Errorf("Template var %q must not have an empty value", k)
		}
	}

	tmpl, err := template.ParseFiles(spec.TemplatePath)
	if err != nil {
		return err
	}

	fmtArg, ok := spec.TemplateVars[spec.OutputPathFmtArgKey]
	if !ok {
		return fmt.Errorf("OutputPathFmtArgKey %q is not in the TemplateVars map",
			spec.OutputPathFmtArgKey)
	}

	outFileName := fmt.Sprintf(spec.OutputPathFmt, fmtArg)
	outFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	fmt.Fprintln(outFile, advisoryComment)
	fmt.Fprintln(outFile)

	return tmpl.Execute(outFile, spec.TemplateVars)
}

func cleanGeneratedFiles(nameFmt string) error {
	names, err := filepath.Glob(fmt.Sprintf(nameFmt, "*"))
	if err != nil {
		return err
	}
	for _, n := range names {
		if err := os.Remove(n); err != nil {
			return err
		}
	}
	return nil
}

func Exit(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}
