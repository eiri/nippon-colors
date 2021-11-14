package main

import (
	color "github.com/eiri/nippon-colors"
)

func main() {
	for i := 0; i < 250; i++ {
		c := color.Color(i)
		c.Printf("███ %s %s %s %s\n", c, c.Name(), c.Shade(), c.Hue())
	}
}
