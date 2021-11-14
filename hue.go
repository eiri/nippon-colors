package nipponcolors

type hue int

const (
	Red hue = iota + 305
	Black
	White
	Purple
	Cyan
	Green
	Yellow
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

func (h hue) Series() []Color {
	colors := make([]Color, 0)
	for i, c := range colorMap {
		if c.Hue == h {
			colors = append(colors, Color(i))
		}
	}
	return colors
}
