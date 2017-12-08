// +build ignore

package main

import (
	"flag"
	"strings"

	"github.com/frou/poor-mans-generics/internal/common"
)

func main() {
	flag.Parse()
	common.Exit(common.Run(common.RunSpec{
		CleanOnly:           *common.FlagClean,
		TemplatePath:        *common.FlagTemplatePath,
		OutputPathFmt:       "set_%v.go",
		OutputPathFmtArgKey: "ElementTypeName",
		TemplateVars: map[string]string{
			"ElementTypeName": strings.ToLower(*common.FlagElementTypeName),
			"SetTypeName":     strings.Title(*common.FlagElementTypeName) + "s",
		},
	}))
}
