// +build ignore

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	flagElementTypeName = flag.String("type", "", "element type name")
	flagTemplatePath    = flag.String("template", "generator.tmpl", "path to template")
	flagClean           = flag.Bool("clean", false, "instead of generating, delete existing generated files")
)

type config struct {
	ElementTypeName string
	SetTypeName     string
	TemplatePath    string
	OutPath         string
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

const generatedFileNameFmt = "set_%v.go"

func run() error {
	if *flagClean {
		names, err := filepath.Glob(fmt.Sprintf(generatedFileNameFmt, "*"))
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

	if strings.TrimSpace(*flagElementTypeName) == "" {
		return errors.New("specify an element type name")
	}

	var cfg config

	cfg.ElementTypeName = strings.ToLower(*flagElementTypeName)
	cfg.SetTypeName = strings.Title(*flagElementTypeName) + "s"
	cfg.TemplatePath = *flagTemplatePath
	cfg.OutPath = fmt.Sprintf(generatedFileNameFmt, cfg.ElementTypeName)

	tmpl, err := template.ParseFiles(cfg.TemplatePath)
	if err != nil {
		return err
	}

	outFile, err := os.Create(cfg.OutPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return tmpl.Execute(outFile, cfg)
}
