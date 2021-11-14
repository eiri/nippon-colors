package nipponcolors

type shade int

const (
	dark shade = iota + 300
	bright
)

func (s shade) String() string {
	return [...]string{300: "dark", 301: "bright"}[s]
}
