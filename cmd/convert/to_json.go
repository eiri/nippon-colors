package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var (
	shades = map[string]int{"d": 300, "b": 301}
	hues   = map[string]int{
		"r": 305,
		"b": 306,
		"w": 307,
		"p": 308,
		"c": 309,
		"g": 310,
		"y": 311,
	}
)

var (
	jsonCmd = &cobra.Command{
		Use:   "to_json",
		Short: "Convert colors js file to json representation",
		RunE:  runToJSONE,
	}
)

func init() {
	jsonCmd.Flags().StringVarP(&input, "input", "i", "", "Colors js file")
	jsonCmd.MarkFlagRequired("input")
	jsonCmd.Flags().StringVarP(&output, "output", "o", "", "Output json file")

	rootCmd.AddCommand(jsonCmd)
}

func runToJSONE(cmd *cobra.Command, args []string) error {
	file, err := os.Open(input)
	if err != nil {
		return err
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
			case "f":
				e["shade"] = shades[f[1]]
			case "c":
				e["hue"] = hues[f[1]]
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
		return err
	}

	// generate embeddable json color map
	j, err := json.Marshal(acc)
	if err != nil {
		return err
	}

	if output == "" {
		fmt.Println(string(j))
	} else {
		ioutil.WriteFile(output, j, os.ModePerm)
	}

	return nil
}
