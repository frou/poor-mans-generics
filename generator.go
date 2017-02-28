// +build ignore

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	flagElementTypeName = flag.String("type", "", "element type name")
	flagTemplatePath    = flag.String("template", "generator.tmpl", "path to template")
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

func run() error {
	if strings.TrimSpace(*flagElementTypeName) == "" {
		return errors.New("specify an element type name")
	}

	var cfg config

	cfg.ElementTypeName = strings.ToLower(*flagElementTypeName)
	cfg.SetTypeName = strings.Title(*flagElementTypeName) + "Set"
	cfg.TemplatePath = *flagTemplatePath
	cfg.OutPath = fmt.Sprintf("set_%v.go", cfg.ElementTypeName)

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
