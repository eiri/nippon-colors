package nipponcolors

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
