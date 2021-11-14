package nipponcolors

import (
	// embed colors json
	_ "embed"
	"encoding/json"

	"github.com/gookit/color"
)

type shade int

const (
	dark shade = iota + 300
	bright
)

func (s shade) String() string {
	return [...]string{300: "dark", 301: "bright"}[s]
}

type hue int

const (
	red hue = iota + 305
	black
	white
	purple
	cyan
	green
	yellow
)

func (h hue) String() string {
	return [...]string{
		305: "red",
		306: "black",
		307: "white",
		308: "purple",
		309: "cyan",
		310: "green",
		311: "yellow",
	}[h]
}

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
