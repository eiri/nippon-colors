package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"unicode"
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

func main() {
	fileName := os.Args[1]
	if fileName == "" {
		log.Fatalf("Error: please provide color map file name")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error: can't open scan color map file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	e := make(map[string]interface{})
	acc := make([]map[string]interface{}, 0)
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.FieldsFunc(line, func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		})
		if len(f) >= 2 {
			switch f[0] {
			case "name":
				e[f[0]] = f[1]
			case "color":
				e[f[0]] = strings.Title(strings.ToLower(f[1]))
			case "Drgb":
				r, _ := strconv.Atoi(f[1])
				g, _ := strconv.Atoi(f[2])
				b, _ := strconv.Atoi(f[3])
				e["red"], e["green"], e["blue"] = r, g, b
				acc = append(acc, e)
				e = make(map[string]interface{})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: can't scan color map source file: %s", err)
	}

	// generate const module
	tmpl, err := template.New("mod").Parse(ModSrc)
	if err != nil {
		log.Fatalf("Error: can't parse template: %s", err)
	}

	modFile, err := os.Create("const.go")
	if err != nil {
		log.Fatalf("Error: can't create const mod file: %s", err)
	}
	defer modFile.Close()

	if err := tmpl.Execute(modFile, acc); err != nil {
		log.Fatalf("Error: can't compile template: %s", err)
	}

	// generate embeddable json color map
	j, err := json.Marshal(acc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(j))
}
