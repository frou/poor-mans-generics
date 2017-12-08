package common

import (
	"fmt"
	"log"
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

func Run(spec RunSpec) {
	if spec.CleanOnly {
		cleanGeneratedFiles(spec.OutputPathFmt)
		return
	}

	for k, v := range spec.TemplateVars {
		if strings.TrimSpace(v) == "" {
			log.Fatalf("Template var %q must not have an empty value", k)
		}
	}

	tmpl, err := template.ParseFiles(spec.TemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	fmtArg, ok := spec.TemplateVars[spec.OutputPathFmtArgKey]
	if !ok {
		log.Fatalf("OutputPathFmtArgKey %q is not in the TemplateVars map",
			spec.OutputPathFmtArgKey)
	}

	outFileName := fmt.Sprintf(spec.OutputPathFmt, fmtArg)
	outFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, spec.TemplateVars); err != nil {
		log.Fatal(err)
	}
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
