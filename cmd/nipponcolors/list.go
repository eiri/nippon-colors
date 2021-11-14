package main

import (
	"github.com/spf13/cobra"

	color "github.com/eiri/nippon-colors"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Ouput list fo all colors",
		Run:   runList,
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) {
	for i := 0; i < 250; i++ {
		c := color.Color(i)
		c.Printf("███ %s %s %s %s\n", c, c.Name(), c.Shade(), c.Hue())
	}
}
