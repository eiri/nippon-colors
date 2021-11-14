package nipponcolors

// Color is an integer representation of a color
type Color uint8

// String implements stringer interface for a given color
func (c Color) String() string {
	return colorMap[c].Color
}

// Name returns original kanji name of a given color
func (c Color) Name() string {
	return colorMap[c].Name
}

// Print marshals arguments to according wrapped gookit.Color method
func (c Color) Print(args ...interface{}) {
	colorMap[c].Print(args...)
}

// Printf marshals arguments to according wrapped gookit.Color method
func (c Color) Printf(format string, args ...interface{}) {
	colorMap[c].Printf(format, args...)
}

// Println marshals arguments to according wrapped gookit.Color method
func (c Color) Println(args ...interface{}) {
	colorMap[c].Println(args...)
}

// Shade returns string representation of a shade for a given color
func (c Color) Shade() string {
	return colorMap[c].Shade.String()
}

// Hue returns string representation of a hue for a given color
func (c Color) Hue() string {
	return colorMap[c].Hue.String()
}
