package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

const (
	// ModSrc is a template for constants module file
	ModSrc = `// THIS MODULE WAS GENERATED, DO NOT EDIT MANUALLY
package nipponcolors

const (
{{- range $i, $c := . }}
{{- if eq $i 0 }}
	{{ $c.color }} Color = iota
{{- else }}
	{{ $c.color }}
{{- end }}
{{- end }}
)
`
)

var (
	constCmd = &cobra.Command{
		Use:   "to_const",
		Short: "Generate const.go module from fiven colors json file",
		RunE:  runToConstE,
	}
)

func init() {
	constCmd.Flags().StringVarP(&input, "input", "i", "", "Colors json file")
	constCmd.MarkFlagRequired("input")
	constCmd.Flags().StringVarP(&output, "output", "o", "", "Output file for const.go source")

	rootCmd.AddCommand(constCmd)
}

func runToConstE(cmd *cobra.Command, args []string) error {
	j, err := ioutil.ReadFile(input)
	if err != nil {
		return err
	}

	var colors []map[string]interface{}
	err = json.Unmarshal(j, &colors)
	if err != nil {
		return err
	}

	tmpl, err := template.New("mod").Parse(ModSrc)
	if err != nil {
		return err
	}

	outputHandler := os.Stdout
	if output != "" {
		var err error
		outputHandler, err = os.Create(output)
		if err != nil {
			return err
		}
		defer outputHandler.Close()
	}

	return tmpl.Execute(outputHandler, colors)
}
