package nipponcolors

type shade int

const (
	Dark shade = iota + 300
	Bright
)

func (s shade) String() string {
	return [...]string{300: "dark", 301: "bright"}[s]
}

func (s shade) Series() []Color {
	colors := make([]Color, 0)
	for i, c := range colorMap {
		if c.Shade == s {
			colors = append(colors, Color(i))
		}
	}
	return colors
}
