package main

import (
	"github.com/spf13/cobra"

	color "github.com/eiri/nippon-colors"
)

var (
	demoCmd = &cobra.Command{
		Use:   "demo",
		Short: "Ouput colors demo",
		Run:   runDemo,
	}
)

func init() {
	rootCmd.AddCommand(demoCmd)
}

func runDemo(cmd *cobra.Command, args []string) {
	for i := 0; i < 250; i++ {
		c := color.Color(i)
		c.Printf("███ %s %s %s %s\n", c, c.Name(), c.Shade(), c.Hue())
	}
}
