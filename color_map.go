package nipponcolors

import (
	// embed colors json
	_ "embed"
	"encoding/json"

	"github.com/gookit/color"
)

var (
	colorMap []NipponColor
)

//go:embed color.json
var colors []byte

// NipponColor is a wrapper for gookit/color's RGBColor type
type NipponColor struct {
	Name           string `json:"name"`
	Color          string `json:"color"`
	Shade          shade  `json:"shade"`
	Hue            hue    `json:"hue"`
	Red            uint8  `json:"red"`
	Green          uint8  `json:"green"`
	Blue           uint8  `json:"blue"`
	color.RGBColor `json:"-"`
}

func init() {
	err := json.Unmarshal(colors, &colorMap)
	if err != nil {
		panic(err)
	}
	for i, c := range colorMap {
		colorMap[i].RGBColor = color.RGB(c.Red, c.Green, c.Blue)
	}
}
