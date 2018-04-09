package common

import "flag"

var (
	FlagClean           = flag.Bool("clean", false, "instead of generating, delete existing generated files")
	FlagTemplatePath    = flag.String("template", "generator.gotemplate", "path to template")
	FlagElementTypeName = flag.String("type", "", "element type name")
)
